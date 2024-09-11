package aper

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type testData struct {
	in  []byte
	Out interface{}
}

type BitStringStructTest1 struct {
	BitString1 BitString `aper:"sizeLB:3,sizeUB:3"`
	BitString2 BitString `aper:"sizeLB:3,sizeUB:3"`
}

var BitStringStructTest1Data = []BitStringStructTest1{
	{BitString{[]byte{0xa0}, 3}, BitString{[]byte{0xa0}, 3}},
}

type BitStringStructTest2 struct {
	BitString1 BitString `aper:"sizeLB:3,sizeUB:3"`
	BitString2 BitString `aper:"sizeLB:4,sizeUB:4"`
}

var BitStringStructTest2Data = []BitStringStructTest2{
	{BitString{[]byte{0xa0}, 3}, BitString{[]byte{0xb0}, 4}},
}

type BitStringStructTest3 struct {
	BitString1 BitString `aper:"sizeLB:3,sizeUB:3"`
	BitString2 BitString `aper:"sizeLB:0,sizeUB:125"`
	BitString3 BitString `aper:"sizeLB:0,sizeUB:255"`
	BitString4 BitString `aper:"sizeLB:0,sizeUB:555"`
}

var BitStringStructTest3Data = []BitStringStructTest3{
	{
		BitString{[]byte{0xa0}, 3},
		BitString{[]byte{0xfe}, 8},
		BitString{[]byte{0xec}, 6},
		BitString{[]byte{0xd8}, 5},
	},
}

var structBitStringTestData = []testData{
	{[]byte{0xB4}, BitStringStructTest1Data[0]},
	{[]byte{0xB6}, BitStringStructTest2Data[0]},
	{[]byte{0xA2, 0x00, 0xFE, 0x06, 0xEC, 0x00, 0x05, 0xD8}, BitStringStructTest3Data[0]},
}

func TestDecodeStructBitString(t *testing.T) {
	for i, test := range structBitStringTestData {
		fmt.Print(1, fmt.Sprintf("[TEST %d]\n", i+1))
		out := reflect.New(reflect.TypeOf(test.Out))
		err := Decode(test.in, out.Interface())
		fmt.Print(2, fmt.Sprintf("	in : %0x", test.in))

		for j := 0; j < reflect.TypeOf(out.Elem().Interface()).NumField(); j++ {
		}
		if err != nil {
		} else if reflect.DeepEqual(test.Out, out.Elem().Interface()) {
			fmt.Print(1, "[PASS]\n")
			continue
		}
		fmt.Print(1, "[FAIL]\n")
		t.Errorf("TEST %d is FAILED", i+1)
	}
}

func TestEncodeBitString(t *testing.T) {
	for i, test := range structBitStringTestData {
		fmt.Print(1, fmt.Sprintf("[TEST %d]\n", i+1))

		out, err := Encode(test.Out)
		if err != nil {
			t.Errorf("TEST %d Marshal failed: %v", i+1, err)
			continue
		}
		if !reflect.DeepEqual(test.in, out) {
			fmt.Print(1, "[FAIL]\n")
			t.Errorf("TEST %d is FAILED\nExpected: %0x\nGot: %0x", i+1, test.in, out)
		} else {
			fmt.Print(1, "[PASS]\n")
		}
	}
}

//------integer
type intTest1 struct {
	Value int64
}

var intTest1Data = []intTest1{
	{3},
	{333333},
	{-333333},
}

// value is definite
type intTest2 struct {
	Value int64 `aper:"valueLB:3,valueUB:3"`
}

var intTest2Data = []intTest2{
	{3},
}

// 2 <= bmax-bmin <= 255
type intTest3 struct {
	Value int64 `aper:"valueLB:1,valueUB:110"`
}

var intTest3Data = []intTest3{
	{11},
	{13},
	{28},
	{110},
}

// bmax-bmin == 256
type intTest4 struct {
	Value int64 `aper:"valueLB:0,valueUB:255"`
}

var intTest4Data = []intTest4{
	{140},
}

