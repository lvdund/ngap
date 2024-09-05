package aper

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math/bits"
	"reflect"
)

const (
	Zero bool = false
	One  bool = true
)

type aperWriter struct {
	w     io.Writer
	bytes []byte
	index uint
}

func NewWriter(w io.Writer) *aperWriter {
	return &aperWriter{
		w:     w,
		index: 0,
	}
}

func (aw *aperWriter) align() error {
	if aw.index > 0 {
		if _, err := aw.w.Write(aw.bytes[:]); err != nil {
			return err
		}
		aw.index = 0
	}
	return nil
}

func (aw *aperWriter) writeBool(bit bool) error {
	var err error
	if bit {
		err = aw.putBitsValue(1, 1)
	} else {
		err = aw.putBitsValue(0, 1)
	}
	return err
}

func (aw *aperWriter) putBitString(content []byte, nbits uint) error {
	var err error
	if nbits > uint(8*len(content)) {
		return aperError("writeBits", ErrUnderflow)
	}
	if nbits == 0 { //write nothing
		return nil
	}
	content = content[:(nbits+7)>>3]
	if aw.index == 0 {
		aw.index = (nbits & 0x7)
		aw.bytes = appendBytes(aw.bytes, content)
	} else {
		bitsLeft := 8 - aw.index
		currentByte := len(aw.bytes) - 1
		if nbits <= bitsLeft {
			aw.bytes[currentByte] |= (content[0] >> aw.index)
		} else {
			content = appendBytes([]byte{0x00}, content)
			shiftBytes := ShiftBytes(content, 8-int(aw.index))
			aw.bytes[currentByte] |= shiftBytes[0]
			aw.bytes = appendBytes(aw.bytes, shiftBytes[1:])
			content = content[1:]
		}
		aw.index = (nbits + aw.index) & 0x07
		return err
	}
	return nil
}

func (aw *aperWriter) putBitsValue(value uint64, nbits uint) error {
	if nbits > 64 {
		return aperError("writeValue", ErrUnderflow)
	}
	value = value << (64 - nbits)
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], value)
	return aw.putBitString(buf[:], nbits)
}

func (aw *aperWriter) writeConstraintValue(valueRange int64, value uint64) error {
	var nbytes uint
	if valueRange <= 255 {
		return aw.putBitsValue(value, uint(bits.Len64(uint64(valueRange-1))))
	} else if valueRange == 256 {
		nbytes = 1
	} else if valueRange <= 65536 {
		nbytes = 2
	} else {
		return aperError("writeConstrainValue", ErrOverflow)
	}
	aw.align()
	return aperError("writeConstrainValue", aw.putBitsValue(value, nbytes*8))
}

func (aw *aperWriter) writeNormallySmallNonNegativeValue(value uint64) (err error) {
	defer func() {
		err = aperError("writeNormallySmallNonNegativeValue", err)
	}()
	if value < 64 {
		if err = aw.writeBool(Zero); err != nil {
			return
		}
		err = aw.putBitsValue(value, 6)
	} else {
		if err = aw.writeBool(One); err != nil {
			return
		}
		err = aw.putSemiConstrainedWholeNumber(value, 0)
	}
	return
}

func (aw *aperWriter) putSemiConstrainedWholeNumber(value uint64, lb uint64) (err error) {
	defer func() {
		err = aperError("writeSemiContrainWholeNumber", err)
	}()
	if lb > value {
		err = ErrUnderflow
	}
	value -= lb
	length := (bits.Len64(value) + 7) >> 3
	if err = aw.writeLength(66537, uint64(length)); err != nil {
		return
	}
	err = aw.putBitsValue(value, uint(length)*8)
	return
}

func (aw *aperWriter) writeLength(sizeRange int64, value uint64) (err error) {
	if sizeRange <= 65536 && sizeRange > 0 {
		return aw.writeConstraintValue(sizeRange, value)
	}
	aw.align()
	if value <= 127 {
		err = aw.putBitsValue(value, 8)
		return
	} else if value <= 16383 {
		value |= 0x8000
		err = aw.putBitsValue(value, 16)
		return
	}
	value = (value >> 14) | 0xc0
	err = aw.putBitsValue(value, 8)
	return
}

