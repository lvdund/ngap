package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type HandoverCancelAcknowledge struct {
	AMFUENGAPID            *AMFUENGAPID            `,ignore,mandatory`
	RANUENGAPID            *RANUENGAPID            `,ignore,mandatory`
	CriticalityDiagnostics *CriticalityDiagnostics `,ignore,optional`
}

func (msg *HandoverCancelAcknowledge) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_HandoverCancel, Criticality_PresentReject, msg.toIes())
}
func (msg *HandoverCancelAcknowledge) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.AMFUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.AMFUENGAPID})
	}
	if msg.RANUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RANUENGAPID})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics})
	}
	return
}
func (msg *HandoverCancelAcknowledge) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *HandoverCancelAcknowledge) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_AMFUENGAPID:
		var tmp AMFUENGAPID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFUENGAPID = &tmp
	case ProtocolIEID_RANUENGAPID:
		var tmp RANUENGAPID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RANUENGAPID = &tmp
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
