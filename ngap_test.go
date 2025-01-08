package ngap

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/lvdund/ngap/aper"
	"github.com/lvdund/ngap/ies"
)

func Test_NgNew(t *testing.T) {
	msg := ies.UERadioCapabilityInfoIndication{
		AMFUENGAPID:       int64(1),
		RANUENGAPID:       int64(1),
		UERadioCapability: []byte{1, 1, 1},
	}
	var b []byte
	var err error
	if b, err = NgapEncode(&msg); err != nil {
		fmt.Println("NgapEncode err:", err)
		return
	}
	fmt.Printf("Encoded: %v", b)

	fmt.Println("================================")
	fmt.Println("================================")
	fmt.Println("Decode:")

	var pdu NgapPdu
	if pdu, err, _ = NgapDecode(b); err != nil {
		fmt.Println("NgapDecode err:", err)
		return
	}
	fmt.Println("Decoded:", pdu)
	decode_msg := pdu.Message.Msg.(*ies.UERadioCapabilityInfoIndication)
	fmt.Println(decode_msg.AMFUENGAPID)
	fmt.Println(decode_msg.RANUENGAPID)
	fmt.Println(decode_msg.UERadioCapability)

	fmt.Println("================================")
	fmt.Println("================================")

	var buf bytes.Buffer
	w := aper.NewWriter(&buf)
	ie := ies.NewBITSTRING([]byte{0xa0}, aper.Constraint{Lb: 1, Ub: 3}, false)
	if err = ie.Encode(w); err != nil {
		fmt.Println("err ie encode:", err)
		return
	}
	w.Close()
	fmt.Println("encode ie:", buf.Bytes())

	r := aper.NewReader(&buf)
	ie = ies.NewBITSTRING(nil, aper.Constraint{Lb: 1, Ub: 3}, false)
	if err = ie.Decode(r); err != nil {
		fmt.Println("err ie encode:", err)
		return
	}
	fmt.Println("Decode ie:", ie.Value)
}

func Test_Sequence(t *testing.T) {
	ie := ies.AdditionalQosFlowInformation{Value: 5}
	elements := []*ies.AdditionalQosFlowInformation{&ie, &ie, &ie}
	s := ies.NewSequence[*ies.AdditionalQosFlowInformation](
		elements,
		aper.Constraint{Lb: 0, Ub: 5},
		false,
	)

	var buf bytes.Buffer
	w := aper.NewWriter(&buf)
	if err := s.Encode(w); err != nil {
		fmt.Println("encode sequence err:", err)
		return
	}
	w.Close()
	fmt.Println("Encode:", buf.Bytes())

	fmt.Println("================================")
	fmt.Println("================================")

	r := aper.NewReader(&buf)
	ss := ies.NewSequence[*ies.AdditionalQosFlowInformation](
		nil,
		aper.Constraint{Lb: 1, Ub: 5},
		false,
	)
	fn := func() *ies.AdditionalQosFlowInformation {
		return &ies.AdditionalQosFlowInformation{}
	}
	if err := ss.Decode(r, fn); err != nil {
		fmt.Println("encode sequence err:", err)
		return
	}
	sss := []*ies.AdditionalQosFlowInformation{}
	sss = append(sss, ss.Value...)
	fmt.Println("Decode", len(sss), sss[0].Value)
}

func Test_Bitstring(t *testing.T) {
	var buf bytes.Buffer
	var err error

	w := aper.NewWriter(&buf)
	b := []byte{0xa0}
	fmt.Println(len(b))
	if err = w.WriteBitString(
		b,
		1,
		&aper.Constraint{Lb: 1, Ub: 3},
		false); err != nil {
		fmt.Println("err ie encode:", err)
		return
	}
	w.Close()
	fmt.Println("encode ie:", buf.Bytes())

	r := aper.NewReader(&buf)
	var content []byte
	var nbits uint
	content, nbits, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 3}, false)
	if err != nil {
		fmt.Println("err ie encode:", err)
		return
	}
	fmt.Println("Decode ie:", content, nbits)
}

func Test_BITSTRING(t *testing.T) {
	ie := ies.NewBITSTRING([]byte{0x10, 0x10, 0x10, 0x10, 0x10}, aper.Constraint{Lb: 36, Ub: 36}, false)

	var buf bytes.Buffer
	var err error

	w := aper.NewWriter(&buf)
	if err = ie.Encode(w); err != nil {
		fmt.Println("err ie encode:", err)
		return
	}
	w.Close()

	fmt.Println("================================")
	fmt.Println("================================")

	r := aper.NewReader(&buf)
	newie := ies.NewBITSTRING(nil, aper.Constraint{Lb: 36, Ub: 36}, false)
	if err = newie.Decode(r); err != nil {
		fmt.Println("err ie encode:", err)
		return
	}
	fmt.Println(newie)
}
