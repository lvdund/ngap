package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type PWSCancelRequest struct {
	MessageIdentifier        *MessageIdentifier        `,reject,mandatory`
	SerialNumber             *SerialNumber             `,reject,mandatory`
	WarningAreaList          *WarningAreaList          `,ignore,optional`
	CancelAllWarningMessages *CancelAllWarningMessages `,reject,optional`
}

func (msg *PWSCancelRequest) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_PWSCancel, Criticality_PresentReject, msg.toIes())
}
func (msg *PWSCancelRequest) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.MessageIdentifier != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_MessageIdentifier},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.MessageIdentifier})
	}
	if msg.SerialNumber != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SerialNumber},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.SerialNumber})
	}
	if msg.WarningAreaList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_WarningAreaList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.WarningAreaList})
	}
	if msg.CancelAllWarningMessages != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CancelAllWarningMessages},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.CancelAllWarningMessages})
	}
	return
}
func (msg *PWSCancelRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *PWSCancelRequest) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_MessageIdentifier:
		var tmp MessageIdentifier
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.MessageIdentifier = &tmp
	case ProtocolIEID_SerialNumber:
		var tmp SerialNumber
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SerialNumber = &tmp
	case ProtocolIEID_WarningAreaList:
		var tmp WarningAreaList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.WarningAreaList = &tmp
	case ProtocolIEID_CancelAllWarningMessages:
		var tmp CancelAllWarningMessages
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.CancelAllWarningMessages = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
