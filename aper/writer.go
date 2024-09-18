package aper

import (
	"bytes"
	"encoding/binary"
	//	"fmt"
	"io"
	"math/bits"
)

type AperWriter interface {
	WriteBool(bool) error
	WriteBits([]byte, uint) error
	WriteOctetString([]byte, *Constraint, bool) error
	WriteOpenType([]byte) error
	WriteBitString([]byte, uint, *Constraint, bool) error
	WriteInteger(uint64, *Constraint, bool) error
	WriteEnumerate(uint64, Constraint, bool) error
	Close() error
}

type aperWriter struct {
	*bitstreamWriter
}

func NewWriter(w io.Writer) *aperWriter {
	return &aperWriter{
		bitstreamWriter: NewBitStreamWriter(w),
	}
}

func (aw *aperWriter) Close() error {
	return aw.flush()
}

func (aw *aperWriter) writeBytes(bytes []byte) error {
	return aw.WriteBits(bytes, uint(8*len(bytes)))
}

func (aw *aperWriter) writeValue(v uint64, nbits uint) (err error) {
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

func (aw *aperWriter) writeSemiConstraintWholeNumber(v uint64, lb uint64) (err error) {
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

func (aw *aperWriter) writeNormallySmallNonNegativeValue(v uint64) (err error) {
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

func (aw *aperWriter) writeLength(r uint64, v uint64) (err error) {
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

func (aw *aperWriter) writeConstraintValue(r uint64, v uint64) (err error) {
	defer func() {
		err = aperError("writeConstraintValue", err)
	}()

	var nBytes uint
	if r < POW_8 { //range is smaller that one byte, write value bits, no alignment
		return aw.writeValue(v, uint(bits.Len64(r)))
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

func (aw *aperWriter) WriteBitString(content []byte, nbits uint, c *Constraint, e bool) (err error) {
	defer func() {
		err = aperError("WriteBitString", err)
	}()

	var lRange uint64 = 0    //length range
	var lowerBound int64 = 0 //length lower bound, default=0

	//write extension bit
	exBit := false
	if c != nil {
		if lowerBound = c.Lb; lowerBound < 0 { //make sure lower bound is not negative
			err = ErrConstraint
			return
		}
		if lRange = c.Range(); lRange > 0 && uint(c.Ub) < nbits { //unconstraint
			if !e { //unconstraint not supported
				err = ErrInextensible
				return
			} else {
				exBit = true
			}
		}
	}
	if e {
		if err = aw.WriteBool(exBit); err != nil {
			return
		}
	}

	if lRange > 0 && uint64(c.Ub) >= POW_16 { //if upper bound is at lest 16bits then set as semi-constrain
		lRange = 0
	}

	if lRange == 1 { //constrain with fixed length; both bounds have the same value
		if int64(nbits) != lowerBound {
			err = ErrFixedLength
			return
		}

		if numBytes := (nbits + 7) >> 3; numBytes > 2 { //if more than 2 bytes, align first
			if err = aw.align(); err != nil {
				return
			}
		}
		//then write content
		err = aw.WriteBits(content, nbits)
		return
	}
	partReader := NewBitStreamReader(bytes.NewReader(content)) //for reading parts of content for writing
	totalLen := uint64(nbits) - uint64(lowerBound)
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
		if err = aw.writeLength(lRange, partLen); err != nil {
			return
		}

		//align last byte
		if err = aw.align(); err != nil {
			return
		}
		//write content part
		partLen += uint64(lowerBound)
		if partBytes, err = partReader.ReadBits(uint(partLen)); err != nil { //get a content part to write
			return
		}
		if err = aw.WriteBits(partBytes, uint(partLen)); err != nil { //write the part
			return
		}
		if completed {
			break
		}
	}
	return
}

func (aw *aperWriter) WriteSequenceOf(items []AperMarshaller, c *Constraint, e bool) error {
	//TODO:
	return nil
}

// constrain must have Lb <= Ub
func (aw *aperWriter) WriteEnumerate(v uint64, c Constraint, e bool) (err error) {
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

func (aw *aperWriter) WriteOpenType(content []byte) (err error) {
	//it is just link writing an OctetString without a constraint and
	//extension bit
	return aw.WriteOctetString(content, nil, false)
}

func (aw *aperWriter) WriteOctetString(content []byte, c *Constraint, e bool) (err error) {
	//TODO: @Duc: It is very similar to WriteBitString
	defer func() {
		err = aperError("WriteOctetString", err)
	}()

	/*
		aww,_ := aw.w.(*bytes.Buffer)
		if aw.index != 0 {
			aw.b[0] = aww.Bytes()[len(aww.Bytes())-1]
			removeLastByte(aw)
		}
		byteLen := uint64(len(content))
		var lb, ub, sRange int64 = 0, -1, -1
		if c != nil {
			lb = c.Lb
			if c.Ub >= c.Lb {
				ub = c.Ub
				if int64(byteLen) <=c.Ub {
					sRange = int64(c.Range())
				} else if !e {
					err = ErrInextensible
					return
				}
				if e {
					if sRange == -1 {
						if err = aw.WriteBool(One); err != nil {
							return
						}
						lb = 0
					} else {
						if err = aw.WriteBool(Zero); err != nil {
							return
						}
					}
				}

			}
		}
		if ub > 65535 {
			sRange = -1
		}

		if sRange == 1 {
			if byteLen != uint64(ub) {
				err := fmt.Errorf("OctetString Length(%d) is not match fix-sized : %d", byteLen, ub)
				return err
			}
			if byteLen > 2 {
				aw.align()
				if _, err := aw.w.Write(content); err != nil {
					return aperError("writeBits", err)
				}
			} else {
				err:= aw.WriteBits(content, uint(byteLen*8))
				if aw.index != 0 {
					aw.w.Write(aw.b[:])
				}
				return err
			}
			return nil
		}

		rawLength := byteLen - uint64(lb)

		var byteOffset, partOfRawLength uint64
		for {
			if rawLength > 65536 {
				partOfRawLength = 65536
			} else if rawLength >= 16384 {
				partOfRawLength = rawLength & 0xc000
			} else {
				partOfRawLength = rawLength
			}
			if err := aw.writeLength(uint64(sRange), partOfRawLength); err != nil {
				return err
			}
			partOfRawLength += uint64(lb)
			if partOfRawLength == 0 {
				aw.w.Write(aw.b[:])
				return nil
			}
			aw.align()
			_,_ = aw.w.Write(content[byteOffset:byteOffset+partOfRawLength])
			rawLength -= (partOfRawLength - uint64(lb))
			if rawLength > 0 {
				byteOffset += partOfRawLength
			} else {
				// aw.align()
				break
			}
		}
	*/
	return
}

func (aw *aperWriter) WriteInteger(v int64, c *Constraint, e bool) (err error) {
	//TODO: @Duc: please check again
	defer func() {
		err = aperError("WriteInteger", err)
	}()

	/*
		var lb, sRange int64 = 0, -1
		if c != nil {
			lb = c.Lb
			if c.Ub >= c.Lb {
				if int64(v) <=c.Ub {
					sRange = int64(c.Range())
				} else if !e {
					err = ErrInextensible
					return
				}
				if e {
					if sRange == -1 {
						if err = aw.WriteBool(One); err != nil {
							return
						}
						lb = 0
					} else {
						if err = aw.WriteBool(Zero); err != nil {
							return
						}
					}
				}

			}
		}
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
			return aw.writeConstrainValue(uint64(sRange),uint64(v-lb))
		} else {
			unsignedValue >>= 8
		}
		for rawLength = 1; rawLength <= 127; rawLength++ {
			if unsignedValue == 0 {
				break
			}
			unsignedValue >>= 8
		}
		if sRange <= 0 {
			aw.align()
			_,_ = aw.w.Write([]byte{byte(rawLength)})
		} else {
			unsignedValueRange := uint64(sRange - 1)
			bitLength := bits.Len64(unsignedValueRange)
			byteLen := uint((bitLength + 7) / 8)
			bitLen := bits.Len(uint(int(byteLen)))
			if err := aw.writeValue(uint64(rawLength-1), uint(bitLen)); err != nil {
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
		return nil
	*/
	return
}
