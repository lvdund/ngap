package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type PDUSessionResourceNotify struct {
	AMFUENGAPID                       int64
	RANUENGAPID                       int64
	PDUSessionResourceNotifyList      *[]PDUSessionResourceNotifyItem
	PDUSessionResourceReleasedListNot *[]PDUSessionResourceReleasedItemNot
	UserLocationInformation           *UserLocationInformation
}

func (msg *PDUSessionResourceNotify) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_PDUSessionResourceNotify, Criticality_PresentIgnore, msg.toIes())
}
func (msg *PDUSessionResourceNotify) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 1099511627775},
			ext:   false,
			Value: aper.Integer(msg.AMFUENGAPID),
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4294967295},
			ext:   false,
			Value: aper.Integer(msg.RANUENGAPID),
		}})
	if msg.PDUSessionResourceNotifyList != nil {
		tmp_PDUSessionResourceNotifyList := Sequence[*PDUSessionResourceNotifyItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range *msg.PDUSessionResourceNotifyList {
			tmp_PDUSessionResourceNotifyList.Value = append(tmp_PDUSessionResourceNotifyList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceNotifyList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_PDUSessionResourceNotifyList,
		})
	}
	if msg.PDUSessionResourceReleasedListNot != nil {
		tmp_PDUSessionResourceReleasedListNot := Sequence[*PDUSessionResourceReleasedItemNot]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range *msg.PDUSessionResourceReleasedListNot {
			tmp_PDUSessionResourceReleasedListNot.Value = append(tmp_PDUSessionResourceReleasedListNot.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceReleasedListNot},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_PDUSessionResourceReleasedListNot,
		})
	}
	if msg.UserLocationInformation != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UserLocationInformation},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UserLocationInformation,
		})
	}
	return
}
func (msg *PDUSessionResourceNotify) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *PDUSessionResourceNotify) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_AMFUENGAPID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 1099511627775},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFUENGAPID = int64(tmp.Value)
	case ProtocolIEID_RANUENGAPID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4294967295},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RANUENGAPID = int64(tmp.Value)
	case ProtocolIEID_PDUSessionResourceNotifyList:
		tmp := Sequence[*PDUSessionResourceNotifyItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceNotifyList = &[]PDUSessionResourceNotifyItem{}
		for _, i := range tmp.Value {
			*msg.PDUSessionResourceNotifyList = append(*msg.PDUSessionResourceNotifyList, *i)
		}
	case ProtocolIEID_PDUSessionResourceReleasedListNot:
		tmp := Sequence[*PDUSessionResourceReleasedItemNot]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceReleasedListNot = &[]PDUSessionResourceReleasedItemNot{}
		for _, i := range tmp.Value {
			*msg.PDUSessionResourceReleasedListNot = append(*msg.PDUSessionResourceReleasedListNot, *i)
		}
	case ProtocolIEID_UserLocationInformation:
		var tmp UserLocationInformation
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UserLocationInformation = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
