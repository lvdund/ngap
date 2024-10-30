package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type WriteReplaceWarningResponse struct {
	MessageIdentifier          *MessageIdentifier          `,reject,mandatory`
	SerialNumber               *SerialNumber               `,reject,mandatory`
	BroadcastCompletedAreaList *BroadcastCompletedAreaList `,ignore,optional`
	CriticalityDiagnostics     *CriticalityDiagnostics     `,ignore,optional`
}

func (msg *WriteReplaceWarningResponse) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_WriteReplaceWarning, Criticality_PresentReject, msg.toIes())
}
func (msg *WriteReplaceWarningResponse) toIes() (ies []NgapMessageIE) {
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
	if msg.BroadcastCompletedAreaList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_BroadcastCompletedAreaList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.BroadcastCompletedAreaList})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics})
	}
	return
}
func (msg *WriteReplaceWarningResponse) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *WriteReplaceWarningResponse) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_BroadcastCompletedAreaList:
		var tmp BroadcastCompletedAreaList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.BroadcastCompletedAreaList = &tmp
	case ProtocolIEID_CriticalityDiagnostics:
		var tmp CriticalityDiagnostics
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.CriticalityDiagnostics = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
