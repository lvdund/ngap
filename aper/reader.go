package aper

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math/bits"
)

type AperReader struct {
	*bitstreamReader
}

func NewReader(r io.Reader) *AperReader {
	return &AperReader{
		bitstreamReader: NewBitStreamReader(r),
	}
}

func (ar *AperReader) readBytes(nbytes uint) (output []byte, err error) {
	return ar.ReadBits(nbytes * 8)
}

func (ar *AperReader) readValue(nbits uint) (v uint64, err error) {
	defer func() {
		if err != nil {
			err = aperError("readValue", err)
		}
	}()

	if nbits > 64 {
		err = ErrOverflow
		return
	}
	var buf []byte
	if buf, err = ar.ReadBits(nbits); err != nil {
		return
	}
	vBytes := make([]byte, 8)
	copy(vBytes[:], buf)
	v = binary.BigEndian.Uint64(vBytes)
	v >>= (64 - nbits)
	return
}

func (ar *AperReader) readConstraintValue(r uint64) (v uint64, err error) {
	defer func() {
		err = aperError("readConstraintValue", err)
	}()

	var nBytes uint

	if r < POW_8 { //smaller than 1 byte, read value bits
		//fmt.Printf("Range for constraint=%d[%d]\n", bits.Len64(r), r)
		v, err = ar.readValue(uint(bits.Len64(r - 1)))
		return
	} else if r == POW_8 {
		nBytes = 1
	} else if r <= POW_16 {
		nBytes = 2
	} else {
		err = ErrOverflow
		return
	}
	//otherwise, align then read whole bytes (1 or 2)
	ar.align()
	v, err = ar.readValue(nBytes * 8)
	return
}

func (ar *AperReader) readSemiConstraintWholeNumber(lb uint64) (v uint64, err error) {
	defer func() {
		err = aperError("readSemiConstraintWholeNumber", err)
	}()

	ar.align()
	var length uint64

	if length, err = ar.readValue(8); err != nil {
		return
	}
	if v, err = ar.readValue(uint(length) * 8); err != nil {
		return
	}
	v += lb
	return
}

func (ar *AperReader) readNormallySmallNonNegativeValue() (v uint64, err error) {
	defer func() {
		err = aperError("readNormallySmallNonNegativeValue", err)
	}()

	var b bool
	if b, err = ar.ReadBool(); err != nil {
		return
	}
	if b {
		v, err = ar.readSemiConstraintWholeNumber(0)
	} else {
		v, err = ar.readValue(6)
	}
	return
}

// decode length of a data part in a multiple-parts content
func (ar *AperReader) readLength(lRange uint64) (value uint64, more bool, err error) {
	defer func() {
		err = aperError("readLength", err)
	}()

	more = false
	if lRange <= POW_16 && lRange > 0 { //range exist, read a contrained value
		if value, err = ar.readConstraintValue(lRange); err != nil {
		}
		return
	}

	//byte align
	ar.align()

	//detect the type of length then decode the value
	var first, second uint64
	if first, err = ar.readValue(8); err != nil { //read first byte for detecting type of encoded length
		err = aperError("read first byte", err)
		return
	}

	if (first & POW_7) == 0 { // first byte has leading Zero -> 7-bits value
		value = first & 0x7F
		return
	} else if (first & POW_6) == 0 { //  first byte has '10' leading bits -> 14bits value
		if second, err = ar.readValue(8); err != nil { //read second byte to calculate the length value
			err = aperError("read second byte", err)
			return
		}

		value = ((first & 63) << 8) | second //remove leading '10' bits then get the 14bits value
		return
	}

	//now handle the case where first byte has '11' leading bits, POW_14 <=
	//length <= POW_16; the length is a multipler of POW_14
	first &= 63                 //strip the '11' leading bits
	if first < 1 || first > 4 { //multipler of POW_14 must be in [1,4]
		err = ErrInvalidLength
		return
	}
	more = true //this is not last content part
	value = POW_14 * first
	return
}

func (ar *AperReader) ReadString(c *Constraint, e bool, isBitstring bool) (content []byte, err error) {
	defer func() {
		if isBitstring {
			err = aperError("ReadString BitString", err)
		}
		err = aperError("ReadString OctetString", err)
	}()
	lRange, lowerBound, err := ar.readExBit(c, e)
	if err != nil {
		return nil, err
	}

	if lRange > 0 && uint64(c.Ub) >= POW_16 { //if upper bound is at least 16 bits then set as semi-constrain
		lRange = 0
	}

	if lRange == 1 { //constrained with fixed length
		var numBytes, nbits uint
		if isBitstring {
			nbits = uint(c.Lb)
			numBytes = (nbits + 7) >> 3
		} else {
			numBytes = uint(c.Lb)
			nbits = numBytes * 8
		}
		if numBytes > 2 { //if more than 2 bytes, need align byte first
			ar.align()
		}
		content, err = ar.ReadBits(nbits)
		return
	}
	var buf bytes.Buffer
	var tmpBytes []byte
	partWriter := NewBitStreamWriter(&buf) //a bitstream writer to write parts of content
	more := true                           //more part to read
	var partLen uint64                     //length of a part to read
	for more {
		//read part length first
		if partLen, more, err = ar.readLength(lRange); err != nil {
			return
		}
		partLen += uint64(lowerBound)
		if partLen == 0 {
			//last part has zeros length, skip reading
			break
		}
		ar.align()
		//then read the  part content
		var partLenBits uint64
		if isBitstring {
			partLenBits = partLen
		} else {
			partLenBits = partLen * 8
		}
		if tmpBytes, err = ar.ReadBits(uint(partLenBits)); err != nil {
			return
		}
		//concat the part to the output bitstream
		if err = partWriter.WriteBits(tmpBytes, uint(partLenBits)); err != nil {
			return
		}
		//nbits += uint(partLen)
	}
	partWriter.flush()    //flush the buffer
	content = buf.Bytes() //return the concatenated output
	return
}

