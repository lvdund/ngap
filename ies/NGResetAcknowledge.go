package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type NGResetAcknowledge struct {
	UEassociatedLogicalNGconnectionList *[]UEassociatedLogicalNGconnectionItem
	CriticalityDiagnostics              *CriticalityDiagnostics
}

func (msg *NGResetAcknowledge) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_NGReset, Criticality_PresentReject, msg.toIes())
}
func (msg *NGResetAcknowledge) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.UEassociatedLogicalNGconnectionList != nil {
		tmp_UEassociatedLogicalNGconnectionList := Sequence[*UEassociatedLogicalNGconnectionItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofNGConnectionsToReset},
			ext: false,
		}
		for _, i := range *msg.UEassociatedLogicalNGconnectionList {
			tmp_UEassociatedLogicalNGconnectionList.Value = append(tmp_UEassociatedLogicalNGconnectionList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UEassociatedLogicalNGconnectionList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_UEassociatedLogicalNGconnectionList,
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
func (msg *NGResetAcknowledge) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *NGResetAcknowledge) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_UEassociatedLogicalNGconnectionList:
		tmp := Sequence[*UEassociatedLogicalNGconnectionItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofNGConnectionsToReset},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UEassociatedLogicalNGconnectionList = &[]UEassociatedLogicalNGconnectionItem{}
		for _, i := range tmp.Value {
			*msg.UEassociatedLogicalNGconnectionList = append(*msg.UEassociatedLogicalNGconnectionList, *i)
		}
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
