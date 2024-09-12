package aper

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	//	"math"

	//	"math"
	"math/bits"
)

const (
	Zero bool = false
	One  bool = true
)

type AperWriter interface {
	WriteBool(bool) error
	WriteBits([]byte, uint) error
	WriteOctetString([]byte, *Constrain, bool) error
	WriteBitString([]byte, uint, *Constrain, bool) error
	WriteInteger(int64, *Constrain, bool) error
	WriteEnumerate(uint64, Constrain, bool) error
	WriteBytes([]byte) error
	
}
type aperWriter struct {
	w     io.Writer
	b     [1]byte
	index uint8 //number o written bits in the buffer/index of the next bit to write [0:7]
}

func NewWriter(w io.Writer) *aperWriter {
	return &aperWriter{
		w:     w,
		index: 0,
	}
}

// write buffer and reset
func (aw *aperWriter) align() error {
	if aw.index > 0 {
		if _, err := aw.w.Write(aw.b[:]); err != nil {
			return err
		}
		aw.index = 0
		aw.b[0] = 0
	}
	return nil
}

func (aw *aperWriter) flush() error {
	fmt.Println("flush")
	if aw.index == 0 { //already flushed, no more write
		return nil
	}
	if _, err := aw.w.Write(aw.b[:]); err != nil {
		return err
	}
	aw.b[0] = 0
	aw.index = 0
	fmt.Println("=============================================================")
	return nil
}

// func (aw *aperWriter) WriteBool(bit bool) error {
// 	if bit {
// 		aw.b[0] |= 1 << (7 - aw.index)
// 	}

// 	aw.index++

// 	if aw.index == 8 {
// 		return aw.flush()
// 	}

// 	return nil
// }


func (aw *aperWriter) WriteBool(bit bool) error {
	var err error
	if bit {
		err = aw.writeValue(1, 1)
	} else {
		err = aw.writeValue(0, 1)
	}
	return err
}

// writes a single byte
func (aw *aperWriter) writeByte(v byte) error {
	aw.b[0] |= v >> aw.index

	if _, err := aw.w.Write(aw.b[:]); err != nil {
		return aperError("WriteByte", err)
	}
	aw.b[0] = v << (8 - aw.index)

	return nil
}

func (aw *aperWriter) WriteBytes(bytes []byte) error {
	err:=aw.WriteBits(bytes, uint(8*len(bytes)))
	return err
}

func (aw *aperWriter) WriteBits(content []byte, nbits uint) (error) {
	//var wb bool
	if nbits > uint(8*len(content)) {
		return aperError("writeBits", ErrUnderflow)
	}
	if nbits == 0 { //write nothing
		return nil
	}
	//truncate input
	content = content[0 : (nbits+7)>>3]
	//a. all bits can be fit on the current buffer
	if nbits <= 8-uint(aw.index) {
		aw.b[0] |= content[0] >> aw.index
		aw.index += uint8(nbits)
		if aw.index == 8 {
			return aw.flush()
		}
	}else{
		//b. need some writes
		nWriteBytes := (nbits + uint(aw.index) + 7) >> 3 //at lease two bytes
		buf := make([]byte, nWriteBytes)                 //buffer to keep all writes in byte alignment
		//fill the first byte
		
		buf[0] = aw.b[0] | content[0]>>aw.index
		//align the input byte for copying to the buffer array
		shiftBytes := ShiftBytes(content, 8-int(aw.index))
		copy(buf[1:], shiftBytes) //safe because buf is at least 2 bytes
		aw.index = uint8((nbits + uint(aw.index)) & 0x07)//determine new aw.index
		if aw.index == 0 {
			_, err := aw.w.Write(buf)
			if err != nil {
				return aperError("writeBits", err)
			}
			return nil 
		} else { //flush all except the last byte which is move to the buffer
			//removeLastByte(aw)
			_, err := aw.w.Write(buf[0 : nWriteBytes-1])
			if err != nil {
				return aperError("writeBits", err)
			}
			aw.b[0] = buf[nWriteBytes-1]
		}
		return nil 
	}
	return nil 
}


func (aw *aperWriter) writeValue(v uint64, nbits uint) error {
	if nbits > 64 {
		return aperError("writeValue", ErrUnderflow)
	}
	v = v << (64 - nbits)
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], v)
	err :=aw.WriteBits(buf[:], nbits)
	return err
}

func (aw *aperWriter) writeSemiConstrainWholeNumber(v uint64, lb uint64) (err error) {
	defer func() {
		err = aperError("writeSemiContrainWholeNumber", err)
	}()
	if lb > v {
		err = ErrUnderflow
		return
	}
	v -= lb
	length := (bits.Len64(v) + 7) >> 3

	if err = aw.writeLength(66537, uint64(length)); err != nil {
		return
	}
	err = aw.writeValue(v, uint(length)*8)
	return
}