// 257 <= bmax-bmin <= 65536
type intTest5 struct {
	Value int64 `aper:"valueLB:0,valueUB:65535"`
}

var intTest5Data = []intTest5{
	{140},
}

// 65537 <= bmax-bmin
type intTest6 struct {
	Value int64 `aper:"valueLB:0,valueUB:4294967295"`
}

var intTest6Data = []intTest6{
	{140},
	{4294967295},
	{65535},
	{65536},
}

// value extensed
type intTest7 struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:45"`
}

var intTest7Data = []intTest7{
	{140},
	{2147483647},
	{65535},
	{65536},
	{33},
}

var integerTestData = []testData{
	{[]byte{0x01, 0x03}, intTest1Data[0]},
	{[]byte{0x03, 0x05, 0x16, 0x15}, intTest1Data[1]},
	{[]byte{0x03, 0xFA, 0xE9, 0xEB}, intTest1Data[2]},
	// I don't know if I should treat it as an error in this case.
	{[]byte{0x00}, intTest1{0}},
	{[]byte{0x00}, intTest2Data[0]},
	{[]byte{0x14}, intTest3Data[0]},
	{[]byte{0x18}, intTest3Data[1]},
	{[]byte{0x36}, intTest3Data[2]},
	{[]byte{0xDA}, intTest3Data[3]},
	{[]byte{0x8C}, intTest4Data[0]},
	{[]byte{0x00, 0x8C}, intTest5Data[0]},
	{[]byte{0x00, 0x8C}, intTest6Data[0]},
	{[]byte{0xC0, 0xFF, 0xFF, 0xFF, 0xFF}, intTest6Data[1]},
	{[]byte{0x40, 0xFF, 0xFF}, intTest6Data[2]},
	{[]byte{0x80, 0x01, 0x00, 0x00}, intTest6Data[3]},
	{[]byte{0x80, 0x02, 0x00, 0x8C}, intTest7Data[0]},
	{[]byte{0x80, 0x04, 0x7F, 0xFF, 0xFF, 0xFF}, intTest7Data[1]},
	{[]byte{0x80, 0x03, 0x00, 0xFF, 0xFF}, intTest7Data[2]},
	{[]byte{0x80, 0x03, 0x01, 0x00, 0x00}, intTest7Data[3]},
	{[]byte{0x42}, intTest7Data[4]},
}

type intStructTest1 struct {
	Int1 int64 `aper:"valueLB:1,valueUB:110"`
	Int2 int64 `aper:"valueLB:0,valueUB:255"`
	Int3 int64 `aper:"valueExt,valueLB:0,valueUB:45"`
}

var intStructTest1Data = []intStructTest1{
	{45, 123, 6445},
}

var integerStructTestData = []testData{
	{[]byte{0x58, 0x7B, 0x80, 0x02, 0x19, 0x2D}, intStructTest1Data[0]},
}

func DeccodeStructInteger(t *testing.T) {
	for i, test := range integerTestData {
		out := reflect.New(reflect.TypeOf(test.Out))
		err := Decode(test.in, out.Interface())
		for j := 0; j < reflect.TypeOf(out.Elem().Interface()).NumField(); j++ {
		}
		if err != nil {
		} else if reflect.DeepEqual(test.Out, out.Elem().Interface()) {
			continue
		}
		t.Errorf("TEST %d is FAILED", i+1)
	}
}

func TestDeccodeStructInteger(t *testing.T) {
	for i, test := range integerStructTestData {
		out := reflect.New(reflect.TypeOf(test.Out))
		err := Decode(test.in, out.Interface())
		for j := 0; j < reflect.TypeOf(out.Elem().Interface()).NumField(); j++ {
		}
		if err != nil {
		} else if reflect.DeepEqual(test.Out, out.Elem().Interface()) {
			continue
		}
		t.Errorf("TEST %d is FAILED", i+1)
	}
}

