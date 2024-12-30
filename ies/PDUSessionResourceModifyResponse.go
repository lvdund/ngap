package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type PDUSessionResourceModifyResponse struct {
	AMFUENGAPID                                int64
	RANUENGAPID                                int64
	PDUSessionResourceModifyListModRes         *[]PDUSessionResourceModifyItemModRes
	PDUSessionResourceFailedToModifyListModRes *[]PDUSessionResourceFailedToModifyItemModRes
	UserLocationInformation                    *UserLocationInformation
	CriticalityDiagnostics                     *CriticalityDiagnostics
}

func (msg *PDUSessionResourceModifyResponse) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_PDUSessionResourceModify, Criticality_PresentReject, msg.toIes())
}
func (msg *PDUSessionResourceModifyResponse) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 1099511627775},
			ext:   false,
			Value: aper.Integer(msg.AMFUENGAPID),
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4294967295},
			ext:   false,
			Value: aper.Integer(msg.RANUENGAPID),
		}})
	if msg.PDUSessionResourceModifyListModRes != nil {
		tmp_PDUSessionResourceModifyListModRes := Sequence[*PDUSessionResourceModifyItemModRes]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range *msg.PDUSessionResourceModifyListModRes {
			tmp_PDUSessionResourceModifyListModRes.Value = append(tmp_PDUSessionResourceModifyListModRes.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceModifyListModRes},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_PDUSessionResourceModifyListModRes,
		})
	}
	if msg.PDUSessionResourceFailedToModifyListModRes != nil {
		tmp_PDUSessionResourceFailedToModifyListModRes := Sequence[*PDUSessionResourceFailedToModifyItemModRes]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range *msg.PDUSessionResourceFailedToModifyListModRes {
			tmp_PDUSessionResourceFailedToModifyListModRes.Value = append(tmp_PDUSessionResourceFailedToModifyListModRes.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceFailedToModifyListModRes},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_PDUSessionResourceFailedToModifyListModRes,
		})
	}
	if msg.UserLocationInformation != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UserLocationInformation},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UserLocationInformation,
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
func (msg *PDUSessionResourceModifyResponse) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *PDUSessionResourceModifyResponse) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_PDUSessionResourceModifyListModRes:
		tmp := Sequence[*PDUSessionResourceModifyItemModRes]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceModifyListModRes = &[]PDUSessionResourceModifyItemModRes{}
		for _, i := range tmp.Value {
			*msg.PDUSessionResourceModifyListModRes = append(*msg.PDUSessionResourceModifyListModRes, *i)
		}
	case ProtocolIEID_PDUSessionResourceFailedToModifyListModRes:
		tmp := Sequence[*PDUSessionResourceFailedToModifyItemModRes]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceFailedToModifyListModRes = &[]PDUSessionResourceFailedToModifyItemModRes{}
		for _, i := range tmp.Value {
			*msg.PDUSessionResourceFailedToModifyListModRes = append(*msg.PDUSessionResourceFailedToModifyListModRes, *i)
		}
	case ProtocolIEID_UserLocationInformation:
		var tmp UserLocationInformation
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UserLocationInformation = &tmp
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
