package aper

import (
	"bytes"
	"fmt"
	"io"
	"math/bits"
	"reflect"
)

type AperReader interface {
	// ReadBool() (bool, error)
	// ReadBits(uint64) ([]byte, error)
	ReadOctetString(*Constrain,bool) (OctetString, error)
	ReadBitString(*Constrain, bool) (BitString, uint, error)
	ReadInteger(*Constrain, bool) (Integer, error)
	ReadEnumerated(*Constrain, bool) (Enumerated, error)
	ReadOpenType(*Constrain, bool) (interface{} , error)
}

type aperReader struct {
	r     io.Reader
	bytes      []byte
	byteOffset uint64
	index      uint
}

func (r *aperReader) ReadConstrainNumber() (uint, error) {
	return 0, nil
}

func GetBitString(srcBytes []byte, index, numBits uint) ([]byte, error) {
	bitsLeft := uint(len(srcBytes))*8 - index
	if numBits > bitsLeft {
		return nil, fmt.Errorf("Get bits overflow, requireBits: %d, leftBits: %d", numBits, bitsLeft)
	}
	numBitsByteLen := (numBits + 7) >> 3
	dstBytes := make([]byte, numBitsByteLen)
	if numBitsByteLen == 0 {
		return dstBytes, nil
	}
	if modEight := numBits & 0x7; modEight != 0 {
		dstBytes[numBitsByteLen-1] &= 0xFF << (8 - modEight)
	}
	for i := uint(0); i < numBits; i++ {
		bytePos := (index + i) >> 3
		bitPos := 7 - ((index + i) & 0x7)
		bit := (srcBytes[bytePos] >> bitPos) & 1
		dstBytes[i>>3] |= bit << (7 - (i & 0x7))
	}
	return dstBytes, nil
}
// update byteoffset and index 
func (ar *aperReader) bitCarry() {
	ar.byteOffset += uint64(ar.index >> 3)
	ar.index = ar.index & 0x07
}
// get bit string and update index +  byteoffset
func (ar *aperReader) getBitString(numBits uint) (dstBytes []byte, err error) {
	if dstBytes, err = GetBitString(ar.bytes[ar.byteOffset:], ar.index, numBits); err !=nil{
		return
	}
	ar.index += numBits
	ar.bitCarry()
	return
}
// get value and update index + byte off set 
func (ar *aperReader) getBitsValue(numBits uint) (value uint64, err error) {
	var dstBytes []byte
	if dstBytes, err = GetBitString(ar.bytes[ar.byteOffset:], ar.index, numBits); err != nil {
		return
	}
	for i, j := 0, numBits; j >= 8; i, j = i+1, j-8 {
		value <<= 8
		value |= uint64(uint(dstBytes[i]))
	}
	if extraBits := numBits & 0x7; extraBits != 0 {
		shiftAmount := 8 - extraBits
		lastByte := dstBytes[len(dstBytes)-1] >> shiftAmount
		value >>= shiftAmount 
		value |= uint64(lastByte)
	}
	ar.index += numBits
	ar.bitCarry() 
	return value, nil
}

func (ar *aperReader) parseAlignBits() error {
	if (ar.index & 0x7) > 0 {
		alignBits := 8 - ((ar.index) & 0x7)
		if val, err := ar.getBitsValue(alignBits); err != nil || val != 0   {
			return aperError("alignBits is ",ErrValue)
		}
	} else if ar.index != 0 {
		ar.bitCarry()
	}
	return nil
}

func (ar *aperReader) parseConstraintValue(valueRange int64) (value uint64, err error) {
	var bytes uint
	if valueRange <= 255 {
		value, err = ar.getBitsValue(uint(bits.Len64(uint64(valueRange-1))))
		return value, err
	} else if valueRange == 256 {
		bytes = 1
	} else if valueRange <= 65536 {
		bytes = 2
	} else {
		err = aperError("writeConstrainValue", ErrOverflow)
		return value, err
	}
	if err = ar.parseAlignBits(); err != nil {
		return value, err
	}
	value, err = ar.getBitsValue(bytes * 8)
	return value, err
}

func (ar *aperReader) parseSemiConstrainedWholeNumber(lb uint64) (value uint64, err error) {
	var repeat bool
	var length uint64
	if length, err = ar.parseLength(-1, &repeat); err != nil {
		return
	}
	if length > 8 || repeat {
		err = fmt.Errorf("Too long length: %d", length)
		return
	}
	if value, err = ar.getBitsValue(uint(length) * 8); err != nil {
		return
	}
	value += lb
	return
}

