package aper

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/bits"
)

const (
	Zero bool = false
	One  bool = true
)

type AperWriter interface {
	WriteBool(bool) error
	WriteBits([]byte, uint64) error
	WriteOctetString([]byte, *Constrain, bool) error
	WriteBitString([]byte, uint, *Constrain, bool) error
	WriteInteger(uint64, *Constrain, bool) error
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
	if aw.index == 0 { //already flushed, no more write
		return nil
	}
	if _, err := aw.w.Write(aw.b[:]); err != nil {
		return err
	}
	aw.b[0] = 0
	aw.index = 0
	return nil
}

func (aw *aperWriter) WriteBool(bit bool) error {
	if bit {
		aw.b[0] |= 1 << (7 - aw.index)
	}

	aw.index++

	if aw.index == 8 {
		return aw.flush()
	}

	return nil
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

func (aw *aperWriter) writeBytes(bytes []byte) error {
	return aw.WriteBits(bytes, uint(8*len(bytes)))
}

// write 'nbits' from 'content' byte array
// putBitString
func (aw *aperWriter) WriteBits(content []byte, nbits uint) error {

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
		aw.b[0] |= (content[0] >> (8 - nbits)) << (8 - aw.index - uint8(nbits))
		aw.index += uint8(nbits)
		if aw.index == 8 {
			return aw.flush()
		}
		return nil
	}
	//b. need some writes
	nWriteBytes := (nbits + uint(aw.index) + 7) >> 3 //at lease two bytes
	buf := make([]byte, nWriteBytes)                 //buffer to keep all writes in byte alignment

	//fill the first byte
	buf[0] = aw.b[0] | content[0]>>aw.index

	//align the input byte for copying to the buffer array
	content = ShiftBytes(content, 8-int(aw.index))
	copy(buf[1:], content) //safe because buf is at least 2 bytes

	aw.index = uint8((nbits + uint(aw.index)) & 0x07) //determine new aw.index
	if aw.index == 0 {                                //flush all
		if _, err := aw.w.Write(buf); err != nil {
			return aperError("writeBits", err)
		}
		aw.b[0] = 0
	} else { //flush all except the last byte which is move to the buffer
		if _, err := aw.w.Write(buf[0 : nWriteBytes-1]); err != nil {
			return aperError("writeBits", err)
		}
		aw.b[0] = buf[nWriteBytes-1]
	}
	return nil
}
//putBitsValue
func (aw *aperWriter) writeValue(v uint64, nbits uint) error {
	if nbits > 64 {
		return aperError("writeValue", ErrUnderflow)
	}
	v = v << (64 - nbits)
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], v)
	return aw.WriteBits(buf[:], nbits)
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
	
	// var sRange int64 = -1
	// var lb, ub int64 = 0, 0
	// if c != nil {
	// 	lb = c.Lb
	// 	if c.Ub >= c.Lb {
	// 		ub = c.Ub
	// 		if nbits <= uint64(c.Ub) {
	// 			sRange = int64(c.Range())
	// 		} else if !e {
	// 			err = ErrInextensible
	// 			return
	// 		}
	// 		if e {
	// 			if sRange == -1 {
	// 				if err = aw.WriteBool(One); err != nil {
	// 					return
	// 				}
	// 				lb = 0
	// 			} else {
	// 				if err = aw.WriteBool(Zero); err != nil {
	// 					return
	// 				}
	// 			}
	// 		}

	// 	}
	// }

	// if ub > 65535 {
	// 	sRange = -1
	// }
	// sizes := (bitsLength + 7) >> 3
	// shift := (8 - bitsLength&0x7)
	// if shift != 8 {
	// 	bytes[sizes-1] &= (0xff << shift)
	// }

	// if sizeRange == 1 {
	// 	if bitsLength != uint64(ub) {
	// 		err = fmt.Errorf("bitString Length(%d) is not match fix-sized : %d", bitsLength, ub)
	// 	}
	// 	log.Debugf("Encoding BIT STRING size %d", ub)
	// 	if sizes > 2 {
	// 		pd.appendAlignBits()
	// 		pd.bytes = append(pd.bytes, bytes...)
	// 		pd.bitsOffset = uint(ub & 0x7)
	// 		log.Debugf("%s", perRawBitLog(bitsLength, len(pd.bytes), pd.bitsOffset, bytes))
	// 	} else {
	// 		err = pd.putBitString(bytes, uint(bitsLength))
	// 	}
	// 	log.Debugf("Encoded BIT STRING (length = %d): 0x%0x", bitsLength, bytes)
	// 	return
	// }
	// rawLength := bitsLength - uint64(lb)

	// var byteOffset, partOfRawLength uint64
	// for {
	// 	if rawLength > 65536 {
	// 		partOfRawLength = 65536
	// 	} else if rawLength >= 16384 {
	// 		partOfRawLength = rawLength & 0xc000
	// 	} else {
	// 		partOfRawLength = rawLength
	// 	}
	// 	if err = pd.appendLength(sizeRange, partOfRawLength); err != nil {
	// 		return
	// 	}
	// 	partOfRawLength += uint64(lb)
	// 	sizes := (partOfRawLength + 7) >> 3
	// 	log.Debugf("Encoding BIT STRING size %d", partOfRawLength)
	// 	if partOfRawLength == 0 {
	// 		return
	// 	}
	// 	pd.appendAlignBits()
	// 	pd.bytes = append(pd.bytes, bytes[byteOffset:byteOffset+sizes]...)
	// 	log.Debugf("%s", perRawBitLog(partOfRawLength, len(pd.bytes), pd.bitsOffset, bytes))
	// 	log.Debugf("Encoded BIT STRING (length = %d): 0x%0x", partOfRawLength,
	// 		bytes[byteOffset:byteOffset+sizes])
	// 	rawLength -= (partOfRawLength - uint64(lb))
	// 	if rawLength > 0 {
	// 		byteOffset += sizes
	// 	} else {
	// 		pd.bitsOffset += uint(partOfRawLength & 0x7)
	// 		// pd.appendAlignBits()
	// 		break
	// 	}
	// }

	//TODO:
	return nil
}

func (aw *aperWriter) WriteOctetString(content []byte, c *Constrain, e bool) error {
	//TODO:
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
	var  valueRange int64 = 0
	if c.Lb != 0 && c.Ub != 0{
		if v < c.Lb && v > c.Ub {
			fmt.Errorf("The integer value is out of the allowed range.")
		}
	}else{
		valueRange = c.Ub - c.Lb +1
	}

	fmt.Println(valueRange)

	return nil
}
