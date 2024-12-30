package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type PDUSessionResourceModifyConfirm struct {
	AMFUENGAPID                                int64
	RANUENGAPID                                int64
	PDUSessionResourceModifyListModCfm         *[]PDUSessionResourceModifyItemModCfm
	PDUSessionResourceFailedToModifyListModCfm *[]PDUSessionResourceFailedToModifyItemModCfm
	CriticalityDiagnostics                     *CriticalityDiagnostics
}

func (msg *PDUSessionResourceModifyConfirm) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_PDUSessionResourceModifyIndication, Criticality_PresentReject, msg.toIes())
}
func (msg *PDUSessionResourceModifyConfirm) toIes() (ies []NgapMessageIE) {
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
	if msg.PDUSessionResourceModifyListModCfm != nil {
		tmp_PDUSessionResourceModifyListModCfm := Sequence[*PDUSessionResourceModifyItemModCfm]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range *msg.PDUSessionResourceModifyListModCfm {
			tmp_PDUSessionResourceModifyListModCfm.Value = append(tmp_PDUSessionResourceModifyListModCfm.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceModifyListModCfm},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_PDUSessionResourceModifyListModCfm,
		})
	}
	if msg.PDUSessionResourceFailedToModifyListModCfm != nil {
		tmp_PDUSessionResourceFailedToModifyListModCfm := Sequence[*PDUSessionResourceFailedToModifyItemModCfm]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range *msg.PDUSessionResourceFailedToModifyListModCfm {
			tmp_PDUSessionResourceFailedToModifyListModCfm.Value = append(tmp_PDUSessionResourceFailedToModifyListModCfm.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceFailedToModifyListModCfm},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_PDUSessionResourceFailedToModifyListModCfm,
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
func (msg *PDUSessionResourceModifyConfirm) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *PDUSessionResourceModifyConfirm) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_PDUSessionResourceModifyListModCfm:
		tmp := Sequence[*PDUSessionResourceModifyItemModCfm]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceModifyListModCfm = &[]PDUSessionResourceModifyItemModCfm{}
		for _, i := range tmp.Value {
			*msg.PDUSessionResourceModifyListModCfm = append(*msg.PDUSessionResourceModifyListModCfm, *i)
		}
	case ProtocolIEID_PDUSessionResourceFailedToModifyListModCfm:
		tmp := Sequence[*PDUSessionResourceFailedToModifyItemModCfm]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceFailedToModifyListModCfm = &[]PDUSessionResourceFailedToModifyItemModCfm{}
		for _, i := range tmp.Value {
			*msg.PDUSessionResourceFailedToModifyListModCfm = append(*msg.PDUSessionResourceFailedToModifyListModCfm, *i)
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