func (ar *aperReader) parseNormallySmallNonNegativeWholeNumber() (value uint64, err error) {
	var notSmallFlag uint64
	if notSmallFlag, err = ar.getBitsValue(1); err != nil {
		return
	}
	if notSmallFlag == 1 {
		if value, err = ar.parseSemiConstrainedWholeNumber(0); err != nil {
			return
		}
	} else {
		if value, err = ar.getBitsValue(6); err != nil {
			return
		}
	}
	return
}

func (ar *aperReader) parseLength(sizeRange int64, repeat *bool) (value uint64, err error) {
	*repeat = false
	var firstByte uint64
	if sizeRange > 0 && sizeRange <= 65536 {
		return ar.parseConstraintValue(sizeRange)
	}
	if err = ar.parseAlignBits(); err != nil {
		return 0, err
	}
	if firstByte, err = ar.getBitsValue(8); err!=nil{
		return 0, err
	}
	switch {
	case firstByte&0x80 == 0:
		value = firstByte & 0x7F
		return value, nil
	case firstByte&0x40 == 0:
		secondByte, err := ar.getBitsValue(8)
		if err != nil {
			return 0, err
		}
		value = ((firstByte & 0x3F) << 8) | secondByte
		return value, nil
	default:
		numMultiplier := firstByte & 0x3F
		if numMultiplier < 1 || numMultiplier > 4 {
			return 0, aperError("Parse Length ",ErrConstrain )
		}
		*repeat = true
		value = 16384 * uint64(numMultiplier)
		return value, nil
	}
}

func (ar *aperReader) parseString(
	extensed bool,
	lowerBoundPtr *int64,
	upperBoundPtr *int64,
	isBitString bool,
) ([]byte, uint64, error) {
	lb, ub, sizeRange := parseConstraint(extensed, lowerBoundPtr, upperBoundPtr)
	var result []byte
	var bitLength uint64
	if !isBitString {
		result = OctetString("")
	}
	if sizeRange == 1 {
		var sizes uint64
		if isBitString {
			sizes = (uint64(ub) + 7) >> 3
			bitLength = uint64(ub)
		} else {
			sizes = uint64(ub)
		}
		if sizes > 2 {
			if err := ar.parseAlignBits(); err != nil {
				return nil, 0, err
			}
			if (ar.byteOffset + sizes) > uint64(len(ar.bytes)) {
				return nil, 0, aperError("PER data", ErrRange)
			}
			result = ar.bytes[ar.byteOffset : ar.byteOffset+sizes]
			ar.byteOffset += sizes
			if isBitString {
				ar.index = uint(ub & 0x7)
				if ar.index > 0 {
					ar.byteOffset--
				}
			}
		} else {
			var bitCount uint64
			bitCount = uint64(ub * 8)
			if isBitString {
				bitCount = uint64(ub)
			}
			bytes, err := ar.getBitString(uint(bitCount))
			if err != nil {
				return nil, 0, err
			}
			result = bytes
		}
		return result, bitLength, nil
	}
	repeat := false
	for {
		rawLength, err := ar.parseLength(sizeRange, &repeat)
		if err != nil {
			return nil, 0, err
		}
		rawLength += uint64(lb)
		if rawLength == 0 {
			return result, bitLength, nil
		}
		var sizes uint64
		if isBitString {
			sizes = (rawLength + 7) >> 3
		} else {
			sizes = rawLength
		}
		if err := ar.parseAlignBits(); err != nil {
			return nil, 0, err
		}
		if (ar.byteOffset + sizes) > uint64(len(ar.bytes)) {
			return nil, 0, aperError("PER data", ErrRange)
		}
		result = appendBytes(result, ar.bytes[ar.byteOffset:ar.byteOffset+sizes])
		ar.byteOffset += sizes
		if isBitString {
			bitLength += rawLength
			ar.index = uint(rawLength & 0x7)
			if ar.index != 0 {
				ar.byteOffset--
			}
		}
		if !repeat {
			break
		}
	}
	return result, bitLength, nil
}

func (ar *aperReader) ReadBitString(extensed bool, lowerBoundPtr *int64, upperBoundPtr *int64) (BitString, error) {
	bytes, bitLength, err := ar.parseString(extensed, lowerBoundPtr, upperBoundPtr, true)
	if err != nil {
		aperError("ReadBitString ", err)
		return BitString{}, err
	}
	return BitString{Bytes: bytes, BitLength: bitLength}, nil
}

