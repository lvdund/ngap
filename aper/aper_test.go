package aper

import (
	"bytes"
	"fmt"
	"testing"
)

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

	r := NewReader(bytes.NewBuffer(buf.Bytes()))
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
func TestWriteBitStringGroups(t *testing.T) {
	fmt.Printf("Test WriteBiString\n")
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
			for _, tt := range group.tests {
				t.Run(tt.name, func(t *testing.T) {
					_ = writer.WriteBitString(tt.content, tt.nbits, tt.constraint, tt.extensible)
				})
			}
			writer.flush()
			if !bytes.Equal(buf.Bytes(), group.expected) {
				t.Errorf("Final buffer = %v, want %v", buf.Bytes(), group.expected)
			}
		})
	}
}

func TestReadBitStringGroups(t *testing.T) {
	fmt.Printf("Test ReadBiString\n")
	testGroups := []struct {
		name  string
		input []byte // Input data for decoding
		tests []struct {
			name       string
			expected   []byte // Expected decoded value (bit string)
			nbits      uint
			constraint *Constraint
			extensible bool
		}
	}{
		{
			name:  "Group 1",
			input: []byte{0xA2, 0x00, 0xFE, 0x06, 0xEC, 0x00, 0x05, 0xD8},
			tests: []struct {
				name       string
				expected   []byte
				nbits      uint
				constraint *Constraint
				extensible bool
			}{
				{
					expected:   []byte{0xa0}, // expected decoded result
					nbits:      3,
					constraint: &Constraint{Lb: 3, Ub: 3},
				},
				{
					expected:   []byte{0xfe}, // expected decoded result
					nbits:      8,
					constraint: &Constraint{Lb: 0, Ub: 125},
				},
				{
					expected:   []byte{0xec}, // expected decoded result
					nbits:      6,
					constraint: &Constraint{Lb: 0, Ub: 255},
				},
				{
					expected:   []byte{0xd8}, // expected decoded result
					nbits:      5,
					constraint: &Constraint{Lb: 0, Ub: 555},
				},
			},
		},
		{
			name:  "Group 2",
			input: []byte{180},
			tests: []struct {
				name       string
				expected   []byte
				nbits      uint
				constraint *Constraint
				extensible bool
			}{
				{
					expected:   []byte{0xa0},
					nbits:      3,
					constraint: &Constraint{Lb: 3, Ub: 3},
				},
				{
					expected:   []byte{0xa0},
					nbits:      3,
					constraint: &Constraint{Lb: 3, Ub: 3},
				},
			},
		},
		{
			name:  "Group 3",
			input: []byte{182},
			tests: []struct {
				name       string
				expected   []byte
				nbits      uint
				constraint *Constraint
				extensible bool
			}{
				{
					expected:   []byte{0xa0},
					nbits:      3,
					constraint: &Constraint{Lb: 3, Ub: 3},
				},
				{
					expected:   []byte{0xb0},
					nbits:      4,
					constraint: &Constraint{Lb: 4, Ub: 4},
				},
			},
		},
	}
	for _, group := range testGroups {
		t.Run(group.name, func(t *testing.T) {
			reader := NewReader(bytes.NewReader(group.input))
			for _, tt := range group.tests {
				t.Run(tt.name, func(t *testing.T) {
					decodedBytes, _, err := reader.ReadBitString(tt.constraint, tt.extensible)
					if err != nil {
						t.Errorf("Error decoding: %v", err)
					}
					if !bytes.Equal(decodedBytes, tt.expected) {
						t.Errorf("Decoded bitstring = %v, want %v", decodedBytes, tt.expected)
					}
				})
			}
		})
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

func TestWriteOctetStringGroup(t *testing.T) {
	testGroups := []struct {
		name  string
		tests []struct {
			name       string
			content    []byte
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
				constraint *Constraint
				extensible bool
			}{
				{
					content:    OctetString("a"),
					constraint: &Constraint{Lb: 1, Ub: 1},
					extensible: true,
				},
				{
					content:    OctetString("bcd"),
					constraint: &Constraint{Lb: 0, Ub: 20},
					extensible: false,
				},
				{
					content:    OctetString("ef"),
					constraint: &Constraint{Lb: 2, Ub: 2},
					extensible: false,
				},
			},
			expected: []byte("\x30\x8Cbcdef"),
		},
		{
			name: "Group 2",
			tests: []struct {
				name       string
				content    []byte
				constraint *Constraint
				extensible bool
			}{
				{
					content:    []byte("a"),
					constraint: &Constraint{Lb: 1, Ub: 1},
					extensible: true,
				},
				{
					content:    []byte("abcdefgh"),
					constraint: &Constraint{Lb: 0, Ub: 20},
					extensible: false,
				},
				{
					content:    []byte("ij"),
					constraint: &Constraint{Lb: 2, Ub: 2},
					extensible: false,
				},
			},
			expected: []byte("\x30\xA0abcdefghij"),
		},

		{
			name: "Group 3",
			tests: []struct {
				name       string
				content    []byte
				constraint *Constraint
				extensible bool
			}{
				{
					content:    []byte("a"),
					constraint: &Constraint{Lb: 1, Ub: 1},
					extensible: true,
				},
				{
					content:    []byte(""),
					constraint: &Constraint{Lb: 0, Ub: 20},
					extensible: false,
				},
				{
					content:    []byte("bc"),
					constraint: &Constraint{Lb: 2, Ub: 2},
					extensible: false,
				},
			},
			expected: []byte("\x30\x81\x89\x8C"),
		},
	}
	for _, group := range testGroups {
		t.Run(group.name, func(t *testing.T) {
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
			if !bytes.Equal(buf.Bytes(), group.expected) {
				t.Errorf("Final buffer = %v, want %v", buf.Bytes(), group.expected)
			}
		})
	}
}
func TestReadOctetStringGroup(t *testing.T) {
	testGroups := []struct {
		name  string
		input []byte // Encoded input to decode
		tests []struct {
			name       string
			constraint *Constraint
			extensible bool
			expected   []byte // Expected decoded octet string
		}
	}{
		{
			name:  "Group 1",
			input: []byte("\x30\x8Cbcdef"),
			tests: []struct {
				name       string
				constraint *Constraint
				extensible bool
				expected   []byte
			}{
				{
					constraint: &Constraint{Lb: 1, Ub: 1},
					extensible: true,
					expected:   []byte("a"),
				},
				{
					constraint: &Constraint{Lb: 0, Ub: 20},
					extensible: false,
					expected:   []byte("bcd"),
				},
				{
					constraint: &Constraint{Lb: 2, Ub: 2},
					extensible: false,
					expected:   []byte("ef"),
				},
			},
		},
		{
			name:  "Group 2",
			input: []byte("\x30\xA0abcdefghij"),
			tests: []struct {
				name       string
				constraint *Constraint
				extensible bool
				expected   []byte
			}{
				{
					constraint: &Constraint{Lb: 1, Ub: 1},
					extensible: true,
					expected:   []byte("a"),
				},
				{
					constraint: &Constraint{Lb: 0, Ub: 20},
					extensible: false,
					expected:   []byte("abcdefgh"),
				},
				{
					constraint: &Constraint{Lb: 2, Ub: 2},
					extensible: false,
					expected:   []byte("ij"),
				},
			},
		},
		{
			name:  "Group 3",
			input: []byte("\x30\x81\x89\x8C"),
			tests: []struct {
				name       string
				constraint *Constraint
				extensible bool
				expected   []byte
			}{
				{
					constraint: &Constraint{Lb: 1, Ub: 1},
					extensible: true,
					expected:   []byte("a"),
				},
				{
					constraint: &Constraint{Lb: 0, Ub: 20},
					extensible: false,
					expected:   []byte(""),
				},
				{
					constraint: &Constraint{Lb: 2, Ub: 2},
					extensible: false,
					expected:   []byte("bc"),
				},
			},
		},
	}

	for _, group := range testGroups {
		t.Run(group.name, func(t *testing.T) {
			reader := NewReader(bytes.NewReader(group.input))
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

func TestWriteInteger(t *testing.T) {
	testGroups := []struct {
		name  string
		tests []struct {
			value      int64
			constraint *Constraint
			extensible bool
		}
		expected []byte // Expected value after writing all tests in the group
	}{
		{
			name: "Group 1",
			tests: []struct {
				value      int64
				constraint *Constraint
				extensible bool
			}{
				{
					value:      45,
					constraint: &Constraint{Lb: 1, Ub: 110},
					extensible: false,
				},
				{
					value:      123,
					constraint: &Constraint{Lb: 0, Ub: 255},
					extensible: false,
				},
				{
					value:      6445,
					constraint: &Constraint{Lb: 0, Ub: 255},
					extensible: true,
				},
			},
			expected: []byte{0x58, 0x7B, 0x80, 0x02, 0x19, 0x2D},
		},
	}

	for _, group := range testGroups {
		buf := new(bytes.Buffer)
		writer := NewWriter(buf)
		t.Run(group.name, func(t *testing.T) {

			for _, tt := range group.tests {
				_ = writer.WriteInteger(tt.value, tt.constraint, tt.extensible)

			}
			writer.flush()
			if !bytes.Equal(buf.Bytes(), group.expected) {
				t.Errorf("Final buffer = %v, want %v", buf.Bytes(), group.expected)
			}
		})
	}
}

func TestReadInteger(t *testing.T) {
	testGroups := []struct {
		name  string
		input []byte // Input data for decoding
		tests []struct {
			expected   int64 // Expected decoded value
			constraint *Constraint
			extensible bool
		}
	}{
		{
			name:  "Group 1",
			input: []byte{0x58, 0x7B, 0x80, 0x02, 0x19, 0x2D},
			tests: []struct {
				expected   int64
				constraint *Constraint
				extensible bool
			}{
				{
					expected:   45,
					constraint: &Constraint{Lb: 1, Ub: 110},
					extensible: false,
				},
				{
					expected:   123,
					constraint: &Constraint{Lb: 0, Ub: 255},
					extensible: false,
				},
				{
					expected:   6445,
					constraint: &Constraint{Lb: 0, Ub: 255},
					extensible: true,
				},
			},
		},
	}

	for _, group := range testGroups {
		t.Run(group.name, func(t *testing.T) {
			reader := NewReader(bytes.NewReader(group.input))
			for _, tt := range group.tests {
				t.Run(fmt.Sprintf("Expected=%d", tt.expected), func(t *testing.T) {
					decoded, err := reader.ReadInteger(tt.constraint, tt.extensible)
					if err != nil {
						t.Errorf("Error decoding: %v", err)
					}
					if decoded != tt.expected {
						t.Errorf("Decoded value = %d, want %d", decoded, tt.expected)
					}
				})
			}
		})
	}
}

func TestReadEnumerate(t *testing.T) {
	fmt.Printf("Test ReadEnumerate\n")
	testGroups := []struct {
		name  string
		input []byte // Input data for decoding
		tests []struct {
			expected   uint64 // Expected decoded value
			constraint *Constraint
			extensible bool
		}
	}{
		{
			name: "Group 1 - Definite Range",
			//input: []byte{0x00, 0x40}, // Input data for enumTest1
			//NOTE:the above input is composed from the two inputs of the 2 test
			//cases in free5gc and it is not correct becuase enumerate value
			//may take a few bits (<8)
			input: []byte{4}, // Input data for enumTest1
			tests: []struct {
				expected   uint64
				constraint *Constraint
				extensible bool
			}{
				{
					expected:   0,
					constraint: &Constraint{Lb: 0, Ub: 3},
					extensible: false,
				},
				{
					expected:   1,
					constraint: &Constraint{Lb: 0, Ub: 3},
					extensible: false,
				},
			},
		},
		{
			name: "Group 2 - With Extension",
			//input: []byte{0x10, 0x20}, // Input data for enumTest2
			//NOTE: similarly, the above input is not correct

			input: []byte{18}, // Input data for enumTest2
			tests: []struct {
				expected   uint64
				constraint *Constraint
				extensible bool
			}{
				{
					expected:   1,
					constraint: &Constraint{Lb: 0, Ub: 4},
					extensible: true,
				},
				{
					expected:   2,
					constraint: &Constraint{Lb: 0, Ub: 4},
					extensible: true,
				},
			},
		},
	}

	for _, group := range testGroups {
		t.Run(group.name, func(t *testing.T) {
			reader := NewReader(bytes.NewReader(group.input))
			for _, tt := range group.tests {
				t.Run(fmt.Sprintf("Expected=%d", tt.expected), func(t *testing.T) {
					decoded, err := reader.ReadEnumerate(*tt.constraint, tt.extensible)
					if err != nil {
						t.Errorf("Error decoding: %v", err)
					}
					if decoded != tt.expected {
						t.Errorf("Decoded value = %d, want %d", decoded, tt.expected)
					}
				})
			}
		})
	}
}

func TestWriteEnumerate(t *testing.T) {
	fmt.Printf("Test WriteEnumerate\n")
	testGroups := []struct {
		name  string
		input []byte // Input data for decoding
		tests []struct {
			expected   uint64 // Expected decoded value
			constraint *Constraint
			extensible bool
		}
	}{
		{
			name: "Group 1 - Definite Range",
			//input: []byte{0x00, 0x40}, // Input data for enumTest1
			input: []byte{4}, // Input data for enumTest1
			tests: []struct {
				expected   uint64
				constraint *Constraint
				extensible bool
			}{
				{
					expected:   0,
					constraint: &Constraint{Lb: 0, Ub: 3},
					extensible: false,
				},
				{
					expected:   1,
					constraint: &Constraint{Lb: 0, Ub: 3},
					extensible: false,
				},
			},
		},
		{
			name: "Group 2 - With Extension",
			//input: []byte{0x10, 0x20}, // Input data for enumTest2
			input: []byte{18}, // Input data for enumTest2
			tests: []struct {
				expected   uint64
				constraint *Constraint
				extensible bool
			}{
				{
					expected:   1,
					constraint: &Constraint{Lb: 0, Ub: 4},
					extensible: true,
				},
				{
					expected:   2,
					constraint: &Constraint{Lb: 0, Ub: 4},
					extensible: true,
				},
			},
		},
	}

	for _, group := range testGroups {
		t.Run(group.name, func(t *testing.T) {
			var buf bytes.Buffer
			writer := NewWriter(&buf)
			for _, tt := range group.tests {
				t.Run(fmt.Sprintf("Expected=%d", tt.expected), func(t *testing.T) {
					if err := writer.WriteEnumerate(tt.expected, *tt.constraint, tt.extensible); err != nil {
						t.Errorf("Error encoding: %v", err)
					}
				})
			}
			writer.Close()
			if !bytes.Equal(buf.Bytes(), group.input) {
				t.Errorf("ouput = %v, want %v", buf.Bytes(), group.input)
			}

		})
	}
}

type TestItem struct {
	id  int64
	msg IE
}

func (item TestItem) Encode(aw AperWriter) (err error) {
	err = aw.WriteInteger(item.id, nil, false)
	return
}

func (item TestItem) Decode(aw AperReader) (err error) {
	item.id, err = aw.ReadInteger(nil, false)
	return
}

func Test_Sequence(t *testing.T) {
	//@Duc: please improve this test

	fmt.Printf("Test Write/Read sequence\n")
	//1. encode sequences
	var buf bytes.Buffer
	writer := NewWriter(&buf)
	items := []TestItem{
		TestItem{
			id: 100,
			msg: &AmfId{},
		},
		TestItem{
			id: 199,
			msg: AmfName("aa"),
		},
	}
	if err := WriteSequenceOf[TestItem](items, writer, nil, false); err != nil {
		t.Errorf("Fail encoding: %+v", err)
	}

	//2. decode sequences
	reader := NewReader(bytes.NewReader(buf.Bytes()))
	if newItems, err := ReadSequenceOfEx[TestItem](reader, nil, false); err != nil {
		t.Errorf("Fail decoding: %+v", err)
	} else {
		//3. compare
		if len(newItems) != len(items) {
			t.Errorf("size not match")
		} else {
			fmt.Printf("num items = %d\n", len(newItems))
			//TODO: compare content
		}
	}

}
