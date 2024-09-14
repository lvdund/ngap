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
		if err:= w.WriteBits(content, nbits); err != nil {
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

func TestWriteBitStringGroups(t *testing.T) {
	testGroups := []struct {
		name     string
		tests    []struct {
			name       string
			content    []byte
			nbits      uint
			constraint *Constrain
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
				constraint *Constrain
				extensible bool
			}{
				{
					content:    []byte{0xa0}, 
					nbits:      3,
					constraint: &Constrain{Lb: 3, Ub: 3},
				},
				{
					content:    []byte{0xfe}, 
					nbits:      8,
					constraint: &Constrain{Lb: 0, Ub: 125},
				},
				{
					content:    []byte{0xec},
					nbits:      6,
					constraint: &Constrain{Lb: 0, Ub: 255},
				},
				{
					content:    []byte{0xd8}, 
					nbits:      5,
					constraint: &Constrain{Lb: 0, Ub: 555},
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
				constraint *Constrain
				extensible bool
			}{
				{
					content:    []byte{0xa0}, 
					nbits:      3,
					constraint: &Constrain{Lb: 3, Ub: 3},
				},
				{
					content:    []byte{0xa0}, 
					nbits:      3,
					constraint: &Constrain{Lb: 3, Ub: 3},
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
				constraint *Constrain
				extensible bool
			}{
				{
					content:    []byte{0xa0}, 
					nbits:      3,
					constraint: &Constrain{Lb: 3, Ub: 3},
				},
				{
					content:    []byte{0xb0}, 
					nbits:      4,
					constraint: &Constrain{Lb: 4, Ub: 4},
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
			if !bytes.Equal(buf.Bytes(), group.expected) {
				t.Errorf("Final buffer = %v, want %v", buf.Bytes(), group.expected)
			}
		})
	}
}


func TestWriteOctetStringGroup(t *testing.T) {
	testGroups := []struct {
		name     string
		tests    []struct {
			name       string
			content    []byte
			constraint *Constrain
			extensible bool
		}
		expected []byte // Expected value after writing all tests in the group
	}{
		{
			name: "Group 1",
			tests: []struct {
				name       string
				content    []byte
				constraint *Constrain
				extensible bool
			}{
				{
					content:    OctetString("a"), 
					constraint: &Constrain{Lb: 1, Ub: 1},
					extensible: true,
				},
				{
					content:    OctetString("bcd"), 
					constraint: &Constrain{Lb: 0, Ub: 20},
					extensible: false,
				},
				{
					content:    OctetString("ef"),
					constraint: &Constrain{Lb: 2, Ub: 2},
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
				constraint *Constrain
				extensible bool
			}{
				{
					content:    []byte("a"), 
					constraint: &Constrain{Lb: 1, Ub: 1},
					extensible: true,
				},
				{
					content:    []byte("abcdefgh"), 
					constraint: &Constrain{Lb: 0, Ub: 20},
					extensible: false,
				},
				{
					content:    []byte("ij"),
					constraint: &Constrain{Lb: 2, Ub: 2},
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
				constraint *Constrain
				extensible bool
			}{
				{
					content:    []byte("a"), 
					constraint: &Constrain{Lb: 1, Ub: 1},
					extensible: true,
				},
				{
					content:    []byte(""), 
					constraint: &Constrain{Lb: 0, Ub: 20},
					extensible: false,
				},
				{
					content:    []byte("bc"),
					constraint: &Constrain{Lb: 2, Ub: 2},
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
					fmt.Println("----------in : ",tt.content)
					fmt.Println("----------out : ",  buf.Bytes() )
				})
			}
			if !bytes.Equal(buf.Bytes(), group.expected) {
				t.Errorf("Final buffer = %v, want %v", buf.Bytes(), group.expected)
			}
		})
	}
}

func TestWriteInteger(t *testing.T) {
	testGroups := []struct {
		name     string
		tests    []struct {
			value      int64
			constraint *Constrain
			extensible bool
		}
		expected []byte // Expected value after writing all tests in the group
	}{
		{
			name: "Group 1",
			tests: []struct {
				value      int64
				constraint *Constrain
				extensible bool
			}{
				{
					value:      45,
					constraint: &Constrain{Lb: 1, Ub: 110},
					extensible: false,
				},
				{
					value:      123,
					constraint: &Constrain{Lb: 0, Ub: 255},
					extensible: false,
				},
				{
					value:      6445,
					constraint: &Constrain{Lb: 0, Ub: 255},
					extensible:  true, 
				},
			},
			expected: []byte{0x58, 0x7B, 0x80, 0x02, 0x19, 0x2D},
		},
	}

	for _, group := range testGroups {
		t.Run(group.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			writer := NewWriter(buf)

			for _, tt := range group.tests {
				_ = writer.WriteInteger(tt.value, tt.constraint, tt.extensible)

			}
			if !bytes.Equal(buf.Bytes(), group.expected) {
				t.Errorf("Final buffer = %v, want %v", buf.Bytes(), group.expected)
			}
		})
	}
}

func TestReadBitStringGroups(t *testing.T) {
	testGroups := []struct {
		name     string
		input    []byte // Input data for decoding
		tests    []struct {
			name       string
			expected   []byte // Expected decoded value (bit string)
			nbits      uint
			constraint *Constrain
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
				constraint *Constrain
				extensible bool
			}{
				{
					expected:   []byte{0xa0}, // expected decoded result
					nbits:      3,
					constraint: &Constrain{Lb: 3, Ub: 3},
				},
				{
					expected:   []byte{0xfe}, // expected decoded result
					nbits:      8,
					constraint: &Constrain{Lb: 0, Ub: 125},
				},
				{
					expected:   []byte{0xec}, // expected decoded result
					nbits:      6,
					constraint: &Constrain{Lb: 0, Ub: 255},
				},
				{
					expected:   []byte{0xd8}, // expected decoded result
					nbits:      5,
					constraint: &Constrain{Lb: 0, Ub: 555},
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
				constraint *Constrain
				extensible bool
			}{
				{
					expected:   []byte{0xa0},
					nbits:      3,
					constraint: &Constrain{Lb: 3, Ub: 3},
				},
				{
					expected:   []byte{0xa0},
					nbits:      3,
					constraint: &Constrain{Lb: 3, Ub: 3},
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
				constraint *Constrain
				extensible bool
			}{
				{
					expected:   []byte{0xa0},
					nbits:      3,
					constraint: &Constrain{Lb: 3, Ub: 3},
				},
				{
					expected:   []byte{0xb0},
					nbits:      4,
					constraint: &Constrain{Lb: 4, Ub: 4},
				},
			},
		},
	}
	for _, group := range testGroups {
		t.Run(group.name, func(t *testing.T) {
			reader := NewReader(bytes.NewReader(group.input))
			for _, tt := range group.tests {
				t.Run(tt.name, func(t *testing.T) {
					decoded, err := reader.ReadBitString(tt.constraint, tt.extensible)
					if err != nil {
						t.Errorf("Error decoding: %v", err)
					}
					if !bytes.Equal(decoded.Bytes, tt.expected) {
						t.Errorf("Decoded bitstring = %v, want %v", decoded, tt.expected)
					}
				})
			}
		})
	}
}

func TestReadInteger(t *testing.T) {
	testGroups := []struct {
		name     string
		input    []byte // Input data for decoding
		tests    []struct {
			expected   int64 // Expected decoded value
			constraint *Constrain
			extensible bool
		}
	}{
		{
			name:  "Group 1",
			input: []byte{0x58, 0x7B, 0x80, 0x02, 0x19, 0x2D},
			tests: []struct {
				expected   int64
				constraint *Constrain
				extensible bool
			}{
				{
					expected:   45,
					constraint: &Constrain{Lb: 1, Ub: 110},
					extensible: false,
				},
				{
					expected:   123,
					constraint: &Constrain{Lb: 0, Ub: 255},
					extensible: false,
				},
				// {
				// 	expected:   6445,
				// 	constraint: &Constrain{Lb: 0, Ub: 255},
				// 	extensible: true,
				// },
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