func (ar *aperReader) readOctetString(extensed bool, lowerBoundPtr *int64, upperBoundPtr *int64) (OctetString, error) {
	bytes, _, err := ar.parseString(extensed, lowerBoundPtr, upperBoundPtr, false)
	if err != nil {
		aperError("readOctetString ", err)
		return OctetString{}, err
	}
	return OctetString(bytes), nil
}

func (ar *aperReader) readBool() (value bool, err error) {
	bit, err1 := ar.getBitsValue(1)
	if err1 != nil {
		err = err1
		return
	}
	if bit == 1 {
		value = true
	} else {
		value = false
	}
	return
}

func (ar *aperReader) readInteger(extensed bool, lowerBoundPtr *int64, upperBoundPtr *int64) (int64, error) {
	lb, ub, valueRange  := parseConstraint(extensed, lowerBoundPtr,upperBoundPtr)
	var rawLength uint
	if valueRange == 1 {
		return ub, nil
	} else if valueRange <= 0 {
		if err := ar.parseAlignBits(); err != nil {
			return int64(0), err
		}
		if ar.byteOffset >= uint64(len(ar.bytes)) {
			return int64(0), aperError("PER data ",ErrRange)
		}
		rawLength = uint(ar.bytes[ar.byteOffset])
		ar.byteOffset++
	} else if valueRange <= 65536 {
		rawValue, err := ar.parseConstraintValue(valueRange)
		if err != nil {
			return int64(0), err
		} else {
			return int64(rawValue) + lb, nil
		}
	} else {
		unsignedValueRange := uint64(valueRange - 1)
		bitLength := bits.Len64(unsignedValueRange) 
		byteLen := uint((bitLength + 7) / 8) 
		bitLen := bits.Len(uint(int(byteLen)))
		if tempLength, err := ar.getBitsValue(uint(bitLen)); err != nil {
			return int64(0), err
		} else {
			rawLength = uint(tempLength)
		}
		rawLength++
		if err := ar.parseAlignBits(); err != nil {
			return int64(0), err
		}
	}
	if rawValue, err := ar.getBitsValue(rawLength * 8); err != nil {
		return int64(0), err
	} else if valueRange < 0 {
		signedBitMask := uint64(1 << (rawLength*8 - 1))
		valueMask := signedBitMask - 1
		if rawValue&signedBitMask > 0 {
			return int64((^rawValue)&valueMask+1) * -1, nil
		}
		return int64(rawValue) + lb, nil
	} else {
		return int64(rawValue) + lb, nil
	}
}

func (ar *aperReader) readEnumerated(extensed bool, lowerBoundPtr *int64, upperBoundPtr *int64) (value uint64,
	err error,
) {
	if lowerBoundPtr == nil || upperBoundPtr == nil {
		err = fmt.Errorf("ENUMERATED value constraint is error")
		return
	}
	lb, ub := *lowerBoundPtr, *upperBoundPtr
	if lb < 0 || lb > ub {
		err = fmt.Errorf("ENUMERATED value constraint is error")
		return
	}
	if extensed {
		if value, err = ar.parseNormallySmallNonNegativeWholeNumber(); err != nil {
			return
		}
		value += uint64(ub) + 1
	} else {
		valueRange := ub - lb + 1
		if valueRange > 1 {
			value, err = ar.parseConstraintValue(valueRange)
		}
	}
	return
}

func (ar *aperReader) parseSequenceOf(sizeExtensed bool, params fieldParameters, sliceType reflect.Type) (
	reflect.Value, error,
) {
	var sliceContent reflect.Value
	var lb int64 = 0
	var sizeRange int64
	if *params.sizeLowerBound < 65536 && *params.sizeUpperBound < 65536 {
		lb, _, sizeRange  = parseConstraint(sizeExtensed, params.sizeLowerBound,params.sizeUpperBound )
	}
	var numElements uint64
	if sizeRange > 1 {
		if numElementsTmp, err := ar.parseConstraintValue(sizeRange); err != nil {
		} else {
			numElements = numElementsTmp
		}
		numElements += uint64(lb)
	} else if sizeRange == 1 {
		numElements += uint64(lb)
	} else {
		if err := ar.parseAlignBits(); err != nil {
			return sliceContent, err
		}
		if ar.byteOffset >= uint64(len(ar.bytes)) {
			err := aperError("PER data ",ErrRange)
			return sliceContent, err
		}
		numElements = uint64(ar.bytes[ar.byteOffset])
		ar.byteOffset++
	}
	params.sizeExtensible = false
	params.sizeUpperBound = nil
	params.sizeLowerBound = nil
	intNumElements := int(numElements)
	sliceContent = reflect.MakeSlice(sliceType, intNumElements, intNumElements)
	for i := 0; i < intNumElements; i++ {
		err := parseField(sliceContent.Index(i), ar, params)
		if err != nil {
			return sliceContent, err
		}
	}
	return sliceContent, nil
}

