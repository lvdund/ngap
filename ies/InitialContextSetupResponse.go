package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type InitialContextSetupResponse struct {
	AMFUENGAPID                               int64
	RANUENGAPID                               int64
	PDUSessionResourceSetupListCxtRes         *[]PDUSessionResourceSetupItemCxtRes
	PDUSessionResourceFailedToSetupListCxtRes *[]PDUSessionResourceFailedToSetupItemCxtRes
	CriticalityDiagnostics                    *CriticalityDiagnostics
}

func (msg *InitialContextSetupResponse) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_InitialContextSetup, Criticality_PresentReject, msg.toIes())
}
func (msg *InitialContextSetupResponse) toIes() (ies []NgapMessageIE) {
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
	if msg.PDUSessionResourceSetupListCxtRes != nil {
		tmp_PDUSessionResourceSetupListCxtRes := Sequence[*PDUSessionResourceSetupItemCxtRes]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range *msg.PDUSessionResourceSetupListCxtRes {
			tmp_PDUSessionResourceSetupListCxtRes.Value = append(tmp_PDUSessionResourceSetupListCxtRes.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceSetupListCxtRes},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_PDUSessionResourceSetupListCxtRes,
		})
	}
	if msg.PDUSessionResourceFailedToSetupListCxtRes != nil {
		tmp_PDUSessionResourceFailedToSetupListCxtRes := Sequence[*PDUSessionResourceFailedToSetupItemCxtRes]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range *msg.PDUSessionResourceFailedToSetupListCxtRes {
			tmp_PDUSessionResourceFailedToSetupListCxtRes.Value = append(tmp_PDUSessionResourceFailedToSetupListCxtRes.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceFailedToSetupListCxtRes},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_PDUSessionResourceFailedToSetupListCxtRes,
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
func (msg *InitialContextSetupResponse) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *InitialContextSetupResponse) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_PDUSessionResourceSetupListCxtRes:
		tmp := Sequence[*PDUSessionResourceSetupItemCxtRes]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceSetupListCxtRes = &[]PDUSessionResourceSetupItemCxtRes{}
		for _, i := range tmp.Value {
			*msg.PDUSessionResourceSetupListCxtRes = append(*msg.PDUSessionResourceSetupListCxtRes, *i)
		}
	case ProtocolIEID_PDUSessionResourceFailedToSetupListCxtRes:
		tmp := Sequence[*PDUSessionResourceFailedToSetupItemCxtRes]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceFailedToSetupListCxtRes = &[]PDUSessionResourceFailedToSetupItemCxtRes{}
		for _, i := range tmp.Value {
			*msg.PDUSessionResourceFailedToSetupListCxtRes = append(*msg.PDUSessionResourceFailedToSetupListCxtRes, *i)
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
