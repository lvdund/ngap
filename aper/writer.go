package aper

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math/bits"
)

type AperWriter struct {
	*bitstreamWriter
}

func NewWriter(w io.Writer) *AperWriter {
	return &AperWriter{
		bitstreamWriter: NewBitStreamWriter(w),
	}
}

func (aw *AperWriter) Close() error {
	return aw.flush()
}

func (aw *AperWriter) writeBytes(bytes []byte) error {
	return aw.WriteBits(bytes, uint(8*len(bytes)))
}

func (aw *AperWriter) writeValue(v uint64, nbits uint) (err error) {
	defer func() {
		err = aperError("writeValue", err)
	}()

	if nbits > 64 {
		err = ErrUnderflow
		return
	}
	v = v << (64 - nbits)
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], v)
	err = aw.WriteBits(buf[:], nbits)
	return
}

func (aw *AperWriter) writeSemiConstraintWholeNumber(v uint64, lb uint64) (err error) {
	defer func() {
		err = aperError("writeSemiContrainWholeNumber", err)
	}()

	if lb > v {
		err = ErrUnderflow
		return
	}
	v -= lb
	length := (bits.Len64(v) + 7) >> 3
	if err = aw.align(); err != nil {
		return
	}
	//since length < 8, just write its value bits
	if err = aw.writeValue(uint64(length), 8); err != nil {
		return
	}
	//then write the value bits
	err = aw.writeValue(v, uint(length)*8)
	return
}

func (aw *AperWriter) writeNormallySmallNonNegativeValue(v uint64) (err error) {
	defer func() {
		err = aperError("writeNormallySmallNonNegativeValue", err)
	}()
	if v < POW_6 { //leading Zero to indicate a small value
		if err = aw.WriteBool(Zero); err != nil {
			return
		}
		err = aw.writeValue(v, 6)
		return
	} else { //leading One to indicate a whole number
		if err = aw.WriteBool(One); err != nil {
			return
		}
		//write as a semi constrained whole number with lower bound zero
		err = aw.writeSemiConstraintWholeNumber(v, 0)
	}
	return
}

func (aw *AperWriter) writeLength(r uint64, v uint64) (err error) {
	defer func() {
		err = aperError("writeLength", err)
	}()

	//if range is within 2 bytes, write value as a constrained value
	if r <= POW_16 && r > 0 {
		err = aw.writeConstraintValue(r, v)
		return
	}
	//otherwise range is zero or more than 2 bytes, consider as no range
	//align first
	if err = aw.align(); err != nil {
		return
	}

	if v < POW_7 { //<=7bits
		err = aw.writeValue(v, 8) //write as one byte with Zero leading
	} else if v < POW_14 { //<=14bits
		v |= 0x8000 //write as 16bits with One is leading
		err = aw.writeValue(v, 16)
	} else {
		//length value is multiple of POW_14
		v = (v >> 14) | 0xc0 //strip off last 14 bits, take one byte, add leading '11'
		err = aw.writeValue(v, 8)
	}
	return
}

func (aw *AperWriter) writeConstraintValue(r uint64, v uint64) (err error) {
	defer func() {
		err = aperError("writeConstraintValue", err)
	}()

	var nBytes uint
	if r < POW_8 { //range is smaller that one byte, write value bits, no alignment
		return aw.writeValue(v, uint(bits.Len64(r-1)))
	} else if r == POW_8 {
		nBytes = 1
	} else if r <= POW_16 {
		nBytes = 2
	} else {
		return ErrOverflow
	}
	//otherwise, align then write the value as whole bytes
	if err = aw.align(); err != nil {
		return
	}
	err = aw.writeValue(v, nBytes*8)
	return
}

func (aw *AperWriter) WriteString(content []byte, len uint64, c *Constraint, e bool, isBitstring bool) (err error) {
	lowerBound, lRange, _ := aw.writeExtBit(len, e, c)
	if lRange > 0 && uint64(c.Ub) >= POW_16 { //if upper bound is at lest 16bits then set as semi-constrain
		lRange = 0
	}
	if lRange == 1 { //constrain with fixed length; both bounds have the same value
		if int64(len) != lowerBound {
			err = ErrFixedLength
			return
		}
		var numByte, nbits uint64
		if isBitstring {
			numByte = (len + 7) >> 3
			nbits = len
		} else {
			numByte = len
			nbits = len * 8
		}
		if numByte > 2 { //if more than 2 bytes, align first
			if err = aw.align(); err != nil {
				return
			}
		}
		//then write content
		err = aw.WriteBits(content, uint(nbits))
		return
	}
	partReader := NewBitStreamReader(bytes.NewReader(content)) //for reading parts of content for writing
	totalLen := uint64(len) - uint64(lowerBound)
	var partLen uint64
	var partBytes []byte
	completed := false
	for {
		if totalLen > POW_16 {
			partLen = POW_16
		} else if totalLen >= POW_14 {
			partLen = totalLen & 0xc000 //strip last 14 bits, keep bit 14,15.
		} else {
			partLen = totalLen
			completed = true //last part to write
			//Last part can have zero length, still it must be encoded to tell
			//reader (decoder) to stop
		}
		totalLen -= partLen //reduce total length

		//encode length
		if err = aw.writeLength(uint64(lRange), partLen); err != nil {
			return
		}

		//write content part
		partLen += uint64(lowerBound)
		if partLen == 0 {
			return
		}

		//align last byte
		if err = aw.align(); err != nil {
			return
		}
		var partLenBits uint
		if !isBitstring {
			partLenBits = uint(partLen * 8)
		} else {
			partLenBits = uint(partLen)
		}
		if partBytes, err = partReader.ReadBits(partLenBits); err != nil { //get a content part to write
			return
		}
		if err = aw.WriteBits(partBytes, partLenBits); err != nil { //write the part
			return
		}
		if completed {
			break
		}
	}
	return
}

