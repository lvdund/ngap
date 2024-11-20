package aper

import (
	"fmt"
	"io"
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
	byteIndex := bitIndex / 8
	bitPosition := bitIndex%8 - 1
	content[byteIndex] |= 1 << (7 - bitPosition)
}

// check if a bit at given index is set
func IsBitSet(content []byte, bitIndex uint) bool {
	byteIndex := bitIndex / 8
	bitPosition := bitIndex%8 - 1
	return (content[byteIndex] & (1 << (7 - bitPosition))) != 0
}

// GetBitString is to get BitString with desire size from source byte array with bit offset
func GetBitString(srcBytes []byte, bitsOffset uint, numBits uint) (dstBytes []byte, err error) {
	bitsLeft := uint(len(srcBytes))*8 - bitsOffset
	if numBits > bitsLeft {
		err = fmt.Errorf("Get bits overflow, requireBits: %d, leftBits: %d", numBits, bitsLeft)
		return
	}
	byteLen := (bitsOffset + numBits + 7) >> 3
	numBitsByteLen := (numBits + 7) >> 3
	dstBytes = make([]byte, numBitsByteLen)
	if numBitsByteLen == 0 {
		return
	}
	numBitsMask := byte(0xff)
	if modEight := numBits & 0x7; modEight != 0 {
		numBitsMask <<= uint8(8 - (modEight))
	}
	for i := 1; i < int(byteLen); i++ {
		dstBytes[i-1] = srcBytes[i-1]<<bitsOffset | srcBytes[i]>>(8-bitsOffset)
	}
	if byteLen == numBitsByteLen {
		dstBytes[byteLen-1] = srcBytes[byteLen-1] << bitsOffset
	}
	dstBytes[numBitsByteLen-1] &= numBitsMask
	return
}

func GetReader(r AperReader) []byte {
	t := r.bitstreamReader.r
	data, _ := io.ReadAll(t)
	return data
}
func GetWriter(w AperWriter) io.Writer {
	t := w.bitstreamWriter.w
	return t
}

func FlushWrite(w *AperWriter) error {
	return w.flush()
}
