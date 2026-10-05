package aper

import (
	"bytes"
	"testing"
)

// bits6 is a 6-bit constrained value, so a list of them ends mid-octet.
type bits6 struct{ v uint64 }

func (b *bits6) Encode(w *AperWriter) error { return w.writeConstraintValue(64, b.v) }
func (b *bits6) Decode(r *AperReader) (err error) {
	b.v, err = r.readConstraintValue(64)
	return
}

// A SEQUENCE OF is followed directly by the next field, with no padding: the
// reference is free5gc/aper, which writes the same octets.
func TestSequenceOfIsNotPaddedAtItsEnd(t *testing.T) {
	var buf bytes.Buffer
	w := NewWriter(&buf)
	items := []*bits6{{1}, {2}}
	if err := WriteSequenceOf(items, w, &Constraint{Lb: 1, Ub: 64}, false); err != nil {
		t.Fatal(err)
	}
	if err := w.writeValue(0x5, 4); err != nil {
		t.Fatal(err)
	}
	w.Close()
	// 000001 (count-1) 000001 000010 (items) 0101 (next field), then the pad
	// Close writes: 22 bits in three octets.
	if want := []byte{0x04, 0x10, 0x94}; !bytes.Equal(buf.Bytes(), want) {
		t.Fatalf("encoded %x, want %x", buf.Bytes(), want)
	}

	r := NewReader(bytes.NewBuffer(buf.Bytes()))
	got, err := ReadSequenceOfEx(func() *bits6 { return new(bits6) }, r, &Constraint{Lb: 1, Ub: 64}, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0].v != 1 || got[1].v != 2 {
		t.Fatalf("decoded %+v", got)
	}
	if next, err := r.readValue(4); err != nil || next != 0x5 {
		t.Fatalf("field after the list read as %x (%v), want 5", next, err)
	}
}