func TestEncodeInteger(t *testing.T) {
	for i, test := range integerStructTestData {
		out, err := Encode(test.Out)
		if err != nil {
			t.Errorf("TEST %d Marshal failed: %v", i+1, err)
			continue
		}
		if !reflect.DeepEqual(test.in, out) {
			t.Errorf("TEST %d is FAILED\nExpected: %0x\nGot: %0x", i+1, test.in, out)
		} else {
		}
	}
}

///test type boolean

type boolTest1 struct {
	Value bool
}

var boolTest1Data = []boolTest1{
	{false},
	{true},
}

var boolTestData = []testData{
	{[]byte{0x00}, boolTest1Data[0]},
	{[]byte{0x80}, boolTest1Data[1]},
}

func TestMarshalBoolean(t *testing.T) {
	for i, test := range boolTestData {

		out, err := Encode(test.Out)
		if err != nil {
			t.Errorf("TEST %d Marshal failed: %v", i+1, err)
			continue
		}

		if !reflect.DeepEqual(test.in, out) {
			t.Errorf("TEST %d is FAILED\nExpected: %0x\nGot: %0x", i+1, test.in, out)
		} else {
		}
	}
}

// ---------------------------------------
// oCTETStringTest1 is for no constraint
type oCTETStringTest1 struct {
	OctetString OctetString
}

var oCTETStringTest1Data = []oCTETStringTest1{
	{OctetString("free5GC")},
	{OctetString("\x23\x34\x52\x97")},
	{OctetString("Jennifer")},
}

// oCTETStringTest2 is for (lmin＝lmax) < 3
type oCTETStringTest2 struct {
	OctetString OctetString `aper:"sizeLB:2,sizeUB:2"`
}

var oCTETStringTest2Data = []oCTETStringTest2{
	{OctetString("\xaa\x56")},
	{OctetString("\x43\x12")},
}

// oCTETStringTest3 is for 3 <= (lmin＝lmax) <= 65536
type oCTETStringTest3 struct {
	OctetString OctetString `aper:"sizeLB:20,sizeUB:20"`
}

var oCTETStringTest3Data = []oCTETStringTest3{
	{OctetString("LLpRB9oV8zOkfraw1Nf5")},
}

// the following is for lmax != min

// oCTETStringTest4 is for 2 <= (lmax-lmin+1) <= 255
type oCTETStringTest4 struct {
	OctetString OctetString `aper:"sizeLB:1,sizeUB:160"`
}

var oCTETStringTest4Data = []oCTETStringTest4{
	{OctetString("LLpRB9oV8zOkfraw1Nf5")},
	{OctetString("1yYPj2WH4Uzex3sU40P1Kq7SgDB2sz0Ksg7fA76zcI5pxVDWtkUrfPti95h7xkzWpAcLaU7fMBBIJ981")},
}

// oCTETStringTest5 is for (lmax-lmin+1) == 256
type oCTETStringTest5 struct {
	OctetString OctetString `aper:"sizeLB:0,sizeUB:255"`
}

var oCTETStringTest5Data = []oCTETStringTest5{
	{OctetString("LLpRB9oV8zOkfraw1Nf5")},
	{OctetString("cGUpp6MH*7@55mntftf$k@eVdd3k2-*dVbGt?BmdTvTvs#ee9cktn6uA5u2g@cvE955P4rUqReG$Ybd83YY?" +
		"r5DqTYqrwDtHzeX+tFVK5RkBmns3GFhU9rPtX-eRfh62+Mmdeav2UFRy$wNghwSm?8RpeqBZTe8W-3Yfm#n=NR..r" +
		"@z6BRXGAX.DMz34ad@-N8Xy-V9AkC-6kPU*Yh$MW7+m-$B6e32!WCCeFe?d-QyV+@z#vKy6meZN87bV2hd")},
}

// oCTETStringTest6 is for 257 <= (lmax-lmin+1) <= 65536
type oCTETStringTest6 struct {
	OctetString OctetString `aper:"sizeLB:0,sizeUB:355"`
}

var oCTETStringTest6Data = []oCTETStringTest6{
	{OctetString("I!nGUXiqNpCP&a")},
	{OctetString("u^YlZwgYxf7swQqweqw")},
	{OctetString("iClFlb&YgrS4basdas")},
	{OctetString("wirelab")},
}

