package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type OverloadStart struct {
	AMFOverloadResponse               *OverloadResponse               `,reject,optional`
	AMFTrafficLoadReductionIndication *TrafficLoadReductionIndication `,ignore,optional`
	OverloadStartNSSAIList            *OverloadStartNSSAIList         `,ignore,optional`
}

func (msg *OverloadStart) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_OverloadStart, Criticality_PresentIgnore, msg.toIes())
}
func (msg *OverloadStart) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.AMFOverloadResponse != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFOverloadResponse},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.AMFOverloadResponse})
	}
	if msg.AMFTrafficLoadReductionIndication != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFTrafficLoadReductionIndication},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.AMFTrafficLoadReductionIndication})
	}
	if msg.OverloadStartNSSAIList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_OverloadStartNSSAIList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.OverloadStartNSSAIList})
	}
	return
}
func (msg *OverloadStart) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *OverloadStart) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_AMFOverloadResponse:
		var tmp OverloadResponse
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFOverloadResponse = &tmp
	case ProtocolIEID_AMFTrafficLoadReductionIndication:
		var tmp TrafficLoadReductionIndication
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFTrafficLoadReductionIndication = &tmp
	case ProtocolIEID_OverloadStartNSSAIList:
		var tmp OverloadStartNSSAIList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.OverloadStartNSSAIList = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
