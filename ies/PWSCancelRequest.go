package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type PWSCancelRequest struct {
	MessageIdentifier        []byte
	SerialNumber             []byte
	WarningAreaList          *WarningAreaList
	CancelAllWarningMessages *CancelAllWarningMessages
}

func (msg *PWSCancelRequest) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_PWSCancel, Criticality_PresentReject, msg.toIes())
}
func (msg *PWSCancelRequest) toIes() (ies []NgapMessageIE) {
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
	if msg.WarningAreaList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_WarningAreaList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.WarningAreaList,
		})
	}
	if msg.CancelAllWarningMessages != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CancelAllWarningMessages},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.CancelAllWarningMessages,
		})
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
