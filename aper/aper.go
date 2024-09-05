package aper

import (
	//"fmt"
	"reflect"
	"strconv"
	"strings"
)

// BIT STRING

// BitString is for an ASN.1 BIT STRING type, BitLength means the effective bits.
type BitString struct {
	Bytes     []byte // bits packed into bytes.
	BitLength uint64 // length in bits.
}

type Constrain struct {
	Lb int64
	Ub int64
}

func (c *Constrain) Range() uint64 {
	return uint64(c.Ub - c.Lb + 1)
}
// OCTET STRING

// OctetString is for an ASN.1 OCTET STRING type
type OctetString []byte

// OBJECT IDENTIFIER

// ObjectIdentifier is for an ASN.1 OBJECT IDENTIFIER type
type ObjectIdentifier []byte

// ENUMERATED

// An Enumerated is represented as a plain uint64.
type Enumerated uint64

var (
	// BitStringType is the type of BitString
	BitStringType = reflect.TypeOf(BitString{})
	// OctetStringType is the type of OctetString
	OctetStringType = reflect.TypeOf(OctetString{})
	// ObjectIdentifierType is the type of ObjectIdentify
	ObjectIdentifierType = reflect.TypeOf(ObjectIdentifier{})
	// EnumeratedType is the type of Enumerated
	EnumeratedType = reflect.TypeOf(Enumerated(0))
)

type fieldParameters struct {
	optional            bool   // true iff the type has OPTIONAL tag.
	sizeExtensible      bool   // true iff the size can be extensed.
	valueExtensible     bool   // true iff the value can be extensed.
	sizeLowerBound      *int64 // a sizeLowerBound is the minimum size of type constraint(maybe nil).
	sizeUpperBound      *int64 // a sizeUpperBound is the maximum size of type constraint(maybe nil).
	valueLowerBound     *int64 // a valueLowerBound is the minimum value of type constraint(maybe nil).
	valueUpperBound     *int64 // a valueUpperBound is the maximum value of type constraint(maybe nil).
	defaultValue        *int64 // a default value for INTEGER and ENUMERATED typed fields (maybe nil).
	openType            bool   // true iff this type is opentype.
	referenceFieldName  string // the field to get to get the corresrponding value of this type(maybe nil).
	referenceFieldValue *int64 // the field value which map to this type(maybe nil).
}

// Given a tag string with the format specified in the package comment,
// parseFieldParameters will parse it into a fieldParameters structure,
// ignoring unknown parts of the string. TODO:PrintableString
func parseFieldParameters(str string) (params fieldParameters) {
	vAny, ok := fieldParametersCache.Load(str)
	if ok {
		return vAny.(fieldParameters)
	}

	for _, part := range strings.Split(str, ",") {
		switch {
		case part == "optional":
			params.optional = true
		case part == "sizeExt":
			params.sizeExtensible = true
		case part == "valueExt":
			params.valueExtensible = true
		case strings.HasPrefix(part, "sizeLB:"):
			i, err := strconv.ParseInt(part[7:], 10, 64)
			if err == nil {
				params.sizeLowerBound = new(int64)
				*params.sizeLowerBound = i
			}
		case strings.HasPrefix(part, "sizeUB:"):
			i, err := strconv.ParseInt(part[7:], 10, 64)
			if err == nil {
				params.sizeUpperBound = new(int64)
				*params.sizeUpperBound = i
			}
		case strings.HasPrefix(part, "valueLB:"):
			i, err := strconv.ParseInt(part[8:], 10, 64)
			if err == nil {
				params.valueLowerBound = new(int64)
				*params.valueLowerBound = i
			}
		case strings.HasPrefix(part, "valueUB:"):
			i, err := strconv.ParseInt(part[8:], 10, 64)
			if err == nil {
				params.valueUpperBound = new(int64)
				*params.valueUpperBound = i
			}
		case strings.HasPrefix(part, "default:"):
			i, err := strconv.ParseInt(part[8:], 10, 64)
			if err == nil {
				params.defaultValue = new(int64)
				*params.defaultValue = i
			}
		case part == "openType":
			params.openType = true
		case strings.HasPrefix(part, "referenceFieldName:"):
			params.referenceFieldName = part[19:]
		case strings.HasPrefix(part, "referenceFieldValue:"):
			i, err := strconv.ParseInt(part[20:], 10, 64)
			if err == nil {
				params.referenceFieldValue = new(int64)
				*params.referenceFieldValue = i
			}
		}
	}

	fieldParametersCache.Store(str, params)

	return params
}


func ShiftBytes(input []byte, k int) (output []byte) {
	length := len(input)
	output = make([]byte, length)
	if k >= 0 { //shift left
		nBytes := k >> 3
		k = k & 0x7
		if nBytes > length {
			return
		}
		for i := nBytes; i < length; i++ {
			if i == length-1 {
				output[i-nBytes] = input[i] << k
			} else {
				output[i-nBytes] = input[i]<<k | input[i+1]>>(8-k)
			}

		}
	} else { //shift right
		k = -k
		nBytes := k >> 3
		k = k & 0x7
		if nBytes > length {
			return
		}
		for i := length - 1; i >= nBytes; i-- {
			if i == nBytes {
				output[i] = input[i-nBytes] >> k
			} else {
				output[i] = input[i-nBytes]>>k | input[i-nBytes-1]<<(8-k)
			}
		}
	}
	return
}



func appendBytes(dstBytes []byte, srcBytes []byte) []byte{
	return append(dstBytes,srcBytes...)
}

func getConstrain(param fieldParameters) Constrain {
	if param.valueLowerBound != nil && param.valueUpperBound != nil {
		return Constrain{
			Lb: *param.valueLowerBound,
			Ub: *param.valueUpperBound,
		}
	}
	return Constrain{}
}

func (aw *aperWriter) handleConstraints(bitsLength uint64, extensive bool, 
    lowerBoundPtr *int64, upperBoundPtr *int64) (int64, int64, int64, error) {
    var lb, ub, sizeRange int64 = 0, -1, -1
    if upperBoundPtr != nil && lowerBoundPtr != nil {
        ub = *upperBoundPtr
        lb = *lowerBoundPtr
        if bitsLength <= uint64(ub) {
            sizeRange = ub - lb + 1
        }else if !extensive {
            return 0, 0, 0, aperError("Length is", ErrBound)
        }
        if extensive {
            if sizeRange == -1 {
                if err := aw.writeBool(One); err != nil {
                    return 0, 0, 0, err
                }
                lb = 0
            } else {
                if err := aw.writeBool(Zero); err != nil {
                    return 0, 0, 0, err
                }
            }
        }
    }else{
		sizeRange = -1
	}
    return lb, ub, sizeRange, nil
}

func parseConstraint(extensed bool, lowerBoundPtr *int64, upperBoundPtr *int64) (lb int64, ub int64, sizeRange int64) {
	lb, ub, sizeRange = 0, -1, -1
	if !extensed {
		if lowerBoundPtr != nil {
			lb = *lowerBoundPtr
		}
		if upperBoundPtr != nil {
			ub = *upperBoundPtr
			sizeRange = ub - lb + 1
		}
	}else{
		sizeRange = -1
	}
	if ub > 65535 {
		sizeRange = -1
	}
	return
}
