package aper

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_Aper(t *testing.T) {
	buf := new(bytes.Buffer)
	w := NewWriter(buf)
	if err := w.WriteBool(One); err != nil {
		t.Fatalf("WriteBit 1 error %+v", err)
	}
	b1 := uint8(0xf0)
	if err := w.writeByte(b1); err != nil {
		t.Fatalf("WriteByte 1 error %+v", err)
	}
	if err := w.WriteBool(One); err != nil {
		t.Fatalf("WriteBit 2 error %+v", err)
	}
	b2 := uint8(0x0f)
	if err := w.writeByte(b2); err != nil {
		t.Fatalf("WriteByte 2 error %+v", err)
	}
	content := []byte{0xf0, 0xff, 0x0f}
	nbits := uint(22)
	for i := 0; i < 5; i++ {
		if err := w.WriteBits(content, nbits); err != nil {
			t.Fatalf("WriteBits error %+v", err)
		}
	}

	if err := w.flush(); err != nil {
		t.Fatalf("Flush error %+v", err)
	}
	fmt.Printf("%.8b\n", buf.Bytes())

	r := NewReader(bytes.NewBuffer(buf.Bytes()))
	if b, err := r.ReadBool(); err != nil || !b {
		t.Fatalf("ReadBit 1 error: %+v-%v", err, b)
	}

	if v, err := r.readByte(); err != nil || v != b1 {
		t.Fatalf("ReadByte 1 error: %+v-%.8b", err, v)
	}

	if b, err := r.ReadBool(); err != nil || !b {
		t.Fatalf("ReadBit 1 error: %+v-%v", err, b)
	}
	if v, err := r.readByte(); err != nil || v != b2 {
		t.Fatalf("ReadByte 2 error: %+v-%.8b", err, v)
	}
	for i := 0; i < 5; i++ {
		if newContent, err := r.ReadBits(nbits); err != nil {
			t.Fatalf("ReadBits error: %+v", err)
		} else {
			fmt.Printf("%.8b\n", newContent)
		}
		/*
			if _, err := r.readBits(4); err != nil {
				t.Fatalf("ReadBits error: %+v", err)
			}
		*/
	}
}