func (ar *AperReader) ReadBitString(c *Constraint, e bool) (content []byte, nbits uint, err error) {
	defer func() {
		err = aperError("ReadBitString", err)
	}()
	content, err = ar.ReadString(c, e, true)
	if err != nil {
		return
	}
	return content, 0, nil
}
func (ar *AperReader) ReadOctetString(c *Constraint, e bool) (content []byte, err error) {
	defer func() {
		err = aperError("ReadOctetString", err)
	}()
	content, err = ar.ReadString(c, e, false)
	if err != nil {
		return
	}
	return content, nil
}

func (ar *AperReader) ReadOpenType() (octets []byte, err error) {
	octets, err = ar.ReadOctetString(nil, false)
	ar.align()
	return
}

func (ar *AperReader) ReadInteger(c *Constraint, e bool) (value int64, err error) {
	defer func() {
		err = aperError("ReadInteger", err)
	}()

	sRange, _, err := ar.readExBit(c, e)

	if err != nil {
		return 0, err
	}
	var rawLength uint
	switch {
	case sRange == 1:
		value = c.Lb
		return
	case sRange == 0:
		ar.align()
		var tmp byte
		if tmp, err = ar.readByte(); err != nil {
			return
		}
		rawLength = uint(tmp)

	case uint64(sRange) <= POW_16:
		var tmp uint64
		if tmp, err = ar.readConstraintValue(uint64(sRange)); err != nil {
			return
		}
		value = int64(tmp) + c.Lb //c is non-nil
		return

	default: //sRange > POW_16, c is non-nil
		unsignedValueRange := uint64(sRange - 1)
		var byteLen uint
		for byteLen = 1; byteLen <= 127; byteLen++ {
			unsignedValueRange >>= 8
			if unsignedValueRange == 0 {
				break
			}
		}
		var bitLength, upper uint
		// 1 ~ 8 bits
		for bitLength = 1; bitLength <= 8; bitLength++ {
			upper = 1 << bitLength
			if upper >= byteLen {
				break
			}
		}
		var tmp uint64
		if tmp, err = ar.readValue(uint(bitLength)); err != nil {
			return
		}
		rawLength = uint(tmp) + 1
		ar.align()
	}

	var rawValue uint64
	if rawValue, err = ar.readValue(rawLength * 8); err != nil {
		return
	}
	if sRange == 0 { //unconstraint
		signedBitMask := uint64(1 << (rawLength*8 - 1))
		valueMask := signedBitMask - 1
		if rawValue&signedBitMask > 0 {
			value = int64((^rawValue)&valueMask+1) * -1
		} else {
			value = int64(rawValue)
		}
	} else { //with constraint
		value = int64(rawValue) + c.Lb //c is non-nil
	}
	return
}

// constrain must have Lb <= Ub
func (ar *AperReader) ReadEnumerate(c Constraint, e bool) (v uint64, err error) {
	defer func() {
		err = aperError("ReadEnumerate", err)
	}()

	if e { //if extensible is true, read the extention bit
		var exBit bool
		if exBit, err = ar.ReadBool(); err != nil {
			return
		}
		if exBit { //read out of range value
			var tmp uint64
			if tmp, err = ar.readNormallySmallNonNegativeValue(); err != nil {
				return
			}
			v = tmp + uint64(c.Ub) + 1 //adjust value with upper bound
			return
		}
	}
	//value is contrained
	if c.Range() > 1 {
		var tmp uint64
		if tmp, err = ar.readConstraintValue(c.Range()); err != nil {
			return
		}
		v = tmp + uint64(c.Lb) //adjust value with lower bound
	} else {
		v = uint64(c.Lb) //range is 1, use the bound
	}
	return
}

func (ar *AperReader) ReadChoice(uBound uint64, e bool) (v uint64, err error) {
	defer func() {
		err = aperError("ReadChoice", err)
	}()

	if e {
		var exBit bool
		if exBit, err = ar.ReadBool(); err != nil {
			return
		}
		if exBit {
			err = fmt.Errorf("Choice extension not supported")
			return
		}
	}
	var tmp uint64
	if tmp, err = ar.readConstraintValue(uBound + 1); err != nil {
		return
	}
	v = tmp + 1
	return
}
