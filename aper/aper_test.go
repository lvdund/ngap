package aper

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func Test_Choice(t *testing.T) {
	buf := new(bytes.Buffer)
	w := NewWriter(buf)
	if err := w.WriteChoice(7, 7, false); err != nil {
		t.Fatalf("WriteChoice error %+v", err)
		return
	}
	if err := w.flush(); err != nil {
		t.Fatalf("Flush error %+v", err)
	}
	fmt.Printf("== %.8b %v\n", buf.Bytes(), buf.Bytes())
	r := NewReader(buf)
	if c, err := r.ReadChoice(7, false); err != nil {
		t.Fatalf("ReadChoice error %+v", err)
	} else {
		fmt.Println(c)
	}
}

func Test_WriteInteger(t *testing.T) {
	buf := new(bytes.Buffer)
	w := NewWriter(buf)
	if err := w.WriteInteger(12, &Constraint{Lb: 1, Ub: 20}, false); err != nil {
		fmt.Println("err write")
		return
	}
	if err := w.flush(); err != nil {
		t.Fatalf("Flush error %+v", err)
	}
	fmt.Printf("== %.8b %v\n", buf.Bytes(), buf.Bytes())
	buf1 := new(bytes.Buffer)
	w1 := NewWriter(buf1)
	if err := w1.WriteInteger(12, &Constraint{Lb: 1, Ub: 19}, false); err != nil {
		fmt.Println("err write")
		return
	}
	if err := w1.flush(); err != nil {
		t.Fatalf("Flush error %+v", err)
	}
	fmt.Printf("== %.8b %v", buf1.Bytes(), buf1.Bytes())
	// r := NewReader(buf)

}

func Test_Aper(t *testing.T) {
	buf := new(bytes.Buffer)
	w := NewWriter(buf)
	if err := w.writeValue(100, 12); err != nil {
		t.Fatalf("writeValue error %+v", err)
	}

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
	if err := w.writeConstraintValue(1024, 987); err != nil {
		t.Fatalf("writeValue error %+v", err)
	}

	if err := w.writeSemiConstraintWholeNumber(1111, 500); err != nil {
		t.Fatalf("writeValue error %+v", err)
	}
	if err := w.writeNormallySmallNonNegativeValue(21); err != nil {
		t.Fatalf("writeValue error %+v", err)
	}
	if err := w.writeNormallySmallNonNegativeValue(1984); err != nil {
		t.Fatalf("writeValue error %+v", err)
	}

	content := []byte{0xf0, 0xff, 0x0f}
	nbits := uint(22)
	for i := 0; i < 5; i++ {
		if err := w.WriteBits(content, nbits); err != nil {
			t.Fatalf("WriteBits error %+v", err)
		}
	}
	if err := w.writeValue(1023, 20); err != nil {
		t.Fatalf("writeValue error %+v", err)
	}
	if err := w.writeConstraintValue(50, 20); err != nil {
		t.Fatalf("writeValue error %+v", err)
	}
	if err := w.writeConstraintValue(256, 100); err != nil {
		t.Fatalf("writeValue error %+v", err)
	}

	if err := w.writeSemiConstraintWholeNumber(9999, 6666); err != nil {
		t.Fatalf("writeValue error %+v", err)
	}
	if err := w.WriteEnumerate(120, Constraint{
		Lb: 0, Ub: 2,
	}, true); err != nil {
		t.Fatalf("WriteEnumerate error %+v", err)
	}
	if err := w.WriteChoice(1, 2, false); err != nil {
		t.Fatalf("WriteChoice error %+v", err)
	}

	if err := w.flush(); err != nil {
		t.Fatalf("Flush error %+v", err)
	}

	//fmt.Printf("%.8b\n", buf.Bytes())

	r := NewReader(buf)
	if v, err := r.readValue(12); err != nil || v != 100 {
		t.Fatalf("readValue error: %+v", err)
	} else {
		fmt.Printf("read value %d\n", v)
	}

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
	if v, err := r.readConstraintValue(1024); err != nil || v != 987 {
		t.Fatalf("readConstraintValue error %+v", err)
	} else {
		fmt.Printf("read constrain value %d\n", v)
	}
	if v, err := r.readSemiConstraintWholeNumber(500); err != nil || v != 1111 {
		t.Fatalf("readSemiConstraintWholeNumber error %+v", err)
	} else {
		fmt.Printf("read semi constrain whole number value %d\n", v)
	}
	if v, err := r.readNormallySmallNonNegativeValue(); err != nil || v != 21 {
		t.Fatalf("readNormallySmallNonNegativeValue error %+v", err)
	} else {
		fmt.Printf("readNormallySmallNonNegativeValue %d\n", v)
	}
	if v, err := r.readNormallySmallNonNegativeValue(); err != nil || v != 1984 {
		t.Fatalf("readNormallySmallNonNegativeValue error %+v", err)
	} else {
		fmt.Printf("readNormallySmallNonNegativeValue %d\n", v)
	}
	for i := 0; i < 5; i++ {
		if newContent, err := r.ReadBits(nbits); err != nil {
			t.Fatalf("ReadBits error: %+v", err)
		} else {
			_ = newContent
			//fmt.Printf("%.8b\n", newContent)
		}
	}
	if v, err := r.readValue(20); err != nil || v != 1023 {
		t.Fatalf("readValue error: %+v", err)
	} else {
		fmt.Printf("read value %d\n", v)
	}
	if v, err := r.readConstraintValue(50); err != nil || v != 20 {
		t.Fatalf("writeValue error %+v", err)
	} else {
		fmt.Printf("read constrain value %d\n", v)
	}
	if v, err := r.readConstraintValue(256); err != nil || v != 100 {
		t.Fatalf("writeValue error %+v", err)
	} else {
		fmt.Printf("read constrain value %d\n", v)
	}
	if v, err := r.readSemiConstraintWholeNumber(6666); err != nil || v != 9999 {
		t.Fatalf("readSemiConstraintWholeNumber error %+v", err)
	} else {
		fmt.Printf("read semi constrain whole number value %d\n", v)
	}
	if v, err := r.ReadEnumerate(Constraint{
		Lb: 0,
		Ub: 2,
	}, true); err != nil || v != 120 {
		t.Fatalf("ReadEnumerate error %+v", err)
	} else {
		fmt.Printf("ReadEnumerate %d\n", v)
	}
	if v, err := r.ReadChoice(2, false); err != nil || v != 1 {
		t.Fatalf("ReadChoice error %+v", err)
	} else {
		fmt.Printf("ReadChoice %d\n", v)
	}

}

