package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PWSCancelResponse struct {
	MessageIdentifier          aper.BitString              `lb:16,ub:16,mandatory,reject`
	SerialNumber               aper.BitString              `lb:16,ub:16,mandatory,reject`
	BroadcastCancelledAreaList *BroadcastCancelledAreaList `optional,ignore`
	CriticalityDiagnostics     *CriticalityDiagnostics     `optional,ignore`
}

func (msg *PWSCancelResponse) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("PWSCancelResponse"), err)
		return
	}
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_PWSCancel, Criticality_PresentReject, ies)
}
func (msg *PWSCancelResponse) toIes() (ies []NgapMessageIE, err error) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_MessageIdentifier},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &BITSTRING{
			c:   aper.Constraint{Lb: 16, Ub: 16},
			ext: false,
			Value: aper.BitString{
				Bytes: msg.MessageIdentifier.Bytes, NumBits: msg.MessageIdentifier.NumBits},
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_SerialNumber},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &BITSTRING{
			c:   aper.Constraint{Lb: 16, Ub: 16},
			ext: false,
			Value: aper.BitString{
				Bytes: msg.SerialNumber.Bytes, NumBits: msg.SerialNumber.NumBits},
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
func (msg *PWSCancelResponse) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("PWSCancelResponse"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := PWSCancelResponseDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*NgapMessageIE),
	}
	if _, err = aper.ReadSequenceOf[NgapMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	if _, ok := decoder.list[ProtocolIEID_MessageIdentifier]; !ok {
		err = fmt.Errorf("Mandatory field MessageIdentifier is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_MessageIdentifier},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_SerialNumber]; !ok {
		err = fmt.Errorf("Mandatory field SerialNumber is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_SerialNumber},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type PWSCancelResponseDecoder struct {
	msg      *PWSCancelResponse
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *PWSCancelResponseDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_MessageIdentifier:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 16, Ub: 16},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read MessageIdentifier", err)
			return
		}
		msg.MessageIdentifier = aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	case ProtocolIEID_SerialNumber:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 16, Ub: 16},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read SerialNumber", err)
			return
		}
		msg.SerialNumber = aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	case ProtocolIEID_BroadcastCancelledAreaList:
		var tmp BroadcastCancelledAreaList
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read BroadcastCancelledAreaList", err)
			return
		}
		msg.BroadcastCancelledAreaList = &tmp
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
