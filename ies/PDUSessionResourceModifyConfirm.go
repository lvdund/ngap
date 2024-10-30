package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type PDUSessionResourceModifyConfirm struct {
	AMFUENGAPID                                *AMFUENGAPID                                `,ignore,mandatory`
	RANUENGAPID                                *RANUENGAPID                                `,ignore,mandatory`
	PDUSessionResourceModifyListModCfm         *PDUSessionResourceModifyListModCfm         `,ignore,optional`
	PDUSessionResourceFailedToModifyListModCfm *PDUSessionResourceFailedToModifyListModCfm `,ignore,optional`
	CriticalityDiagnostics                     *CriticalityDiagnostics                     `,ignore,optional`
}

func (msg *PDUSessionResourceModifyConfirm) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_PDUSessionResourceModifyIndication, Criticality_PresentReject, msg.toIes())
}
func (msg *PDUSessionResourceModifyConfirm) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.AMFUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.AMFUENGAPID})
	}
	if msg.RANUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RANUENGAPID})
	}
	if msg.PDUSessionResourceModifyListModCfm != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceModifyListModCfm},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.PDUSessionResourceModifyListModCfm})
	}
	if msg.PDUSessionResourceFailedToModifyListModCfm != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceFailedToModifyListModCfm},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.PDUSessionResourceFailedToModifyListModCfm})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics})
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
	id, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: int64(aper.POW_16) - 1}, false)
	if err != nil {
		return
	}
	msgIe = new(NgapMessageIE)
	msgIe.Id.Value = aper.Integer(id)
	c, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	if err != nil {
		return
	}
	msgIe.Criticality.Value = aper.Enumerated(c)
	var buf []byte
	if buf, err = r.ReadOpenType(); err != nil {
		return
	}
	ieR := aper.NewReader(bytes.NewReader(buf))
	switch msgIe.Id.Value {
	case ProtocolIEID_AMFUENGAPID:
		var tmp AMFUENGAPID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFUENGAPID = &tmp
	case ProtocolIEID_RANUENGAPID:
		var tmp RANUENGAPID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RANUENGAPID = &tmp
	case ProtocolIEID_PDUSessionResourceModifyListModCfm:
		var tmp PDUSessionResourceModifyListModCfm
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceModifyListModCfm = &tmp
	case ProtocolIEID_PDUSessionResourceFailedToModifyListModCfm:
		var tmp PDUSessionResourceFailedToModifyListModCfm
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceFailedToModifyListModCfm = &tmp
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