func (aw *aperWriter) writeString(
	bytes []byte,
	bitsLength uint64,
	extensive bool,
	lowerBoundPtr *int64,
	upperBoundPtr *int64,
	isBitString bool,
) error {
	lb, ub, sizeRange, err := aw.handleConstraints(bitsLength, extensive, lowerBoundPtr, upperBoundPtr)
	if err != nil {
		return err
	}
	if ub > 65535 {
		sizeRange = -1
	}
	byteLen := uint64(len(bytes))
	var bitLen uint64
	if isBitString {
		bitLen = bitsLength
	}
	sizes := (bitsLength + 7) >> 3
	if isBitString && (8-bitsLength&0x7) != 8 {
		bytes[sizes-1] &= (0xff << (8 - bitsLength&0x7))
	}
	if !isBitString {
		sizes = byteLen
	}
	if sizeRange == 1 {
		if isBitString && bitsLength != uint64(ub) {
			return fmt.Errorf("bitString Length(%d) does not match fixed size: %d", bitsLength, ub)
		}
		if !isBitString && sizes != uint64(ub) {
			return fmt.Errorf("OctetString Length(%d) does not match fixed size: %d", len(bytes), ub)
		}
		if sizes > 2 {
			aw.align()
			if isBitString {
				aw.index = uint(ub & 0x7)
			}
			aw.bytes = appendBytes(aw.bytes, bytes)
		} else {
			if isBitString {
				return aw.putBitString(bytes, uint(bitLen))
			}
			return aw.putBitString(bytes, uint(len(bytes)*8))
		}
		return nil
	}
	rawLength := bitsLength - uint64(lb)
	if !isBitString {
		rawLength = uint64(len(bytes)) - uint64(lb)
	}
	var byteOffset, partOfRawLength uint64
	for {
		if rawLength > 65536 {
			partOfRawLength = 65536
		} else if rawLength >= 16384 {
			partOfRawLength = rawLength & 0xc000
		} else {
			partOfRawLength = rawLength
		}
		if err := aw.writeLength(sizeRange, partOfRawLength); err != nil {
			return err
		}
		partOfRawLength += uint64(lb)
		if partOfRawLength == 0 {
			return nil
		}
		aw.align()
		if isBitString {
			aw.bytes = appendBytes(aw.bytes, bytes[byteOffset:byteOffset+((partOfRawLength+7)>>3)])
		} else {
			aw.bytes = appendBytes(aw.bytes, bytes[byteOffset:byteOffset+partOfRawLength])
		}
		rawLength -= partOfRawLength - uint64(lb)
		if rawLength > 0 {
			byteOffset += partOfRawLength
		}else{
			break
		}
		if isBitString {
			aw.index = uint(partOfRawLength & 0x7)
		}
	}
	return nil
}

func (aw *aperWriter) writeBitString(
	bytes []byte,
	bitsLength uint64,
	extensive bool,
	lowerBoundPtr *int64,
	upperBoundPtr *int64,
) error {
	return aw.writeString(bytes, bitsLength, extensive, lowerBoundPtr, upperBoundPtr, true)
}

func (aw *aperWriter) writeOctetString(
	bytes []byte,
	extensive bool,
	lowerBoundPtr *int64,
	upperBoundPtr *int64,
) error {
	return aw.writeString(bytes, uint64(len(bytes)), extensive, lowerBoundPtr, upperBoundPtr, false)
}


func (aw *aperWriter) writeInteger(value int64, extensive bool, lowerBoundPtr *int64, upperBoundPtr *int64) error {
	lb,_, valueRange, err := aw.handleConstraints(uint64(value), extensive, lowerBoundPtr, upperBoundPtr)
	if err != nil {
		return err
	}
	unsignedValue := uint64(value)
	var rawLength uint
	if valueRange == 1 {
		return nil
	}
	if value < 0 {
		y := value >> 63
		unsignedValue = uint64(((value ^ y) - y)) - 1
	}
	if valueRange <= 0 {
		unsignedValue >>= 7
	} else if valueRange <= 65536 {
		return aw.writeConstraintValue(valueRange, uint64(value-lb))
	} else {
		unsignedValue >>= 8
	}
	for rawLength = 1; rawLength <= 127; rawLength++ {
		if unsignedValue == 0 {
			break
		}
		unsignedValue >>= 8
	}
	if valueRange <= 0 {
		aw.align()
		aw.bytes = append(aw.bytes, byte(rawLength))
	} else {
		unsignedValueRange := uint64(valueRange - 1)
		bitLength := bits.Len64(unsignedValueRange) 
		byteLen := uint((bitLength + 7) / 8) 
		bitLen := bits.Len(uint(int(byteLen)))
		if err := aw.putBitsValue(uint64(rawLength-1), uint(bitLen)); err != nil {
			return err
		}
	}
	rawLength *= 8
	aw.align()
	if valueRange < 0 {
		mask := int64(1<<rawLength - 1)
		return aw.putBitsValue(uint64(value&mask), rawLength)
	} else {
		value -= lb
		return aw.putBitsValue(uint64(value), rawLength)
	}
}

