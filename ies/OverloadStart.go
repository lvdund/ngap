package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type OverloadStart struct {
	AMFOverloadResponse               *OverloadResponse        `optional,mandatory,reject`
	AMFTrafficLoadReductionIndication *int64                   `lb:1,ub:99,optional,mandatory,ignore`
	OverloadStartNSSAIList            []OverloadStartNSSAIItem `lb:1,ub:maxnoofSliceItems,optional,mandatory,ignore`
}

func (msg *OverloadStart) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		return
	}
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_OverloadStart, Criticality_PresentIgnore, ies)
}
func (msg *OverloadStart) toIes() (ies []NgapMessageIE, err error) {
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
	if len(msg.OverloadStartNSSAIList) > 0 {
		tmp_OverloadStartNSSAIList := Sequence[*OverloadStartNSSAIItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSliceItems},
			ext: false,
		}
		for _, i := range msg.OverloadStartNSSAIList {
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
func (msg *OverloadStart) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("OverloadStart"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := OverloadStartDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*NgapMessageIE),
	}
	if _, err = aper.ReadSequenceOf[NgapMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	return
}

type OverloadStartDecoder struct {
	msg      *OverloadStart
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *OverloadStartDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_AMFOverloadResponse:
		var tmp OverloadResponse
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read AMFOverloadResponse", err)
			return
		}
		msg.AMFOverloadResponse = &tmp
	case ProtocolIEID_AMFTrafficLoadReductionIndication:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 99},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read AMFTrafficLoadReductionIndication", err)
			return
		}
		*msg.AMFTrafficLoadReductionIndication = int64(tmp.Value)
	case ProtocolIEID_OverloadStartNSSAIList:
		tmp := Sequence[*OverloadStartNSSAIItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSliceItems},
			ext: false,
		}
		fn := func() *OverloadStartNSSAIItem { return new(OverloadStartNSSAIItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read OverloadStartNSSAIList", err)
			return
		}
		msg.OverloadStartNSSAIList = []OverloadStartNSSAIItem{}
		for _, i := range tmp.Value {
			msg.OverloadStartNSSAIList = append(msg.OverloadStartNSSAIList, *i)
		}
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
