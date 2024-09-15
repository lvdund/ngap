package aper

import "fmt"

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

func (w *aperWriter) WritePresent(present int, c *Constrain) error {
	var ub int64
	rawChoice := present - 1
	if c.Ub == 0 {
		return aperError("WritePresent", fmt.Errorf("The upper bound of Present is 0"))
	} else if ub = c.Ub; ub < 0 {
		return aperError("WritePresent", fmt.Errorf("The upper bound of Present is negative"))
	}

	if err := w.writeConstrainValue(uint64(ub+1), uint64(rawChoice)); err != nil {
		return err
	}
	return nil
}