func (aw *aperWriter) writeEnumerated(value uint64, extensive bool, lowerBoundPtr *int64,
	upperBoundPtr *int64,
) error {
	if lowerBoundPtr == nil || upperBoundPtr == nil {
		return fmt.Errorf("ENUMERATED value constraint is error")
	}
	lb, ub := *lowerBoundPtr, *upperBoundPtr
	if lb < 0 || lb > ub {
		return fmt.Errorf("ENUMERATED value constraint is error")
	}
	fmt.Printf("Value: %d, lb: %d, ub: %d, extensive: %v\n", value, lb, ub, extensive)
	if extensive {
		if value <= uint64(ub) {
			if err := aw.writeBool(Zero); err != nil {
				return err
			}
			valueRange := ub - lb + 1
			if valueRange > 1 {
				return aw.writeConstraintValue(valueRange, value)
			}
		} else {
			if err := aw.writeBool(One); err != nil {
				return err
			}
			return aw.writeNormallySmallNonNegativeValue(value - uint64(ub) - 1)
		}
	} else {
		if value > uint64(ub) {
			return fmt.Errorf("ENUMERATED value is larger than upperbound: %d > %d", value, ub)
		}
		valueRange := ub - lb + 1
		if valueRange > 1 {
			return aw.writeConstraintValue(valueRange, value)
		}
	}
	return nil
}

func (aw *aperWriter) writeSequenceOf(v reflect.Value, params fieldParameters) error {
    var lb, ub, sizeRange int64 = 0, -1, -1
    numElements := int64(v.Len())
    if params.sizeLowerBound != nil && params.sizeUpperBound != nil {
        if *params.sizeLowerBound < 65536 && *params.sizeUpperBound < 65536 {
            lb, ub, sizeRange, _ = aw.handleConstraints(uint64(numElements), params.sizeExtensible, params.sizeLowerBound, params.sizeUpperBound)
        }
    }
    if numElements < lb {
        return fmt.Errorf("SEQUENCE OF Size is lower than lowerbound")
    } else if sizeRange == 1 {
        if numElements != ub {
            return fmt.Errorf("Encoding Length %d != fix-size %d", numElements, ub)
        }
    } else if sizeRange > 0 {
        if err := aw.writeConstraintValue(sizeRange, uint64(numElements-lb)); err != nil {
            return err
        }
    } else {
        aw.align()
        aw.bytes = append(aw.bytes, byte(numElements&0xff))
    }
    params.sizeExtensible = false
    params.sizeUpperBound = nil
    params.sizeLowerBound = nil

    for i := 0; i < v.Len(); i++ {
        if err := aw.makeField(v.Index(i), params); err != nil {
            return err
        }
    }
    return nil
}


func (aw *aperWriter) writeChoiceIndex(present int, extensive bool, upperBoundPtr *int64) error {
	var ub int64
	rawChoice := present - 1
	if upperBoundPtr == nil {
		return fmt.Errorf("The upper bound of CHIOCE is missing")
	} else if ub = *upperBoundPtr; ub < 0 {
		return fmt.Errorf("The upper bound of CHIOCE is negative")
	} else if extensive && rawChoice > int(ub) {
		return fmt.Errorf("Unsupport value of CHOICE type is in Extensed")
	}
	if err := aw.writeConstraintValue(ub+1, uint64(rawChoice)); err != nil {
		return err
	}
	return nil
}

func (aw *aperWriter) writeOpenType(v reflect.Value, params fieldParameters) error {
	buf := new(bytes.Buffer)
	awOpenType := NewWriter(buf)
	if err := awOpenType.makeField(v, params); err != nil {
		return err
	}
	openTypeBytes := awOpenType.bytes
	rawLength := uint64(len(awOpenType.bytes))
	var byteOffset, partOfRawLength uint64
	for {
		if rawLength > 65536 {
			partOfRawLength = 65536
		} else if rawLength >= 16384 {
			partOfRawLength = rawLength & 0xc000
		} else {
			partOfRawLength = rawLength
		}
		if err := aw.writeLength(-1, partOfRawLength); err != nil {
			return err
		}
		if partOfRawLength == 0 {
			return nil
		}
		aw.align()
		aw.bytes = appendBytes(aw.bytes, openTypeBytes[byteOffset:byteOffset+partOfRawLength])
		rawLength -= partOfRawLength
		if rawLength > 0 {
			byteOffset += partOfRawLength
		} else {
			aw.align()
			break
		}
	}
	return nil
}