func (aw *AperWriter) WriteBitString(content []byte, nbits uint, c *Constraint, e bool) (err error) {
	defer func() {
		err = aperError("WriteBitString", err)
	}()
	err = aw.WriteString(content, uint64(nbits), c, e, true)
	return
}

func (aw *AperWriter) WriteOctetString(content []byte, c *Constraint, e bool) (err error) {
	defer func() {
		err = aperError("WriteOctetString", err)
	}()
	byteLen := uint64(len(content))
	err = aw.WriteString(content, byteLen, c, e, false)
	return
}

// constrain must have Lb <= Ub
func (aw *AperWriter) WriteEnumerate(v uint64, c Constraint, e bool) (err error) {
	defer func() {
		err = aperError("WriteEnumerate", err)
	}()

	if v <= uint64(c.Ub) { //value is in range
		if e {
			if err = aw.WriteBool(Zero); err != nil {
				return
			}
		}
		vRange := c.Range()
		if vRange > 1 {
			err = aw.writeConstraintValue(vRange, v-uint64(c.Lb))
			return
		}
		//in case Lb == Ub, no need to write value, when reading, just use the
		//bound value
	} else { //value is of of range
		if !e { //not extensible
			err = ErrInextensible
			return
		}

		if err = aw.WriteBool(One); err != nil {
			return
		}
		err = aw.writeNormallySmallNonNegativeValue(v - uint64(c.Ub) - 1)
	}

	return
}

func (aw *AperWriter) WriteOpenType(content []byte) (err error) {
	//it is just like writing an OctetString without a constraint and
	//extension bit
	if err = aw.WriteOctetString(content, nil, false); err != nil {
		return
	}
	err = aw.align()
	return
}

func (aw *AperWriter) WriteInteger(v int64, c *Constraint, e bool) (err error) {
	defer func() {
		err = aperError("WriteInteger", err)
	}()
	lb, sRange, _ := aw.writeExtBit(uint64(v), e, c)
	unsignedValue := uint64(v)
	var rawLength uint
	if sRange == 1 {
		return nil
	}

	if v < 0 {
		y := v >> 63
		unsignedValue = uint64(((v ^ y) - y)) - 1
	}
	if sRange <= 0 {
		unsignedValue >>= 7
	} else if sRange <= 65536 {
		return aw.writeConstraintValue(uint64(sRange), uint64(v-lb))
	} else {
		unsignedValue >>= 8
	}

	for rawLength = 1; rawLength <= 127; rawLength++ {
		if unsignedValue == 0 {
			break
		}
		unsignedValue >>= 8
	}
	// write length
	if sRange <= 0 {
		aw.align()
		_ = aw.writeBytes([]byte{byte(rawLength)})
	} else {
		unsignedValueRange := uint64(sRange - 1)
		bitLen := bits.Len64(unsignedValueRange)
		byteLen := uint((bitLen + 7) / 8)
		bitLenngth := bits.Len(uint(int(byteLen)))
		if err := aw.writeValue(uint64(rawLength-1), uint(bitLenngth)); err != nil {
			return err
		}
	}
	rawLength *= 8
	aw.align()
	if sRange < 0 {
		mask := int64(1<<rawLength - 1)
		return aw.writeValue(uint64(v&mask), rawLength)
	} else {
		v -= lb
		return aw.writeValue(uint64(v), rawLength)
	}
}

func (aw *AperWriter) WriteChoice(v uint64, uBound uint64, e bool) (err error) {
	defer func() {
		err = aperError("WriteChoice", err)
	}()
	if v < 1 {
		err = fmt.Errorf("Choice must be larger than 1")
		return
	}
	v -= 1
	if v > uBound {
		err = fmt.Errorf("Choice extension not supported")
		return
	}

	if e {
		if err = aw.WriteBool(Zero); err != nil {
			return
		}
	}
	err = aw.writeConstraintValue(uBound+1, v)
	return
}