func TestWriteReadBitStringGroups(t *testing.T) {
	testGroups := []struct {
		name  string
		tests []struct {
			name       string
			content    []byte
			nbits      uint
			constraint *Constraint
			extensible bool
		}
		expected []byte // Expected value after writing all tests in the group
	}{
		{
			name: "Group 1",
			tests: []struct {
				name       string
				content    []byte
				nbits      uint
				constraint *Constraint
				extensible bool
			}{
				{
					content:    []byte{0xa0},
					nbits:      3,
					constraint: &Constraint{Lb: 3, Ub: 3},
				},
				{
					content:    []byte{0xfe},
					nbits:      8,
					constraint: &Constraint{Lb: 0, Ub: 125},
				},
				{
					content:    []byte{0xec},
					nbits:      6,
					constraint: &Constraint{Lb: 0, Ub: 255},
				},
				{
					content:    []byte{0xd8},
					nbits:      5,
					constraint: &Constraint{Lb: 0, Ub: 555},
				},
			},
			expected: []byte{0xA2, 0x00, 0xFE, 0x06, 0xEC, 0x00, 0x05, 0xD8},
		},
		{
			name: "Group 2",
			tests: []struct {
				name       string
				content    []byte
				nbits      uint
				constraint *Constraint
				extensible bool
			}{
				{
					content:    []byte{0xa0},
					nbits:      3,
					constraint: &Constraint{Lb: 3, Ub: 3},
				},
				{
					content:    []byte{0xa0},
					nbits:      3,
					constraint: &Constraint{Lb: 3, Ub: 3},
				},
			},
			expected: []byte{180},
		},
		{
			name: "Group 3",
			tests: []struct {
				name       string
				content    []byte
				nbits      uint
				constraint *Constraint
				extensible bool
			}{
				{
					content:    []byte{0xa0},
					nbits:      3,
					constraint: &Constraint{Lb: 3, Ub: 3},
				},
				{
					content:    []byte{0xb0},
					nbits:      4,
					constraint: &Constraint{Lb: 4, Ub: 4},
				},
			},
			expected: []byte{182},
		},
	}

	for _, group := range testGroups {
		t.Run(group.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			writer := NewWriter(buf)

			// Test writing
			for _, tt := range group.tests {
				t.Run("Write: "+tt.name, func(t *testing.T) {
					_ = writer.WriteBitString(tt.content, tt.nbits, tt.constraint, tt.extensible)
				})
			}
			writer.flush()
			if !bytes.Equal(buf.Bytes(), group.expected) {
				t.Errorf("Final buffer after write = %v, want %v", buf.Bytes(), group.expected)
			}

			// Test reading
			reader := NewReader(bytes.NewReader(buf.Bytes()))
			for _, tt := range group.tests {
				t.Run("Read: "+tt.name, func(t *testing.T) {
					decodedBytes, _, err := reader.ReadBitString(tt.constraint, tt.extensible)
					if err != nil {
						t.Errorf("Error decoding: %v", err)
					}
					if !bytes.Equal(decodedBytes, tt.content) {
						t.Errorf("Decoded bitstring = %v, want %v", decodedBytes, tt.content)
					}
				})
			}
		})
	}
}