func (aw *aperWriter) makeField(v reflect.Value, params fieldParameters) error {
	if !v.IsValid() {
		return fmt.Errorf("aper: cannot marshal nil value")
	}
	// If the field is an interface{} then recurse into it.
	if v.Kind() == reflect.Interface && v.Type().NumMethod() == 0 {
		return aw.makeField(v.Elem(), params)
	}
	if v.Kind() == reflect.Ptr {
		return aw.makeField(v.Elem(), params)
	}
	fieldType := v.Type()
	switch fieldType {
	case BitStringType:
		err := aw.writeBitString(v.Field(0).Bytes(), v.Field(1).Uint(), params.sizeExtensible, params.sizeLowerBound,
			params.sizeUpperBound)
		return err
	case ObjectIdentifierType:
		err := fmt.Errorf("Unsupport ObjectIdenfier type")
		return err
	case OctetStringType:
		err := aw.writeOctetString(v.Bytes(), params.sizeExtensible, params.sizeLowerBound, params.sizeUpperBound)
		return err
	case EnumeratedType:
		err := aw.writeEnumerated(v.Uint(), params.valueExtensible, params.valueLowerBound, params.valueUpperBound)
		return err
	}
	switch val := v; val.Kind() {
	case reflect.Bool:
		err := aw.writeBool(v.Bool())
		return err
	case reflect.Int, reflect.Int32, reflect.Int64:
		err := aw.writeInteger(v.Int(), params.valueExtensible, params.valueLowerBound, params.valueUpperBound)
		return err

	case reflect.Struct:
		structType := fieldType
		structField, err := structFieldCache.load(structType)
		if err != nil {
			return err
		}
		var optionalCount uint
		var optionalPresents uint64
		var sequenceType bool
		if params.valueExtensible {
			if err := aw.putBitsValue(0, 1); err != nil {
				return err
			}
		}
		sequenceType = len(structField) <= 0 || structField[0].FieldName != "Present"
		if sequenceType {
			for i := 0; i < len(structField); i++ {
				// for optional flag
				if structField[i].FieldParameters.optional {
					optionalCount++
					optionalPresents <<= 1
					if !v.Field(i).IsNil() {
						optionalPresents++
					}
				} else if v.Field(i).Type().Kind() == reflect.Ptr && v.Field(i).IsNil() {
					return fmt.Errorf("nil element in SEQUENCE type")
				}
			}
		}
		if optionalCount > 0 {
			if err := aw.putBitsValue(optionalPresents, optionalCount); err != nil {
				return err
			}
		}
		// CHOICE or OpenType
		if !sequenceType {
			present := int(v.Field(0).Int())
			if present == 0 {
				return fmt.Errorf("CHOICE or OpenType present is 0(present's field number)")
			} else if present >= len(structField) {
				return fmt.Errorf("Present is bigger than number of struct field")
			} else if params.openType {
				if params.referenceFieldValue == nil {
					return fmt.Errorf("OpenType reference value is empty")
				}
				refValue := *params.referenceFieldValue
				if structField[present].FieldParameters.referenceFieldValue == nil ||
					*structField[present].FieldParameters.referenceFieldValue != refValue {
					return fmt.Errorf("reference value and present reference value is not match")
				}
				if err := aw.writeOpenType(val.Field(present), structField[present].FieldParameters); err != nil {
					return err
				}
			} else {
				if err := aw.writeChoiceIndex(present, params.valueExtensible, params.valueUpperBound); err != nil {
					return err
				}
				if err := aw.makeField(val.Field(present), structField[present].FieldParameters); err != nil {
					return err
				}
			}
			return nil
		}
		for i := 0; i < structType.NumField(); i++ {
			// optional
			if structField[i].FieldParameters.optional && optionalCount > 0 {
				optionalCount--
				if optionalPresents&(1<<optionalCount) == 0 {
					continue
				} else {
				}
			}
			// for open type reference
			tempFieldParameters := structField[i].FieldParameters
			if tempFieldParameters.openType {
				fieldName := tempFieldParameters.referenceFieldName
				var index int
				for index = 0; index < i; index++ {
					if structField[index].FieldName == fieldName {
						break
					}
				}
				if index == i {
					return fmt.Errorf("Open type is not reference to the other field in the struct")
				}
				tempFieldParameters.referenceFieldValue = new(int64)
				if value, err := getReferenceFieldValue(val.Field(index)); err != nil {
					return err
				} else {
					*tempFieldParameters.referenceFieldValue = value
				}
			}
			if err := aw.makeField(val.Field(i), tempFieldParameters); err != nil {
				return err
			}
		}
		return nil
	case reflect.Slice:
		err := aw.writeSequenceOf(v, params)
		return err
	case reflect.String:
		printableString := v.String()
		err := aw.writeOctetString([]byte(printableString), params.sizeExtensible, params.sizeLowerBound,
			params.sizeUpperBound)
		return err
	}
	return fmt.Errorf("unsupported: " + v.Type().String())
}

func Encode(val interface{}) ([]byte, error) {
	return MarshalWithParams(val, "")
}
func MarshalWithParams(val interface{}, params string) ([]byte, error) {
	buf := new(bytes.Buffer)
	aw := NewWriter(buf)
	err := aw.makeField(reflect.ValueOf(val), parseFieldParameters(params))
	if err != nil {
		return nil, err
	} else if len(aw.bytes) == 0 {
		aw.bytes = make([]byte, 1)
	}
	return aw.bytes, nil
}
