package aper

import (
	"io"
)

const (
	Zero bool = false
	One  bool = true
)

/********** BITSTREAM WRTIER ***************/
type bitstreamWriter struct {
	w     io.Writer
	b     [1]byte
	index uint8 //number of written bits in the buffer/index of the next bit to write [0:7]
}

func NewBitStreamWriter(w io.Writer) *bitstreamWriter {
	return &bitstreamWriter{
		w:     w,
		index: 0,
	}
}

// write buffer and reset
func (bs *bitstreamWriter) align() error {
	if bs.index > 0 {
		shift := 8 - bs.index
		v := (bs.b[0] >> shift) << shift //set remaining bit to zeros
		if _, err := bs.w.Write([]byte{v}); err != nil {
			return err
		}
		bs.index = 0
		bs.b[0] = 0
	}
	return nil
}

func (bs *bitstreamWriter) flush() error {
	return bs.align()
	/*
		if bs.index == 0 { //already flushed, no more write
			return nil
		}
		if _, err := bs.w.Write(bs.b[:]); err != nil {
			return err
		}
		bs.b[0] = 0
		bs.index = 0
		return nil
	*/
}

func (bs *bitstreamWriter) WriteBool(bit bool) error {
	if bit {
		bs.b[0] |= 1 << (7 - bs.index)
	}

	bs.index++

	if bs.index == 8 {
		return bs.flush()
	}

	return nil
}

// writes a single byte
func (bs *bitstreamWriter) writeByte(v byte) error {
	bs.b[0] |= v >> bs.index

	if _, err := bs.w.Write(bs.b[:]); err != nil {
		return aperError("WriteByte", err)
	}
	bs.b[0] = v << (8 - bs.index)

	return nil
}

// write 'nbits' from 'content' byte array
func (bs *bitstreamWriter) WriteBits(content []byte, nbits uint) (err error) {
	defer func() {
		err = aperError("WriteBits", err)
	}()

	if nbits > uint(8*len(content)) {
		err = ErrUnderflow
		return
	}

	if nbits == 0 { //write nothing
		return
	}

	//truncate input
	nBytes := (nbits + 7) >> 3
	content = content[0:nBytes]
	nSpareBits := uint8(nbits & 0x07)
	if nSpareBits > 0 {
		tmp := content[nBytes-1]
		content[nBytes-1] = (tmp >> (8 - nSpareBits)) << (8 - nSpareBits)
	}

	//a. all bits can be fit on the current buffer
	if nbits <= 8-uint(bs.index) {
		bs.b[0] |= (content[0] >> (8 - nbits)) << (8 - bs.index - uint8(nbits))
		bs.index += uint8(nbits)
		if bs.index == 8 {
			err = bs.flush()
		}
		return
	}

	//b. need some writes
	nWriteBytes := (nbits + uint(bs.index) + 7) >> 3 //at lease two bytes
	buf := make([]byte, nWriteBytes)                 //buffer to keep all writes in byte alignment

	//fill the first byte
	buf[0] = bs.b[0] | content[0]>>bs.index

	//align the input byte for copying to the buffer array
	content = ShiftBytes(content, 8-int(bs.index))
	copy(buf[1:], content) //safe because buf is at least 2 bytes

	bs.index = uint8((nbits + uint(bs.index)) & 0x07) //determine new bs.index
	if bs.index == 0 {                                //flush all
		if _, err = bs.w.Write(buf); err != nil {
			return
		}
		bs.b[0] = 0
	} else { //flush all except the last byte which is move to the buffer
		if _, err = bs.w.Write(buf[0 : nWriteBytes-1]); err != nil {
			return
		}
		bs.b[0] = buf[nWriteBytes-1]
	}
	return
}

/********** BITSTREAM READER ***************/
type bitstreamReader struct {
	r     io.Reader
	b     [1]byte //read buffer
	index uint8   //number of read bits / index of the next bit to read [0:8]
}

func NewBitStreamReader(r io.Reader) *bitstreamReader {
	return &bitstreamReader{
		r:     r,
		index: 8, //indicate new buffer on next read
	}
}

func (bs *bitstreamReader) ReadBool() (bool, error) {
	if bs.index == 8 { //read next byte to the buffer
		if _, err := bs.r.Read(bs.b[:]); err != nil && err != io.EOF {
			return Zero, err
		}
		bs.index = 0
	}
	bitMask := uint8(1) << (7 - bs.index)
	d := bs.b[0] & bitMask
	bs.index++
	return d == bitMask, nil
}

func (bs *bitstreamReader) ReadBits(nbits uint) (output []byte, err error) {
	defer func() {
		err = aperError("ReadBits", err)
	}()

	if nbits == 0 { //read nothing
		return
	}

	nOutputBytes := (nbits + 7) >> 3    //number of output bytes
	output = make([]byte, nOutputBytes) //prepare output

	//1. no need to read the next byte
	if nbits <= 8-uint(bs.index) { //smaller than number of remaining bits
		output[0] = bs.b[0] >> (8 - uint8(nbits) - bs.index) << (8 - uint8(nbits))
		bs.index += uint8(nbits)
		return
	}

	//2. must read some bytes
	offset := uint(bs.index)      //1 to 8
	output[0] = bs.b[0] << offset //consume remaining bits from the buffer

	//number of remaining bits to read: nbits - 8 + offset
	nReadBytes := (nbits + offset - 1) >> 3 //number of remaining bytes to read (at least 1)
	buf := make([]byte, nReadBytes)
	//read all needed bytes
	if _, err = bs.r.Read(buf); err != nil /*&& err != io.EOF*/ {
		return
	}

	bs.b[0] = buf[nReadBytes-1] //last read byte to the buffer
	//determine the bit index after reading all bits
	if bs.index = uint8((nbits + offset - 8) & 0x07); bs.index == 0 {
		bs.index = 8
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

func (bs *bitstreamReader) align() {
	bs.index = 8
}

// ReadByte reads a single byte from the stream
func (bs *bitstreamReader) readByte() (byte, error) {
	v := bs.b[0] << bs.index

	if _, err := bs.r.Read(bs.b[:]); err != nil && err != io.EOF {
		bs.b[0] = 0
		return v, err
	}

	v |= bs.b[0] >> (8 - bs.index)

	return v, nil
}