func (aw *aperWriter) writeNormallySmallNonNegativeValue(v uint64) (err error) {
	defer func() {
		err = aperError("writeNormallySmallNonNegativeValue", err)
	}()
	if v < 64 {
		if err = aw.WriteBool(Zero); err != nil {
			return
		}
		err = aw.writeValue(v, 6)
		return
	} else {
		if err = aw.WriteBool(One); err != nil {
			return
		}
		err = aw.writeSemiConstrainWholeNumber(v, 0)
	}
	return
}

func (aw *aperWriter) writeLength(r uint64, v uint64) error {
	if r <= 65536 && r > 0 {
		return aw.writeConstrainValue(r, v)
	}
	if err := aw.align(); err != nil {
		return err
	}
	if v <= 127 {
		return aw.writeValue(v, 8)
	} else if v <= 16383 {
		v |= 0x8000
		return aw.writeValue(v, 16)
	}
	v = (v >> 14) | 0xc0
	return aw.writeValue(v, 8)
}

func (aw *aperWriter) writeConstrainValue(r uint64, v uint64) error {
	var nbytes uint
	if r <= 255 {
		return aw.writeValue(v, uint(bits.Len64(r)))
	} else if r == 256 {
		nbytes = 1
	} else if r <= 65536 {
		nbytes = 2
	} else {
		return aperError("writeConstrainValue", ErrOverflow)
	}
	if err := aw.align(); err != nil {
		return aperError("writeConstrainValue", err)
	}
	return aperError("writeConstrainValue", aw.writeValue(v, nbytes*8))
}

func (aw *aperWriter) WriteBitString(content []byte, nbits uint, c *Constrain, e bool) (err error) {
	// e - valueExtensible
	aww,_ := aw.w.(*bytes.Buffer)
	if aw.index != 0 {
		aw.b[0] = aww.Bytes()[len(aww.Bytes())-1]
		removeLastByte(aw)
	}
	var sRange int64 = -1
	var lb, ub int64 = 0, -1
	if c != nil {
		lb = c.Lb
		if c.Ub >= c.Lb {
			ub = c.Ub
			if nbits <=uint(c.Ub) {
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
	sizes := (nbits + 7) >> 3
	shift := (8 - nbits&0x7)
	if shift != 8 {
		content[sizes-1] &= (0xff << shift)
	}
	if sRange == 1 {
		if uint64(nbits) != uint64(ub) {
		}
		if sizes > 2 {
			aw.align()
			if _, err := aw.w.Write(content); err != nil {
				return aperError("writeBits", err)
			}
			aw.index = uint8(c.Ub & 0x7)
		} else {
			err = aw.WriteBits(content, uint(nbits))
			if aw.index != 0 {
				aw.w.Write(aw.b[:])
			}
		}
		return
	}
	rawLength := uint64(nbits) - uint64(lb)
	var byteOffset, partOfRawLength uint64
	for {
		if rawLength > 65536 {
			partOfRawLength = 65536
		} else if rawLength >= 16384 {
			partOfRawLength = rawLength & 0xc000
		} else {
			partOfRawLength = rawLength
		}
		if err = aw.writeLength(uint64(sRange), partOfRawLength); err != nil {
			return err
		}
		partOfRawLength += uint64(lb)
		sizes := (partOfRawLength + 7) >> 3
		if partOfRawLength == 0 {
			return err
		}
		aw.align()
		_,_ = aw.w.Write(content[byteOffset:byteOffset+sizes])
		rawLength -= (partOfRawLength - uint64(lb))
		if rawLength > 0 {
			byteOffset += sizes
		} else {
			aw.index += uint8(partOfRawLength & 0x7)
			// aw.align()
			fmt.Println(aw.index)
			break
		}
	}
	return err
}

func (aw *aperWriter) WriteOctetString(content []byte, c *Constrain, e bool) (err error) {
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
			fmt.Println("break")
			break
		}
	}
	return nil
}

func (aw *aperWriter) WriteSequenceOf(items []AperMarshaller, c *Constrain, e bool) error {
	//TODO:
	return nil
}

func (aw *aperWriter) WriteEnumerate(v uint64, c Constrain, e bool) error {
	if v <= uint64(c.Ub) {
		if e {
			if err := aw.WriteBool(Zero); err != nil {
				return aperError("WriteEnumerate", err)
			}
		}
		vRange := c.Range()
		if vRange > 1 {
			return aw.writeConstrainValue(vRange, v)
		}
	} else {
		if !e {
			return aperError("WriteEnumerate", ErrInextensible)
		}

		if err := aw.WriteBool(One); err != nil {
			return aperError("WriteEnumerate", err)
		}
		return aw.writeNormallySmallNonNegativeValue(v - uint64(c.Ub) - 1)
	}

	return nil
}

func (aw *aperWriter) WriteInteger(v int64, c *Constrain, e bool) (err error) {
	//TODO:
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
	
}