// oCTETStringTest7 is for extensed data
type oCTETStringTest7 struct {
	OctetString OctetString `aper:"sizeExt,sizeLB:1,sizeUB:1"`
}

var oCTETStringTest7Data = []oCTETStringTest7{
	{OctetString("a")},
	{OctetString("free5GC")},
}

// oCTETStringTest8 is for 65537 <= (lmax-lmin+1)
type oCTETStringTest8 struct {
	OctetString OctetString `aper:"sizeLB:0,sizeUB:343434"`
}

var bigOctetData = "\xC4" + strings.Repeat("ab", 32768) + "\xB8\x82" + strings.Repeat("ab", 7232) + "cd"

var oCTETStringTest8Data = []oCTETStringTest8{
	{OctetString("I!nGUXiqNpCP&a")},
	{OctetString(strings.Repeat("ab", 40000) + "cd")},
}

var singleOctetStringTestData = []testData{
	{[]byte("\x07free5GC"), oCTETStringTest1Data[0]},
	{[]byte("\x04\x23\x34\x52\x97"), oCTETStringTest1Data[1]},
	{[]byte("\x08Jennifer"), oCTETStringTest1Data[2]},
	{[]byte("\xaa\x56"), oCTETStringTest2Data[0]},
	{[]byte("\x43\x12"), oCTETStringTest2Data[1]},
	{[]byte("LLpRB9oV8zOkfraw1Nf5"), oCTETStringTest3Data[0]},
	{[]byte("\x13LLpRB9oV8zOkfraw1Nf5"), oCTETStringTest4Data[0]},
	{
		[]byte("O1yYPj2WH4Uzex3sU40P1Kq7SgDB2sz0Ksg7fA76zcI5pxVDWtkUrfPti95h7xkzWpAcLaU7fMBBIJ981"),
		oCTETStringTest4Data[1],
	},
	{[]byte("\x14LLpRB9oV8zOkfraw1Nf5"), oCTETStringTest5Data[0]},
	{
		[]byte("\xFFcGUpp6MH*7@55mntftf$k@eVdd3k2-*dVbGt?BmdTvTvs#ee9cktn6uA5u2g@cvE955P4rUqReG$Ybd83YY" +
			"?r5DqTYqrwDtHzeX+tFVK5RkBmns3GFhU9rPtX-eRfh62+Mmdeav2UFRy$wNghwSm?8RpeqBZTe8W-3Yfm#n=N" +
			"R..r@z6BRXGAX.DMz34ad@-N8Xy-V9AkC-6kPU*Yh$MW7+m-$B6e32!WCCeFe?d-QyV+@z#vKy6meZN87bV2hd"),
		oCTETStringTest5Data[1],
	},
	{[]byte("\x00\x0EI!nGUXiqNpCP&a"), oCTETStringTest6Data[0]},
	{[]byte("\x00\x13u^YlZwgYxf7swQqweqw"), oCTETStringTest6Data[1]},
	{[]byte("\x00\x12iClFlb&YgrS4basdas"), oCTETStringTest6Data[2]},
	{[]byte("\x00\x07wirelab"), oCTETStringTest6Data[3]},
	{[]byte("\x30\x80"), oCTETStringTest7Data[0]},
	{[]byte("\x80\x07free5GC"), oCTETStringTest7Data[1]},
	{[]byte("\x0EI!nGUXiqNpCP&a"), oCTETStringTest8Data[0]},
	{[]byte(bigOctetData), oCTETStringTest8Data[1]},
}

func TestDeocodeSingleOctetString(t *testing.T) {
	for i, test := range singleOctetStringTestData {
		out := reflect.New(reflect.TypeOf(test.Out))
		err := Decode(test.in, out.Interface())
		if err != nil {
		} else if reflect.DeepEqual(test.Out, out.Elem().Interface()) {
			continue
		}
		t.Errorf("TEST %d is FAILED", i+1)
	}
}

