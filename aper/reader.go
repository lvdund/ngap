package aper

import (
	"io"
)

type AperReader interface {
	ReadBool() (bool, error)
	ReadBits(uint64) ([]byte, error)
	ReadOctetString(*Constrain, bool) ([]byte, error)
	ReadBitString(*Constrain, bool) ([]byte, uint, error)
}

type aperReader struct {
	r     io.Reader
	b     [1]byte //read buffer
	index uint8   //number of read bits / index of the next bit to read [0:8]
}

func NewReader(r io.Reader) *aperReader {
	return &aperReader{
		r:     r,
		index: 8, //indicate new buffer on next read
	}
}

func (ar *aperReader) ReadBool() (bool, error) {
	if ar.index == 8 { //read next byte to the buffer
		if _, err := ar.r.Read(ar.b[:]); err != nil && err != io.EOF {
			return Zero, err
		}
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

	v |= ar.b[0] >> (8 - ar.index)
	return v, nil
}

func (ar *aperReader) readBytes(nbytes uint) (output []byte, err error) {
	return ar.ReadBits(nbytes * 8)
}

// Bit alignment
func (ar *aperReader) readAlignBits() error {
	return nil
}

// read numbits to get value and update index
func (ar *aperReader) getBitsValue(numBits uint) (value uint64, err error) {
	// TODO :
	return value, nil
}

//extract value in case the data is limited by a low value (lb)
func (ar *aperReader) readConstraintValue(valueRange int64) (value uint64, err error) {
	// todo :
	return value, err
}


// read length of a data
func (ar *aperReader) readLength(sizeRange int64, repeat *bool) (value uint64, err error) {
	// todo :
	return value, nil
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
	buf := make([]byte, nReadBytes)
	//read all needed bytes
	if _, err = ar.r.Read(buf); err != nil && err != io.EOF {
		err = aperError("readBits", err)
		return
	}
	ar.b[0] = buf[nReadBytes-1] //last read byte to the buffer
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


func (ar *aperReader) ReadBitString(c *Constrain, e bool) (bs []byte, nbits uint, err error) {
	//TODO:

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