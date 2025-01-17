package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type Paging struct {
	UEPagingIdentity           UEPagingIdentity            `mandatory,ignore`
	PagingDRX                  *PagingDRX                  `optional,ignore`
	TAIListForPaging           []TAIListForPagingItem      `lb:1,ub:maxnoofTAIforPaging,mandatory,ignore`
	PagingPriority             *PagingPriority             `optional,ignore`
	UERadioCapabilityForPaging *UERadioCapabilityForPaging `optional,ignore`
	PagingOrigin               *PagingOrigin               `optional,ignore`
	AssistanceDataForPaging    *AssistanceDataForPaging    `optional,ignore`
}

func (msg *Paging) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		return
	}
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_Paging, Criticality_PresentIgnore, ies)
}
func (msg *Paging) toIes() (ies []NgapMessageIE, err error) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_UEPagingIdentity},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &msg.UEPagingIdentity,
	})
	if msg.PagingDRX != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PagingDRX},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.PagingDRX,
		})
	}
	if len(msg.TAIListForPaging) > 0 {
		tmp_TAIListForPaging := Sequence[*TAIListForPagingItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIforPaging},
			ext: false,
		}
		for _, i := range msg.TAIListForPaging {
			tmp_TAIListForPaging.Value = append(tmp_TAIListForPaging.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TAIListForPaging},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_TAIListForPaging,
		})
	} else {
		err = utils.WrapError("TAIListForPaging is nil", err)
		return
	}
	if msg.PagingPriority != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PagingPriority},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.PagingPriority,
		})
	}
	if msg.UERadioCapabilityForPaging != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UERadioCapabilityForPaging},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UERadioCapabilityForPaging,
		})
	}
	if msg.PagingOrigin != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PagingOrigin},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.PagingOrigin,
		})
	}
	if msg.AssistanceDataForPaging != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AssistanceDataForPaging},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.AssistanceDataForPaging,
		})
	}
	return
}
func (msg *Paging) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("Paging"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := PagingDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*NgapMessageIE),
	}
	if _, err = aper.ReadSequenceOf[NgapMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	if _, ok := decoder.list[ProtocolIEID_UEPagingIdentity]; !ok {
		err = fmt.Errorf("Mandatory field UEPagingIdentity is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_UEPagingIdentity},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_TAIListForPaging]; !ok {
		err = fmt.Errorf("Mandatory field TAIListForPaging is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_TAIListForPaging},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type PagingDecoder struct {
	msg      *Paging
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *PagingDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_UEPagingIdentity:
		var tmp UEPagingIdentity
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read UEPagingIdentity", err)
			return
		}
		msg.UEPagingIdentity = tmp
	case ProtocolIEID_PagingDRX:
		var tmp PagingDRX
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read PagingDRX", err)
			return
		}
		msg.PagingDRX = &tmp
	case ProtocolIEID_TAIListForPaging:
		tmp := Sequence[*TAIListForPagingItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIforPaging},
			ext: false,
		}
		fn := func() *TAIListForPagingItem { return new(TAIListForPagingItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read TAIListForPaging", err)
			return
		}
		msg.TAIListForPaging = []TAIListForPagingItem{}
		for _, i := range tmp.Value {
			msg.TAIListForPaging = append(msg.TAIListForPaging, *i)
		}
	case ProtocolIEID_PagingPriority:
		var tmp PagingPriority
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read PagingPriority", err)
			return
		}
		msg.PagingPriority = &tmp
	case ProtocolIEID_UERadioCapabilityForPaging:
		var tmp UERadioCapabilityForPaging
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read UERadioCapabilityForPaging", err)
			return
		}
		msg.UERadioCapabilityForPaging = &tmp
	case ProtocolIEID_PagingOrigin:
		var tmp PagingOrigin
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read PagingOrigin", err)
			return
		}
		msg.PagingOrigin = &tmp
	case ProtocolIEID_AssistanceDataForPaging:
		var tmp AssistanceDataForPaging
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read AssistanceDataForPaging", err)
			return
		}
		msg.AssistanceDataForPaging = &tmp
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
