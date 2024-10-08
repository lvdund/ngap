package ngap

import (
	"bytes"
	"fmt"
	"io"
	"ngap/aper"
	"ngap/ie"
)

type NgapMessageEncoder interface {
	Encode(io.Writer) error
}

// represent an IE in Ngap messages
type NgapMessageIE struct {
	Id          ie.NgapProtocolIeId //protocol IE identity
	Criticality ie.Criticality
	Value       aper.AperMarshaller //open type
}

func (ie NgapMessageIE) Encode(w *aper.AperWriter) (err error) {
	//1. encode protocol Ie Id
	if err = ie.Id.Encode(w); err != nil {
		return
	}
	//2. encode criticality
	if err = ie.Criticality.Encode(w); err != nil {
		return
	}
	//3. encode NgapIE
	//encode IE into a byte array first
	var buf bytes.Buffer
	ieW := aper.NewWriter(&buf)
	if err = ie.Value.Encode(ieW); err != nil {
		return
	}
	//then write the array as open type
	err = w.WriteOpenType(buf.Bytes())
	return
}

//encode an Ngap Message
func encodeMessage(w io.Writer, present uint8, procedureCode int64, criticality aper.Enumerated, ies []NgapMessageIE) (err error) {
	aw := aper.NewWriter(w)
	//0. write extension bit
	if err = aw.WriteBool(aper.Zero); err != nil {
		return
	}
	//1. write present
	if err = aw.WriteChoice(uint64(present), 2, true); err != nil {
		return
	}

	//2. write procedure code
	pCode := ie.ProcedureCode{
		Value: aper.Integer(procedureCode),
	}

	if err = pCode.Encode(aw); err != nil {
		return
	}

	//3. write criticality
	cr := ie.Criticality{
		Value: criticality,
	}
	if err = cr.Encode(aw); err != nil {
		return
	}

	//4. write IE list container
	if len(ies) == 0 {
		err = fmt.Errorf("empty message")
		return
	}

	var buf bytes.Buffer
	cW := aper.NewWriter(&buf) //container writer

	//4.1 container has an extension bit
	cW.WriteBool(aper.Zero)
	//4.2 encode all IEs
	if err = aper.WriteSequenceOf[NgapMessageIE](ies, cW, &aper.Constraint{
		Lb: 0,
		Ub: int64(aper.POW_16 - 1),
	}, false); err != nil {
		return
	}

	if err = cW.Close(); err != nil { //finalize
		return
	}
	// 4.3 write the container
	if err = aw.WriteOpenType(buf.Bytes()); err != nil {
		return
	}
	//5. flush buffer
	err = aw.Close()
	return
}

//API to encode an Ngap Message
func NgapEncode(msg NgapMessageEncoder) (wire []byte, err error) {
	var buf bytes.Buffer
	if err = msg.Encode(&buf); err == nil {
		wire = buf.Bytes()
	}
	return
}