func TestOctetStringEncodeDecode(t *testing.T) {
	testGroups := []struct {
		name  string
		tests []struct {
			name       string
			content    []byte
			constraint *Constraint
			extensible bool
			expected   []byte
		}
		expectedBytes []byte // Expected value after encoding
	}{
		{
			name: "Group 1",
			tests: []struct {
				name       string
				content    []byte
				constraint *Constraint
				extensible bool
				expected   []byte
			}{
				{
					content:    []byte("a"),
					constraint: &Constraint{Lb: 1, Ub: 1},
					extensible: true,
					expected:   []byte("a"),
				},
				{
					content:    []byte("bcd"),
					constraint: &Constraint{Lb: 0, Ub: 20},
					extensible: false,
					expected:   []byte("bcd"),
				},
				{
					content:    []byte("ef"),
					constraint: &Constraint{Lb: 2, Ub: 2},
					extensible: false,
					expected:   []byte("ef"),
				},
			},
			expectedBytes: []byte("\x30\x8Cbcdef"),
		},
		{
			name: "Group 2",
			tests: []struct {
				name       string
				content    []byte
				constraint *Constraint
				extensible bool
				expected   []byte
			}{
				{
					content:    []byte("a"),
					constraint: &Constraint{Lb: 1, Ub: 1},
					extensible: true,
					expected:   []byte("a"),
				},
				{
					content:    []byte("abcdefgh"),
					constraint: &Constraint{Lb: 0, Ub: 20},
					extensible: false,
					expected:   []byte("abcdefgh"),
				},
				{
					content:    []byte("ij"),
					constraint: &Constraint{Lb: 2, Ub: 2},
					extensible: false,
					expected:   []byte("ij"),
				},
			},
			expectedBytes: []byte("\x30\xA0abcdefghij"),
		},
		{
			name: "Group 3",
			tests: []struct {
				name       string
				content    []byte
				constraint *Constraint
				extensible bool
				expected   []byte
			}{
				{
					content:    []byte("a"),
					constraint: &Constraint{Lb: 1, Ub: 1},
					extensible: true,
					expected:   []byte("a"),
				},
				{
					content:    []byte(""),
					constraint: &Constraint{Lb: 0, Ub: 20},
					extensible: false,
					expected:   []byte(""),
				},
				{
					content:    []byte("bc"),
					constraint: &Constraint{Lb: 2, Ub: 2},
					extensible: false,
					expected:   []byte("bc"),
				},
			},
			expectedBytes: []byte("\x30\x81\x89\x8C"),
		},
	}

	for _, group := range testGroups {
		t.Run(group.name, func(t *testing.T) {
			// 1. Encode octet strings
			buf := new(bytes.Buffer)
			writer := NewWriter(buf)
			for _, tt := range group.tests {
				t.Run(tt.name, func(t *testing.T) {
					err := writer.WriteOctetString(tt.content, tt.constraint, tt.extensible)
					if err != nil {
						t.Errorf("WriteOctetString() error = %v", err)
					}
				})
			}
			writer.flush()

			// 2. Verify the encoded result
			if !bytes.Equal(buf.Bytes(), group.expectedBytes) {
				t.Errorf("Encoded buffer = %v, want %v", buf.Bytes(), group.expectedBytes)
			}

			// 3. Decode and compare with expected values
			reader := NewReader(bytes.NewReader(buf.Bytes()))
			for _, tt := range group.tests {
				t.Run(tt.name, func(t *testing.T) {
					decoded, err := reader.ReadOctetString(tt.constraint, tt.extensible)
					if err != nil {
						t.Errorf("ReadOctetString() error = %v", err)
					}
					if !bytes.Equal(decoded, tt.expected) {
						t.Errorf("Decoded octet string = %v, want %v", decoded, tt.expected)
					}
				})
			}
		})
	}
}