func (ar *aperReader) getChoiceIndex(extensed bool, upperBoundPtr *int64) (present int, err error) {
	if extensed {
		err = fmt.Errorf("Unsupport value of CHOICE type is in Extensed")
	} else if upperBoundPtr == nil {
		err = fmt.Errorf("The upper bound of CHIOCE is missing")
	} else if ub := *upperBoundPtr; ub < 0 {
		err = fmt.Errorf("The upper bound of CHIOCE is negative")
	} else if rawChoice, err1 := ar.parseConstraintValue(ub + 1); err1 != nil {
		err = err1
	} else {
		present = int(rawChoice) + 1
	}
	return
}

func getReferenceFieldValue(v reflect.Value) (value int64, err error) {
	fieldType := v.Type()
	switch v.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		value = v.Int()
	case reflect.Struct:
		if fieldType.Field(0).Name == "Present" {
			present := int(v.Field(0).Int())
			if present == 0 {
				err = fmt.Errorf("ReferenceField Value present is 0(present's field number)")
			} else if present >= fieldType.NumField() {
				err = fmt.Errorf("Present is bigger than number of struct field")
			} else {
				value, err = getReferenceFieldValue(v.Field(present))
			}
		} else {
			value, err = getReferenceFieldValue(v.Field(0))
		}
	default:
		err = fmt.Errorf("OpenType reference only support INTEGER")
	}
	return
}

func (ar *aperReader) parseOpenType(skip bool, v reflect.Value, params fieldParameters) error {
	buf := new(bytes.Buffer)
	arOpenType := &aperReader{buf,[]byte(""), 0, 0}
	repeat := false
	for {
		var rawLength uint64
		if rawLengthTmp, err := ar.parseLength(-1, &repeat); err != nil {
			return err
		} else {
			rawLength = rawLengthTmp
		}
		if rawLength == 0 {
			break
		} else if err := ar.parseAlignBits(); err != nil {
			return err
		}
		if (rawLength + ar.byteOffset) > uint64(len(ar.bytes)) {
			return aperError("PER data ",ErrRange)
		}
		arOpenType.bytes = appendBytes(arOpenType.bytes, ar.bytes[ar.byteOffset:ar.byteOffset+rawLength])
		ar.byteOffset += rawLength

		if !repeat {
			if err := ar.parseAlignBits(); err != nil {
				return err
			}
			break
		}
	}
	if skip {
		return nil
	} else {
		err := parseField(v, arOpenType, params)
		return err
	}
}

