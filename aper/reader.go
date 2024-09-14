package aper

import (
	"fmt"
	"io"
	"math/bits"
)
//==========================
type AperReader interface {
	ReadBool() (bool, error)
	ReadBits(uint) ([]byte, error)
	ReadOctetString(*Constrain, bool) ([]byte, error)
	ReadBitString(*Constrain, bool) ([]byte, uint, error)
	ReadEnumerate(*Constrain, bool) (uint64,  error)
	ReadInteger(*Constrain, bool) (int64, error)
}

type aperReader struct {
	r     io.Reader
	b     [1]byte //read buffer]
	byteOffset uint64
	index uint8   //number of read bits / index of the next bit to read [0:8]
}

func NewReader(r io.Reader) *aperReader {
	return &aperReader{
		r:     r,
		byteOffset: 0,
		index: 8, //indicate new buffer on next read
	}
}

func (ar *aperReader) ReadBool() (bool, error) {
	if ar.index == 8 { //read next byte to the buffer
		if _, err := ar.r.Read(ar.b[:]); err != nil && err != io.EOF {
			return Zero, err
		}
		ar.byteOffset +=1
		ar.index = 0
	}
	bitMask := uint8(1) << (7 - ar.index)
	d := ar.b[0] & bitMask
	ar.index++
	return d == bitMask, nil
}

// ReadByte reads a single byte from the stream
func (ar *aperReader) readByte() (byte, error) {
	v := ar.b[0] << ar.index
	if _, err := ar.r.Read(ar.b[:]); err != nil && err != io.EOF {
		ar.b[0] = 0
		return v, err
	}
	ar.byteOffset +=1

	v |= ar.b[0] >> (8 - ar.index)
	return v, nil
}

func (ar *aperReader) readBytes(nbytes uint) (output []byte, err error) {
	return ar.ReadBits(nbytes * 8)
}

// Bit alignment
func (ar *aperReader) readAlignBits() error {
	restartReaderPointer(ar.r)
	buf := make([]byte, uint(ar.byteOffset) + 1)
	if _, err := ar.r.Read(buf); err != nil {
		return nil
	}
	ar.b[0] = buf[ar.byteOffset]
	ar.byteOffset +=1
	ar.index = 0
	return nil
}

// read numbits to get value and update index
func (ar *aperReader) getValue(nbits uint) (value uint64, err error) {
    if nbits > 64 {
        return 0, fmt.Errorf("cannot read more than 64 bits into uint64")
    }
    buf, err := ar.ReadBits(nbits)
    if err != nil {
        return 0, err
    }
	for i, j := 0, nbits; j >= 8; i, j = i+1, j-8 {
		value <<= 8
		value |= uint64(uint(buf[i]))
	}
	if extraBits := nbits & 0x7; extraBits != 0 {
		shiftAmount := 8 - extraBits
		lastByte := buf[len(buf)-1] >> shiftAmount
		value >>= shiftAmount 
		value |= uint64(lastByte)
	}
    return value, nil
}

//extract value in case the data is limited by a low value (lb)
func (ar *aperReader) readConstraintValue(valueRange int64) (value uint64, err error) {
	// todo :
	var bytes uint
	if valueRange <= 255 {
		value, err = ar.getValue(uint(bits.Len64(uint64(valueRange-1))))
		return value, err
	} else if valueRange == 256 {
		bytes = 1
	} else if valueRange <= 65536 {
		bytes = 2
	} else {
		err = fmt.Errorf("Constraint Value is large than 65536")
		return value, err
	}
	if err = ar.readAlignBits(); err != nil {
		return value, err
	}
	value, err = ar.getValue(bytes * 8)
	return value, err
}


// read length of a data
func (ar *aperReader) readLength(sRange int64, repeat *bool) (value uint64, err error) {
	*repeat = false
	if sRange <= 65536 && sRange > 0 {
		return ar.readConstraintValue(sRange)
	}

	if err = ar.readAlignBits(); err != nil {
		return value, err
	}
	firstByte, err := ar.getValue(8)
	if err != nil {
		return value, err
	}
	if (firstByte & 128) == 0 { // #10.9.3.6
		value = firstByte & 0x7F
		return value, err
	} else if (firstByte & 64) == 0 { // #10.9.3.7
		var secondByte uint64
		if secondByte, err = ar.getValue(8); err != nil {
			return value, err
		}
		value = ((firstByte & 63) << 8) | secondByte
		return value, err
	}
	firstByte &= 63
	if firstByte < 1 || firstByte > 4 {
		err = fmt.Errorf("Parse Length Out of Constraint")
		return value, err
	}
	*repeat = true
	value = 16384 * firstByte
	return value, err
}


