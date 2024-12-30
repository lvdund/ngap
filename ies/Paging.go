package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type Paging struct {
	UEPagingIdentity           UEPagingIdentity
	PagingDRX                  *PagingDRX
	TAIListForPaging           []TAIListForPagingItem
	PagingPriority             *PagingPriority
	UERadioCapabilityForPaging *UERadioCapabilityForPaging
	PagingOrigin               *PagingOrigin
	AssistanceDataForPaging    *AssistanceDataForPaging
}

func (msg *Paging) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_Paging, Criticality_PresentIgnore, msg.toIes())
}
func (msg *Paging) toIes() (ies []NgapMessageIE) {
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
func (msg *Paging) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *Paging) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_UEPagingIdentity:
		var tmp UEPagingIdentity
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UEPagingIdentity = tmp
	case ProtocolIEID_PagingDRX:
		var tmp PagingDRX
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PagingDRX = &tmp
	case ProtocolIEID_TAIListForPaging:
		tmp := Sequence[*TAIListForPagingItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIforPaging},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.TAIListForPaging = []TAIListForPagingItem{}
		for _, i := range tmp.Value {
			msg.TAIListForPaging = append(msg.TAIListForPaging, *i)
		}
	case ProtocolIEID_PagingPriority:
		var tmp PagingPriority
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PagingPriority = &tmp
	case ProtocolIEID_UERadioCapabilityForPaging:
		var tmp UERadioCapabilityForPaging
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UERadioCapabilityForPaging = &tmp
	case ProtocolIEID_PagingOrigin:
		var tmp PagingOrigin
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PagingOrigin = &tmp
	case ProtocolIEID_AssistanceDataForPaging:
		var tmp AssistanceDataForPaging
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AssistanceDataForPaging = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