func parseField(v reflect.Value, ar *aperReader, params fieldParameters) error {
	fieldType := v.Type()
	if ar.byteOffset == uint64(len(ar.bytes)) {
		return fmt.Errorf("sequence truncated")
	}
	if v.Kind() == reflect.Ptr {
		ptr := reflect.New(fieldType.Elem())
		v.Set(ptr)
		return parseField(v.Elem(), ar, params)
	}
	sizeExtensible := false
	valueExtensible := false
	if params.sizeExtensible {
		if bitsValue, err1 := ar.getBitsValue(1); err1 != nil {
			return err1
		} else if bitsValue != 0 {
			sizeExtensible = true
		}
	}
	if params.valueExtensible && v.Kind() != reflect.Slice {
		if bitsValue, err1 := ar.getBitsValue(1); err1 != nil {
			return err1
		} else if bitsValue != 0 {
			valueExtensible = true
		}
	}
	switch fieldType {
	case BitStringType:
		bitString, err1 := ar.ReadBitString(sizeExtensible, params.sizeLowerBound, params.sizeUpperBound)
		if err1 != nil {
			return err1
		}
		v.Set(reflect.ValueOf(bitString))
		return nil
	case ObjectIdentifierType:
		return fmt.Errorf("Unsupport ObjectIdenfier type")
	case OctetStringType:
		if octetString, err := ar.readOctetString(sizeExtensible, params.sizeLowerBound, params.sizeUpperBound); err != nil {
			return err
		} else {
			v.Set(reflect.ValueOf(octetString))
			return nil
		}
	case EnumeratedType:
		if parsedEnum, err := ar.readEnumerated(valueExtensible, params.valueLowerBound,
			params.valueUpperBound); err != nil {
			return err
		} else {
			v.SetUint(parsedEnum)
			return nil
		}
	}
	switch val := v; val.Kind() {
	case reflect.Bool:
		if parsedBool, err := ar.readBool(); err != nil {
			return err
		} else {
			val.SetBool(parsedBool)
			return nil
		}
	case reflect.Int, reflect.Int32, reflect.Int64:
		if parsedInt, err := ar.readInteger(valueExtensible, params.valueLowerBound, params.valueUpperBound); err != nil {
			return err
		} else {
			val.SetInt(parsedInt)
			return nil
		}
	case reflect.Struct:
		structType := fieldType
		structField, err := structFieldCache.load(structType)
		if err != nil {
			return err
		}
		var optionalCount uint
		var optionalPresents uint64
		for i := 0; i < len(structField); i++ {
			if structField[i].FieldParameters.optional {
				optionalCount++
			}
		}

		if optionalCount > 0 {
			if optionalPresentsTmp, err := ar.getBitsValue(optionalCount); err != nil {
				return err
			} else {
				optionalPresents = optionalPresentsTmp
			}
		}

		if len(structField) > 0 && structField[0].FieldName == "Present" {
			var present int = 0
			if params.openType {
				if params.referenceFieldValue == nil {
					return fmt.Errorf("OpenType reference value is empty")
				}
				refValue := *params.referenceFieldValue
				for j, param := range structField {
					if j == 0 {
						continue
					}
					if param.FieldParameters.referenceFieldValue != nil && *param.FieldParameters.referenceFieldValue == refValue {
						present = j
						break
					}
				}
				if present == 0 {
					val.Field(0).SetInt(0)
					return ar.parseOpenType(true, reflect.Value{}, fieldParameters{})
				} else if present >= len(structField) {
					return fmt.Errorf("OpenType Present is bigger than number of struct field")
				} else {
					val.Field(0).SetInt(int64(present))
					return ar.parseOpenType(false, val.Field(present), structField[present].FieldParameters)
				}
			} else {
				if presentTmp, err := ar.getChoiceIndex(valueExtensible, params.valueUpperBound); err != nil {
				} else {
					present = presentTmp
				}
				val.Field(0).SetInt(int64(present))
				if present == 0 {
					return fmt.Errorf("CHOICE present is 0(present's field number)")
				} else if present >= len(structField) {
					return fmt.Errorf("CHOICE Present is bigger than number of struct field")
				} else {
					return parseField(val.Field(present), ar, structField[present].FieldParameters)
				}
			}
		}

		for i := 0; i < len(structField); i++ {
			if structField[i].FieldParameters.optional && optionalCount > 0 {
				optionalCount--
				if optionalPresents&(1<<optionalCount) == 0 {
					continue
				} else {
				}
			}
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
				if referenceFieldValue, err := getReferenceFieldValue(val.Field(index)); err != nil {
					return err
				} else {
					*tempFieldParameters.referenceFieldValue = referenceFieldValue
				}
			}
			if err := parseField(val.Field(i), ar, tempFieldParameters); err != nil {
				return err
			}
		}
		return nil
	case reflect.Slice:
		sliceType := fieldType
		if newSlice, err := ar.parseSequenceOf(sizeExtensible, params, sliceType); err != nil {
			return err
		} else {
			val.Set(newSlice)
			return nil
		}
	case reflect.String:
		if octetString, err := ar.readOctetString(sizeExtensible, params.sizeLowerBound, params.sizeUpperBound); err != nil {
			return err
		} else {
			printableString := string(octetString)
			val.SetString(printableString)
			return nil
		}
	}
	return fmt.Errorf("unsupported: " + v.Type().String())
}

func Decode(b []byte, value interface{}) error {
	params := ""
	v := reflect.ValueOf(value).Elem()
	buf := new(bytes.Buffer)
	ar := &aperReader{buf,b, 0, 0}
	return parseField(v, ar, parseFieldParameters(params))
}

