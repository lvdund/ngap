package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type RerouteNASRequest struct {
	RANUENGAPID                         *RANUENGAPID                         `,reject,mandatory`
	AMFUENGAPID                         *AMFUENGAPID                         `,ignore,optional`
	NGAPMessage                         *NGAPMessage                         `,reject,mandatory`
	AMFSetID                            *AMFSetID                            `,reject,mandatory`
	AllowedNSSAI                        *AllowedNSSAI                        `,reject,optional`
	SourceToTargetAMFInformationReroute *SourceToTargetAMFInformationReroute `,ignore,optional`
}

func (msg *RerouteNASRequest) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_RerouteNASRequest, Criticality_PresentReject, msg.toIes())
}
func (msg *RerouteNASRequest) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.RANUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.RANUENGAPID})
	}
	if msg.AMFUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.AMFUENGAPID})
	}
	if msg.NGAPMessage != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NGAPMessage},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.NGAPMessage})
	}
	if msg.AMFSetID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFSetID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.AMFSetID})
	}
	if msg.AllowedNSSAI != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AllowedNSSAI},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.AllowedNSSAI})
	}
	if msg.SourceToTargetAMFInformationReroute != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SourceToTargetAMFInformationReroute},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.SourceToTargetAMFInformationReroute})
	}
	return
}
func (msg *RerouteNASRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *RerouteNASRequest) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_RANUENGAPID:
		var tmp RANUENGAPID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RANUENGAPID = &tmp
	case ProtocolIEID_AMFUENGAPID:
		var tmp AMFUENGAPID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFUENGAPID = &tmp
	case ProtocolIEID_NGAPMessage:
		var tmp NGAPMessage
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NGAPMessage = &tmp
	case ProtocolIEID_AMFSetID:
		var tmp AMFSetID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFSetID = &tmp
	case ProtocolIEID_AllowedNSSAI:
		var tmp AllowedNSSAI
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AllowedNSSAI = &tmp
	case ProtocolIEID_SourceToTargetAMFInformationReroute:
		var tmp SourceToTargetAMFInformationReroute
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SourceToTargetAMFInformationReroute = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