func (ar *aperReader) ReadBits(nbits uint) (output []byte, err error) {
	if nbits == 0 { //read nothing
		return
	}
	nOutputBytes := (nbits + 7) >> 3    //number of output bytes
	output = make([]byte, nOutputBytes) //prepare output
	//1. no need to read the next byte
	if nbits <= 8-uint(ar.index) { //smaller than number of remaining bits
		output[0] = ar.b[0] >> (8 - uint8(nbits) - ar.index) << (8 - uint8(nbits))
		ar.index += uint8(nbits)
		return
	}
	//2. must read some bytes
	offset := uint(ar.index)      //1 to 8
	output[0] = ar.b[0] << offset //consume remaining bits from the buffer
	//number of remaining bits to read: nbits - 8 + offset
	nReadBytes := (nbits + offset - 1) >> 3 //number of remaining bytes to read (at least 1)
	buf := make([]byte, uint(ar.byteOffset) + nReadBytes)
	//read all needed bytes
	restartReaderPointer(ar.r)
	if _, err = ar.r.Read(buf); err != nil && err != io.EOF {
		err = aperError("readBits", err)
		return
	}
	buf = buf[ar.byteOffset:]
	ar.byteOffset +=uint64(nReadBytes)
	//printReaderContents(ar.r)
	ar.b[0] = buf[ nReadBytes - 1] //last read byte to the buffer
	//determine the bit index after reading all bits
	if ar.index = uint8((nbits + offset - 8) & 0x07); ar.index == 0 {
		ar.index = 8
	}
	output[0] |= buf[0] >> (8 - offset) //complete the first output byte
	buf = ShiftBytes(buf, int(offset)) //shift left to remove consumed bits for aligning with the output
	//copy to the output
	if nOutputBytes > 1 {
		copy(output[1:], buf)
	}
	//truncate the last byte of the output if needs
	if numSpareBits := uint8(nbits & 0x07); numSpareBits > 0 {
		output[nOutputBytes-1] &= (1<<numSpareBits - 1) << (8 - numSpareBits)
	}
	return
}

func (ar *aperReader) ReadBitString(c *Constrain, e bool) (bs BitString, err error) {
	//TODO:
	if ar.index == 8{
		ar.readAlignBits()
	}
	var sRange int64 = -1
	//var bitLength uint64
	if !e {
		sRange = c.Ub - c.Lb + 1
	}
	if c.Ub > 65535 {
		sRange = -1
	}
	if sRange == 1 {
		sizes := (c.Ub + 7) >> 3
		bs.NumBits = uint64(c.Ub)
		if sizes >2 {
			if err := ar.readAlignBits(); err != nil {
				return bs, err
			}
			buf := make([]byte, sizes)
			if _, err = ar.r.Read(buf); err != nil && err != io.EOF {
				err = aperError("readBits", err)
				return
			}
			ar.byteOffset += uint64(sizes)
			bs.Bytes = buf
			ar.index = uint8(c.Ub & 0x7)
		}else{
			if buf , err := ar.ReadBits(uint(c.Ub)); err != nil{
				return bs , err
			}else{
				bs.Bytes = buf 
			}
		}
		return 
	}else{
		repeat := false
		for {
			var rawLength uint64
			if length, err := ar.readLength(sRange, &repeat); err != nil {
				return bs, err
			} else {
				rawLength = length
			}
			rawLength += uint64(c.Lb)
			if rawLength == 0 {
				return bs, nil
			}
			sizes := (rawLength + 7) >> 3
			if err := ar.readAlignBits(); err != nil {
				return bs, err
			}
			buf := make([]byte, sizes)
			if buffer,err := ar.readBytes(uint(sizes)); err != nil {
				return bs, err
			}else{
				copy(buf, buffer)
			}
			bs.Bytes = append(bs.Bytes, buf...)
			bs.NumBits += rawLength
			ar.index = uint8(rawLength & 0x7)
			if !repeat {
				break
			}
			if !repeat{
				break
			}
		}
	}
	return
}

func (ar *aperReader) ReadOctetString(c *Constrain, e bool) (octets []byte, err error) {
	//TODO:
	return
}

func (ar *aperReader) ReadInteger(c *Constrain, e bool) (value int64, err error) {
	//TODO:
	return
}

func (ar *aperReader) ReadEnumerate(c *Constrain, e bool) (value uint64, err error) {
	//TODO:
	return
}

/*
func (ar *aperReader) ReadSequenceOf(c *Constrain, e bool) (value uint64, err error) {
	//TODO:
	return
}
*/