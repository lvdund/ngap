package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type NGResetAcknowledge struct {
	UEassociatedLogicalNGconnectionList []UEassociatedLogicalNGconnectionItem `lb:1,ub:maxnoofNGConnectionsToReset,optional,ignore`
	CriticalityDiagnostics              *CriticalityDiagnostics               `optional,ignore`
}

func (msg *NGResetAcknowledge) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		return
	}
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_NGReset, Criticality_PresentReject, ies)
}
func (msg *NGResetAcknowledge) toIes() (ies []NgapMessageIE, err error) {
	ies = []NgapMessageIE{}
	if len(msg.UEassociatedLogicalNGconnectionList) > 0 {
		tmp_UEassociatedLogicalNGconnectionList := Sequence[*UEassociatedLogicalNGconnectionItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofNGConnectionsToReset},
			ext: false,
		}
		for _, i := range msg.UEassociatedLogicalNGconnectionList {
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
func (msg *NGResetAcknowledge) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("NGResetAcknowledge"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := NGResetAcknowledgeDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*NgapMessageIE),
	}
	if _, err = aper.ReadSequenceOf[NgapMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	return
}

type NGResetAcknowledgeDecoder struct {
	msg      *NGResetAcknowledge
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *NGResetAcknowledgeDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	ieId := msgIe.Id.Value
	if _, ok := decoder.list[ieId]; ok {
		err = fmt.Errorf("Duplicated protocol IEID[%d] found", ieId)
		return
	}
	decoder.list[ieId] = msgIe
	ieR := aper.NewReader(bytes.NewReader(buf))
	msg := decoder.msg
	switch msgIe.Id.Value {
	case ProtocolIEID_UEassociatedLogicalNGconnectionList:
		tmp := Sequence[*UEassociatedLogicalNGconnectionItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofNGConnectionsToReset},
			ext: false,
		}
		fn := func() *UEassociatedLogicalNGconnectionItem { return new(UEassociatedLogicalNGconnectionItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read UEassociatedLogicalNGconnectionList", err)
			return
		}
		msg.UEassociatedLogicalNGconnectionList = []UEassociatedLogicalNGconnectionItem{}
		for _, i := range tmp.Value {
			msg.UEassociatedLogicalNGconnectionList = append(msg.UEassociatedLogicalNGconnectionList, *i)
		}
	case ProtocolIEID_CriticalityDiagnostics:
		var tmp CriticalityDiagnostics
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read CriticalityDiagnostics", err)
			return
		}
		msg.CriticalityDiagnostics = &tmp
	default:
		switch msgIe.Criticality.Value {
		case Criticality_PresentReject:
			fmt.Errorf("Not comprehended IE ID 0x%04x (criticality: reject)", msgIe.Id.Value)
		case Criticality_PresentIgnore:
			fmt.Errorf("Not comprehended IE ID 0x%04x (criticality: ignore)", msgIe.Id.Value)
		case Criticality_PresentNotify:
			fmt.Errorf("Not comprehended IE ID 0x%04x (criticality: notify)", msgIe.Id.Value)
		}
		if msgIe.Criticality.Value != Criticality_PresentIgnore {
			decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
				IECriticality: msgIe.Criticality,
				IEID:          msgIe.Id,
				TypeOfError:   TypeOfError{Value: TypeOfErrorNotunderstood},
			})
		}
	}
	return
}