func TestIntegerEncodeDecode(t *testing.T) {
	testGroups := []struct {
		name  string
		tests []struct {
			value      int64
			constraint *Constraint
			extensible bool
			expected   int64
		}
		expectedBytes []byte // Expected bytes after encoding all tests in the group
	}{
		{
			name: "Group 1",
			tests: []struct {
				value      int64
				constraint *Constraint
				extensible bool
				expected   int64
			}{
				{
					value:      45,
					constraint: &Constraint{Lb: 1, Ub: 110},
					extensible: false,
					expected:   45,
				},
				{
					value:      123,
					constraint: &Constraint{Lb: 0, Ub: 255},
					extensible: false,
					expected:   123,
				},
				{
					value:      6445,
					constraint: &Constraint{Lb: 0, Ub: 255},
					extensible: true,
					expected:   6445,
				},
			},
			expectedBytes: []byte{0x58, 0x7B, 0x80, 0x02, 0x19, 0x2D},
		},
	}

	for _, group := range testGroups {
		t.Run(group.name, func(t *testing.T) {
			// 1. Encode integers
			buf := new(bytes.Buffer)
			writer := NewWriter(buf)
			for _, tt := range group.tests {
				t.Run(fmt.Sprintf("Value=%d", tt.value), func(t *testing.T) {
					err := writer.WriteInteger(tt.value, tt.constraint, tt.extensible)
					if err != nil {
						t.Errorf("WriteInteger() error = %v", err)
					}
				})
			}
			writer.flush()

			// 2. Verify the encoded result
			if !bytes.Equal(buf.Bytes(), group.expectedBytes) {
				t.Errorf("Encoded buffer = %v, want %v", buf.Bytes(), group.expectedBytes)
			}

			// 3. Decode and compare with expected values
			reader := NewReader(bytes.NewReader(buf.Bytes()))
			for _, tt := range group.tests {
				t.Run(fmt.Sprintf("Expected=%d", tt.expected), func(t *testing.T) {
					decoded, err := reader.ReadInteger(tt.constraint, tt.extensible)
					if err != nil {
						t.Errorf("ReadInteger() error = %v", err)
					}
					if decoded != tt.expected {
						t.Errorf("Decoded value = %d, want %d", decoded, tt.expected)
					}
				})
			}
		})
	}
}
func TestEnumerateEncodeDecode(t *testing.T) {
	testGroups := []struct {
		name  string
		tests []struct {
			value      uint64
			constraint *Constraint
			extensible bool
			expected   uint64
		}
		expectedBytes []byte // Expected bytes after encoding all tests in the group
	}{
		{
			name: "Group 1 - Definite Range",
			tests: []struct {
				value      uint64
				constraint *Constraint
				extensible bool
				expected   uint64
			}{
				{
					value:      0,
					constraint: &Constraint{Lb: 0, Ub: 3},
					extensible: false,
					expected:   0,
				},
				{
					value:      1,
					constraint: &Constraint{Lb: 0, Ub: 3},
					extensible: false,
					expected:   1,
				},
			},
			expectedBytes: []byte{4},
		},
		{
			name: "Group 2 - With Extension",
			tests: []struct {
				value      uint64
				constraint *Constraint
				extensible bool
				expected   uint64
			}{
				{
					value:      1,
					constraint: &Constraint{Lb: 0, Ub: 4},
					extensible: true,
					expected:   1,
				},
				{
					value:      2,
					constraint: &Constraint{Lb: 0, Ub: 4},
					extensible: true,
					expected:   2,
				},
			},
			expectedBytes: []byte{18},
		},
	}

	for _, group := range testGroups {
		t.Run(group.name, func(t *testing.T) {
			// 1. Encode enumerate values
			buf := new(bytes.Buffer)
			writer := NewWriter(buf)
			for _, tt := range group.tests {
				t.Run(fmt.Sprintf("Value=%d", tt.value), func(t *testing.T) {
					if err := writer.WriteEnumerate(tt.value, *tt.constraint, tt.extensible); err != nil {
						t.Errorf("WriteEnumerate() error = %v", err)
					}
				})
			}
			writer.Close()

			// 2. Verify the encoded result
			if !bytes.Equal(buf.Bytes(), group.expectedBytes) {
				t.Errorf("Encoded buffer = %v, want %v", buf.Bytes(), group.expectedBytes)
			}

			// 3. Decode and compare with expected values
			reader := NewReader(bytes.NewReader(buf.Bytes()))
			for _, tt := range group.tests {
				t.Run(fmt.Sprintf("Expected=%d", tt.expected), func(t *testing.T) {
					decoded, err := reader.ReadEnumerate(*tt.constraint, tt.extensible)
					if err != nil {
						t.Errorf("ReadEnumerate() error = %v", err)
					}
					if decoded != tt.expected {
						t.Errorf("Decoded value = %d, want %d", decoded, tt.expected)
					}
				})
			}
		})
	}
}