func TestMarshalOctetString(t *testing.T) {
	for i, test := range singleOctetStringTestData {
		out, err := Encode(test.Out)
		if err != nil {
			t.Errorf("TEST %d Marshal failed: %v", i+1, err)
			continue
		}
		if !reflect.DeepEqual(test.in, out) {
			t.Errorf("TEST %d is FAILED\nExpected: %0x\nGot: %0x", i+1, test.in, out)
		} else {
		}
	}
}

type oCTETStringStructTest1 struct {
	OctetString1 OctetString `aper:"sizeExt,sizeLB:1,sizeUB:1"`
	OctetString2 OctetString `aper:"sizeLB:0,sizeUB:20"`
	OctetString3 OctetString `aper:"sizeLB:2,sizeUB:2"`
}

var oCTETStringStructTest1Data = []oCTETStringStructTest1{
	{OctetString("a"), OctetString("bcd"), OctetString("ef")},
	{OctetString("a"), OctetString("abcdefgh"), OctetString("ij")},
	{OctetString("a"), OctetString(""), OctetString("bc")},
}

type oCTETStringStructTest2 struct {
	OctetString1 OctetString `aper:"sizeExt,sizeLB:1,sizeUB:1"`
	OctetString3 OctetString `aper:"sizeLB:2,sizeUB:2"`
	OctetString2 OctetString `aper:"sizeLB:0,sizeUB:20"`
}

var oCTETStringStructTest2Data = []oCTETStringStructTest2{
	{OctetString("a"), OctetString("bc"), OctetString("de")},
	{OctetString("a"), OctetString("34"), OctetString("5678")},
	{OctetString("a"), OctetString("12"), OctetString("")},
}

var structOctetStringTestData = []testData{
	{[]byte("\x30\x8Cbcdef"), oCTETStringStructTest1Data[0]},
	{[]byte("\x30\xA0abcdefghij"), oCTETStringStructTest1Data[1]},
	{[]byte("\x30\x81\x89\x8C"), oCTETStringStructTest1Data[2]},
	{[]byte("\x30\xB1\x31\x88de"), oCTETStringStructTest2Data[0]},
	{[]byte("\x30\x99\x9A\x105678"), oCTETStringStructTest2Data[1]},
	{[]byte("\x30\x98\x99\x00"), oCTETStringStructTest2Data[2]},
}

func TestMarshalStructOctetString(t *testing.T) {
	for i, test := range structOctetStringTestData {
		out, err := Encode(test.Out)
		if err != nil {
		} else if reflect.DeepEqual(test.in, out) {
			continue
		}
		t.Errorf("TEST %d is FAILED", i+1)
	}
}


func TestDecodetructOctetString(t *testing.T) {
	for i, test := range structOctetStringTestData {
		out := reflect.New(reflect.TypeOf(test.Out))
		err := Decode(test.in, out.Interface())

		for j := 0; j < reflect.TypeOf(out.Elem().Interface()).NumField(); j++ {
		}
		if err != nil {
		} else if reflect.DeepEqual(test.Out, out.Elem().Interface()) {
			continue
		}
		t.Errorf("TEST %d is FAILED", i+1)
	}
}

//test

// TEST POINTER

type ptrTest1 struct {
	Ptr *intTest1
}

var ptrTest1Data = []ptrTest1{
	{&intTest1Data[0]},
	{&intTest1Data[1]},
	{&intTest1Data[2]},
}

type ptrTest2 struct {
	Ptr *intTest2
}

var ptrTest2Data = []ptrTest2{
	{&intTest2Data[0]},
}

type ptrTest3 struct {
	Ptr *intTest3
}

var ptrTest3Data = []ptrTest3{
	{&intTest3Data[0]},
}

type ptrTest4 struct {
	Ptr *intTest4
}

var ptrTest4Data = []ptrTest4{
	{&intTest4Data[0]},
}

