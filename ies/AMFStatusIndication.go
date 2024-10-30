package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type AMFStatusIndication struct {
	UnavailableGUAMIList *UnavailableGUAMIList `,reject,mandatory`
}

func (msg *AMFStatusIndication) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_AMFStatusIndication, Criticality_PresentIgnore, msg.toIes())
}
func (msg *AMFStatusIndication) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.UnavailableGUAMIList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UnavailableGUAMIList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.UnavailableGUAMIList})
	}
	return
}
func (msg *AMFStatusIndication) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *AMFStatusIndication) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
	id, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: int64(aper.POW_16) - 1}, false)
	if err != nil {
		return
	}
	msgIe = new(NgapMessageIE)
	msgIe.Id.Value = aper.Integer(id)
	c, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	if err != nil {
		return
	}
	msgIe.Criticality.Value = aper.Enumerated(c)
	var buf []byte
	if buf, err = r.ReadOpenType(); err != nil {
		return
	}
	ieR := aper.NewReader(bytes.NewReader(buf))
	switch msgIe.Id.Value {
	case ProtocolIEID_UnavailableGUAMIList:
		var tmp UnavailableGUAMIList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UnavailableGUAMIList = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
