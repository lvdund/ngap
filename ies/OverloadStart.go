package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type OverloadStart struct {
	AMFOverloadResponse               *OverloadResponse
	AMFTrafficLoadReductionIndication *int64
	OverloadStartNSSAIList            *[]OverloadStartNSSAIItem
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
			Value:       msg.AMFOverloadResponse,
		})
	}
	if msg.AMFTrafficLoadReductionIndication != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFTrafficLoadReductionIndication},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 1, Ub: 99},
				ext:   false,
				Value: aper.Integer(*msg.AMFTrafficLoadReductionIndication),
			}})
	}
	if msg.OverloadStartNSSAIList != nil {
		tmp_OverloadStartNSSAIList := Sequence[*OverloadStartNSSAIItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSliceItems},
			ext: false,
		}
		for _, i := range *msg.OverloadStartNSSAIList {
			tmp_OverloadStartNSSAIList.Value = append(tmp_OverloadStartNSSAIList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_OverloadStartNSSAIList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_OverloadStartNSSAIList,
		})
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
	var id int64
	var c uint64
	var buf []byte
	if id, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: int64(aper.POW_16) - 1}, false); err != nil {
		return
	}
	msgIe = new(NgapMessageIE)
	msgIe.Id.Value = aper.Integer(id)
	if c, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return
	}
	msgIe.Criticality.Value = aper.Enumerated(c)
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
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 99},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.AMFTrafficLoadReductionIndication = int64(tmp.Value)
	case ProtocolIEID_OverloadStartNSSAIList:
		tmp := Sequence[*OverloadStartNSSAIItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSliceItems},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.OverloadStartNSSAIList = &[]OverloadStartNSSAIItem{}
		for _, i := range tmp.Value {
			*msg.OverloadStartNSSAIList = append(*msg.OverloadStartNSSAIList, *i)
		}
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