var ptrTestData = []testData{
	{[]byte{0x01, 0x03}, ptrTest1Data[0]},
	{[]byte{0x03, 0x05, 0x16, 0x15}, ptrTest1Data[1]},
	{[]byte{0x03, 0xFA, 0xE9, 0xEB}, ptrTest1Data[2]},
	{[]byte{0x00}, ptrTest2Data[0]},
	{[]byte{0x14}, ptrTest3Data[0]},
	{[]byte{0x8C}, ptrTest4Data[0]},
}

//--------------------------------------

type seqofTest1 struct {
	List []intTest1 `aper:"sizeLB:0,sizeUB:3"`
}

var seqofTest1Data = []seqofTest1{
	{intTest1Data},
}

type seqofTest2 struct {
	List []intStructTest1 `aper:"sizeLB:0,sizeUB:30"`
}

var seqofTest2Data = []seqofTest2{
	{intStructTest1Data},
}

type seqofTest3 struct {
	List []BitStringStructTest3 `aper:"sizeLB:0,sizeUB:50"`
}

var seqofTest3Data = []seqofTest3{
	{BitStringStructTest3Data},
}

type seqofTest4 struct {
	List []intTest7 `aper:"sizeLB:0,sizeUB:16"`
}

var seqofTest4Data = []seqofTest4{
	{intTest7Data},
}

type seqofTest5 struct {
	List []intTest3 `aper:"sizeLB:0,sizeUB:255"`
}

var seqofTest5Data = []seqofTest5{
	{intTest3Data},
}

type seqofTest6 struct {
	List []intTest3
}

var seqofTest6Data = []seqofTest6{
	{intTest3Data},
}

type seqofTest7 struct {
	List []intTest3 `aper:"sizeLB:4,sizeUB:4"`
}

var seqofTest7Data = []seqofTest7{
	{intTest3Data},
}

var seqofTestData = []testData{
	{[]byte{0x0A, 0xC0, 0x7B, 0x80, 0x02, 0x19, 0x2D}, seqofTest2Data[0]},
	{[]byte{0x06, 0x88, 0xFE, 0x06, 0xEC, 0x00, 0x05, 0xD8}, seqofTest3Data[0]},
	{[]byte("\x2C\x02\x00\x8C\x80\x04\x7F\xFF\xFF\xFF\x80\x03\x00\xFF\xFF\x80\x03\x01\x00\x00\x42"), seqofTest4Data[0]},
	{[]byte{0x04, 0x14, 0x30, 0xDE, 0xD0}, seqofTest5Data[0]},
	{[]byte{0x04, 0x14, 0x30, 0xDE, 0xD0}, seqofTest6Data[0]},
	{[]byte{0x14, 0x30, 0xDE, 0xD0}, seqofTest7Data[0]},
}

func TestEncodeSingleSequenceOf(t *testing.T) {
	for i, test := range seqofTestData {
		encoded, err := Encode(test.Out)
		fmt.Println(test.in)
		fmt.Println(encoded)
		if err != nil {
		} else if reflect.DeepEqual(test.in, encoded) {

			continue
		}
		t.Errorf("TEST %d is FAILED", i+1)
	}
}

// -------------enu

// value is unconstraint
type enumTest1 struct {
	Value Enumerated `aper:"valueLB:0,valueUB:3"`
}

var enumTest1Data = []enumTest1{
	{0},
	{1},
}

// value is definite
type enumTest2 struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:4"`
}

var enumTest2Data = []enumTest2{
	{1},
	{2},
}

var enumTestData = []testData{
	{[]byte{0x00}, enumTest1Data[0]},
	{[]byte{0x40}, enumTest1Data[1]},
	{[]byte{0x10}, enumTest2Data[0]},
	{[]byte{0x20}, enumTest2Data[1]},
}

func TestEncodeEnum(t *testing.T) {
	for i, test := range enumTestData {
		encoded, err := Encode(test.Out)
		fmt.Println(test.in)
		fmt.Println(encoded)
		if err != nil {
		} else if reflect.DeepEqual(encoded, test.in) {
			continue
		}
		t.Errorf("TEST ENCODE %d is FAILED", i+1)
	}
}
