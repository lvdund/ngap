package aper

import (
	"bytes"
	"fmt"
)

type AperMarshaller interface {
	AperEncode(w AperWriter) error
}

type AperUnmarshaller interface {
	AperDecode(w AperWriter) error
}

type BitString struct {
	Bytes   []byte
	NumBits uint64
}

type OctetString []byte

type Integer int64
type Enumerated int64

type Constrain struct {
	Lb int64
	Ub int64
}

func (c *Constrain) Range() uint64 {
	return uint64(c.Ub - c.Lb + 1)
}

func removeLastByte(aw *aperWriter) error {
    buf, ok := aw.w.(*bytes.Buffer)
    if !ok {
        return fmt.Errorf("writer is not a *bytes.Buffer")
    }
    currentBytes := buf.Bytes()
    if len(currentBytes) == 0 {
        return fmt.Errorf("buffer is empty, nothing to remove")
    }
    currentBytes = currentBytes[:len(currentBytes)-1]
    buf.Reset()
    _, err := buf.Write(currentBytes)
    if err != nil {
        return fmt.Errorf("error writing to buffer: %v", err)
    }
    return nil
}
func countSetBits(b byte) int {
	count := 0
	for b > 0 {
		count += int(b & 1) 
		b >>= 1             
	}
	return count
}


// Unmarshal parses the APER-encoded ASN.1 data structure b
// and uses the reflect package to fill in an arbitrary value pointed at by value.
// Because Unmarshal uses the reflect package, the structs
// being written to must use upper case field names.
//
// An ASN.1 INTEGER can be written to an int, int32, int64,
// If the encoded value does not fit in the Go type,
// Unmarshal returns a parse error.
//
// An ASN.1 BIT STRING can be written to a BitString.
//
// An ASN.1 OCTET STRING can be written to a []byte.
//
// An ASN.1 OBJECT IDENTIFIER can be written to an
// ObjectIdentifier.
//
// An ASN.1 ENUMERATED can be written to an Enumerated.
//
// Any of the above ASN.1 values can be written to an interface{}.
// The value stored in the interface has the corresponding Go type.
// For integers, that type is int64.
//
// An ASN.1 SEQUENCE OF x can be written
// to a slice if an x can be written to the slice's element type.
//
// An ASN.1 SEQUENCE can be written to a struct
// if each of the elements in the sequence can be
// written to the corresponding element in the struct.
//
// The following tags on struct fields have special meaning to Unmarshal:
//
//		optional        	OPTIONAL tag in SEQUENCE
//		sizeExt             specifies that size  is extensible
//		valueExt            specifies that value is extensible
//		sizeLB		        set the minimum value of size constraint
//		sizeUB              set the maximum value of value constraint
//		valueLB		        set the minimum value of size constraint
//		valueUB             set the maximum value of value constraint
//		default             sets the default value
//		openType            specifies the open Type
//	 referenceFieldName	the string of the reference field for this type (only if openType used)
//	 referenceFieldValue	the corresponding value of the reference field for this type (only if openType used)
//
// Other ASN.1 types are not supported; if it encounters them,
// Unmarshal returns a parse error.