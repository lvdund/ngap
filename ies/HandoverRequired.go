package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type HandoverRequired struct {
	AMFUENGAPID                        int64
	RANUENGAPID                        int64
	HandoverType                       HandoverType
	Cause                              Cause
	TargetID                           TargetID
	DirectForwardingPathAvailability   *DirectForwardingPathAvailability
	PDUSessionResourceListHORqd        []PDUSessionResourceItemHORqd
	SourceToTargetTransparentContainer []byte
}

func (msg *HandoverRequired) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_HandoverPreparation, Criticality_PresentReject, msg.toIes())
}
func (msg *HandoverRequired) toIes() (ies []NgapMessageIE) {
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
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_HandoverType},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.HandoverType,
	})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_Cause},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &msg.Cause,
	})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TargetID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.TargetID,
	})
	if msg.DirectForwardingPathAvailability != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DirectForwardingPathAvailability},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.DirectForwardingPathAvailability,
		})
	}
	tmp_PDUSessionResourceListHORqd := Sequence[*PDUSessionResourceItemHORqd]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
		ext: false,
	}
	for _, i := range msg.PDUSessionResourceListHORqd {
		tmp_PDUSessionResourceListHORqd.Value = append(tmp_PDUSessionResourceListHORqd.Value, &i)
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceListHORqd},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &tmp_PDUSessionResourceListHORqd,
	})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_SourceToTargetTransparentContainer},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &OCTETSTRING{
			c:     aper.Constraint{Lb: 0, Ub: 0},
			ext:   false,
			Value: msg.SourceToTargetTransparentContainer,
		}})
	return
}
func (msg *HandoverRequired) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *HandoverRequired) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_HandoverType:
		var tmp HandoverType
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.HandoverType = tmp
	case ProtocolIEID_Cause:
		var tmp Cause
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.Cause = tmp
	case ProtocolIEID_TargetID:
		var tmp TargetID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.TargetID = tmp
	case ProtocolIEID_DirectForwardingPathAvailability:
		var tmp DirectForwardingPathAvailability
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.DirectForwardingPathAvailability = &tmp
	case ProtocolIEID_PDUSessionResourceListHORqd:
		tmp := Sequence[*PDUSessionResourceItemHORqd]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceListHORqd = []PDUSessionResourceItemHORqd{}
		for _, i := range tmp.Value {
			msg.PDUSessionResourceListHORqd = append(msg.PDUSessionResourceListHORqd, *i)
		}
	case ProtocolIEID_SourceToTargetTransparentContainer:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SourceToTargetTransparentContainer = tmp.Value
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
