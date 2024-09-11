package aper

import (
	"bytes"
	"fmt"
)

type AperMarshaller interface {
	AperEncode(w AperWriter) error
}

type AperUnmarshaller interface {
	AperDecode(w AperWriter) error
}

type BitString struct {
	Bytes   []byte
	NumBits uint64
}

type OctetString []byte

//type Integer int64

type Constrain struct {
	Lb int64
	Ub int64
}

func (c *Constrain) Range() uint64 {
	return uint64(c.Ub - c.Lb + 1)
}

func removeLastByte(aw *aperWriter) error {
    buf, ok := aw.w.(*bytes.Buffer)
    if !ok {
        return fmt.Errorf("writer is not a *bytes.Buffer")
    }
    currentBytes := buf.Bytes()
    if len(currentBytes) == 0 {
        return fmt.Errorf("buffer is empty, nothing to remove")
    }
    currentBytes = currentBytes[:len(currentBytes)-1]
    buf.Reset()
    _, err := buf.Write(currentBytes)
    if err != nil {
        return fmt.Errorf("error writing to buffer: %v", err)
    }
    return nil
}
func countSetBits(b byte) int {
	count := 0
	for b > 0 {
		count += int(b & 1) 
		b >>= 1             
	}
	return count
}