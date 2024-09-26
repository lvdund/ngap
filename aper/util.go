package aper

import (
	"bytes"

)

// shift byte array by a number of bits (positive for left, negative for right)
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

// Set a bit given its index in a byte array
func SetBit(content []byte, bitIndex uint) {
	//TODO:
}

// check if a bit at given index is set
func IsBitSet(content []byte, bitIndex uint) bool {
	//TODO:
	return false
}

func (aw *aperWriter) WriteFlush() error {
	return aw.flush()
}

func (w aperWriter) GetBuf() []byte {
	buf, ok := w.w.(*bytes.Buffer)
	if ok {
		return buf.Bytes()
	}
	return []byte("")
}
