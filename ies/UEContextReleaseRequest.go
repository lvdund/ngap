package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type UEContextReleaseRequest struct {
	AMFUENGAPID                     int64
	RANUENGAPID                     int64
	PDUSessionResourceListCxtRelReq *[]PDUSessionResourceItemCxtRelReq
	Cause                           Cause
}

func (msg *UEContextReleaseRequest) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_UEContextReleaseRequest, Criticality_PresentIgnore, msg.toIes())
}
func (msg *UEContextReleaseRequest) toIes() (ies []NgapMessageIE) {
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
	if msg.PDUSessionResourceListCxtRelReq != nil {
		tmp_PDUSessionResourceListCxtRelReq := Sequence[*PDUSessionResourceItemCxtRelReq]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range *msg.PDUSessionResourceListCxtRelReq {
			tmp_PDUSessionResourceListCxtRelReq.Value = append(tmp_PDUSessionResourceListCxtRelReq.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceListCxtRelReq},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_PDUSessionResourceListCxtRelReq,
		})
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_Cause},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &msg.Cause,
	})
	return
}
func (msg *UEContextReleaseRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *UEContextReleaseRequest) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_PDUSessionResourceListCxtRelReq:
		tmp := Sequence[*PDUSessionResourceItemCxtRelReq]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceListCxtRelReq = &[]PDUSessionResourceItemCxtRelReq{}
		for _, i := range tmp.Value {
			*msg.PDUSessionResourceListCxtRelReq = append(*msg.PDUSessionResourceListCxtRelReq, *i)
		}
	case ProtocolIEID_Cause:
		var tmp Cause
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.Cause = tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