type TestItem struct {
	id     int64
	region []byte
	item2  TestItem2
}

type TestItem2 struct {
	id     int64
	region []byte
}

func (item *TestItem2) Encode(aw *AperWriter) error {
	if err := aw.WriteInteger(item.id, nil, false); err != nil {
		return err
	}
	if err := aw.WriteBitString(item.region, 3, &Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return err
	}
	return nil
}

func (item *TestItem2) Decode(aw *AperReader) error {
	id, err := aw.ReadInteger(nil, false)
	if err != nil {
		return err
	} else {
		item.id = id
	}

	item.region, _, err = aw.ReadBitString(&Constraint{Lb: 3, Ub: 3}, false)
	if err != nil {
		return err
	}
	return nil
}

func (item *TestItem) Encode(aw *AperWriter) error {
	if err := aw.WriteInteger(item.id, nil, false); err != nil {
		return err
	}
	if err := aw.WriteBitString(item.region, 8, &Constraint{Lb: 0, Ub: 125}, false); err != nil {
		return err
	}
	item.item2.Encode(aw)
	return nil
}

func (item *TestItem) Decode(aw *AperReader) error {
	var err error
	item.id, err = aw.ReadInteger(nil, false)
	if err != nil {
		return err
	}
	item.region, _, err = aw.ReadBitString(&Constraint{Lb: 0, Ub: 125}, false)
	if err != nil {
		return err
	}
	item.item2.Decode(aw)
	return err
}

func Test_Sequence(t *testing.T) {
	fmt.Printf("Test Write/Read sequence\n")
	// 1. encode sequences
	var buf bytes.Buffer
	writer := NewWriter(&buf)
	items := []*TestItem{
		{id: 100, region: []byte{0xfa}, item2: TestItem2{id: 169, region: []byte{0xa0}}},
		{id: 199, region: []byte{0xfb}, item2: TestItem2{id: 170, region: []byte{0xa0}}},
		{id: 100, region: []byte{0xfc}, item2: TestItem2{id: 171, region: []byte{0xa0}}},
		{id: 201, region: []byte{0xfe}, item2: TestItem2{id: 172, region: []byte{0xa0}}},
	}

	if err := WriteSequenceOf[*TestItem](items, writer, nil, false); err != nil {
		t.Errorf("Fail encoding: %+v", err)
	}

	// 2. decode sequences
	reader := NewReader(bytes.NewReader(buf.Bytes()))

	newItems, err := ReadSequenceOfEx[*TestItem](func() *TestItem {
		return new(TestItem)
	}, reader, nil, false)

	if err != nil {
		t.Errorf("Fail decoding: %+v", err)
	}

	// 3. compare
	if len(newItems) != len(items) {
		t.Errorf("Size mismatch: expected %d, got %d", len(items), len(newItems))
	} else {
		for i := range newItems {
			if !reflect.DeepEqual(newItems[i], items[i]) {
				t.Errorf("Item %d does not match: expected %+v, got %+v", i, items[i], newItems[i])
			} else {
				fmt.Printf("Item %d matches: %+v\n", i, newItems[i])
			}
		}
	}
}
