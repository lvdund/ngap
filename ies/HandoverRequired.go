package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type HandoverRequired struct {
	AMFUENGAPID                        *AMFUENGAPID                        `,reject,mandatory`
	RANUENGAPID                        *RANUENGAPID                        `,reject,mandatory`
	HandoverType                       *HandoverType                       `,reject,mandatory`
	Cause                              *Cause                              `,ignore,mandatory`
	TargetID                           *TargetID                           `,reject,mandatory`
	DirectForwardingPathAvailability   *DirectForwardingPathAvailability   `,ignore,optional`
	PDUSessionResourceListHORqd        *PDUSessionResourceListHORqd        `,reject,mandatory`
	SourceToTargetTransparentContainer *SourceToTargetTransparentContainer `,reject,mandatory`
}

func (msg *HandoverRequired) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_HandoverPreparation, Criticality_PresentReject, msg.toIes())
}
func (msg *HandoverRequired) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.AMFUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.AMFUENGAPID})
	}
	if msg.RANUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.RANUENGAPID})
	}
	if msg.HandoverType != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_HandoverType},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.HandoverType})
	}
	if msg.Cause != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_Cause},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.Cause})
	}
	if msg.TargetID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TargetID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.TargetID})
	}
	if msg.DirectForwardingPathAvailability != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DirectForwardingPathAvailability},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.DirectForwardingPathAvailability})
	}
	if msg.PDUSessionResourceListHORqd != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceListHORqd},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.PDUSessionResourceListHORqd})
	}
	if msg.SourceToTargetTransparentContainer != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SourceToTargetTransparentContainer},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.SourceToTargetTransparentContainer})
	}
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
	case ProtocolIEID_HandoverType:
		var tmp HandoverType
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.HandoverType = &tmp
	case ProtocolIEID_Cause:
		var tmp Cause
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.Cause = &tmp
	case ProtocolIEID_TargetID:
		var tmp TargetID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.TargetID = &tmp
	case ProtocolIEID_DirectForwardingPathAvailability:
		var tmp DirectForwardingPathAvailability
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.DirectForwardingPathAvailability = &tmp
	case ProtocolIEID_PDUSessionResourceListHORqd:
		var tmp PDUSessionResourceListHORqd
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceListHORqd = &tmp
	case ProtocolIEID_SourceToTargetTransparentContainer:
		var tmp SourceToTargetTransparentContainer
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SourceToTargetTransparentContainer = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
