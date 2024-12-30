package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type PWSCancelResponse struct {
	MessageIdentifier          []byte
	SerialNumber               []byte
	BroadcastCancelledAreaList *BroadcastCancelledAreaList
	CriticalityDiagnostics     *CriticalityDiagnostics
}

func (msg *PWSCancelResponse) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_PWSCancel, Criticality_PresentReject, msg.toIes())
}
func (msg *PWSCancelResponse) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_MessageIdentifier},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &BITSTRING{
			Value: aper.BitString{
				Bytes:   msg.MessageIdentifier,
				NumBits: uint64(len(msg.MessageIdentifier))},
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_SerialNumber},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &BITSTRING{
			Value: aper.BitString{
				Bytes:   msg.SerialNumber,
				NumBits: uint64(len(msg.SerialNumber))},
		}})
	if msg.BroadcastCancelledAreaList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_BroadcastCancelledAreaList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.BroadcastCancelledAreaList,
		})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	return
}
func (msg *PWSCancelResponse) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *PWSCancelResponse) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_MessageIdentifier:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 16, Ub: 16},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.MessageIdentifier = tmp.Value.Bytes
	case ProtocolIEID_SerialNumber:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 16, Ub: 16},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SerialNumber = tmp.Value.Bytes
	case ProtocolIEID_BroadcastCancelledAreaList:
		var tmp BroadcastCancelledAreaList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.BroadcastCancelledAreaList = &tmp
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
