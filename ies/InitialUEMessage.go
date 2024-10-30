package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type InitialUEMessage struct {
	RANUENGAPID                         *RANUENGAPID                         `,reject,mandatory`
	NASPDU                              *NASPDU                              `,reject,mandatory`
	UserLocationInformation             *UserLocationInformation             `,reject,mandatory`
	RRCEstablishmentCause               *RRCEstablishmentCause               `,ignore,mandatory`
	FiveGSTMSI                          *FiveGSTMSI                          `,reject,optional`
	AMFSetID                            *AMFSetID                            `,ignore,optional`
	UEContextRequest                    *UEContextRequest                    `,ignore,optional`
	AllowedNSSAI                        *AllowedNSSAI                        `,reject,optional`
	SourceToTargetAMFInformationReroute *SourceToTargetAMFInformationReroute `,ignore,optional`
	SelectedPLMNIdentity                *PLMNIdentity                        `,ignore,optional`
}

func (msg *InitialUEMessage) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_InitialUEMessage, Criticality_PresentIgnore, msg.toIes())
}
func (msg *InitialUEMessage) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.RANUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.RANUENGAPID})
	}
	if msg.NASPDU != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NASPDU},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.NASPDU})
	}
	if msg.UserLocationInformation != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UserLocationInformation},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.UserLocationInformation})
	}
	if msg.RRCEstablishmentCause != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RRCEstablishmentCause},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RRCEstablishmentCause})
	}
	if msg.FiveGSTMSI != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_FiveGSTMSI},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.FiveGSTMSI})
	}
	if msg.AMFSetID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFSetID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.AMFSetID})
	}
	if msg.UEContextRequest != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UEContextRequest},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UEContextRequest})
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
	if msg.SelectedPLMNIdentity != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SelectedPLMNIdentity},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.SelectedPLMNIdentity})
	}
	return
}
func (msg *InitialUEMessage) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *InitialUEMessage) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_NASPDU:
		var tmp NASPDU
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NASPDU = &tmp
	case ProtocolIEID_UserLocationInformation:
		var tmp UserLocationInformation
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UserLocationInformation = &tmp
	case ProtocolIEID_RRCEstablishmentCause:
		var tmp RRCEstablishmentCause
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RRCEstablishmentCause = &tmp
	case ProtocolIEID_FiveGSTMSI:
		var tmp FiveGSTMSI
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.FiveGSTMSI = &tmp
	case ProtocolIEID_AMFSetID:
		var tmp AMFSetID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFSetID = &tmp
	case ProtocolIEID_UEContextRequest:
		var tmp UEContextRequest
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UEContextRequest = &tmp
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
	case ProtocolIEID_SelectedPLMNIdentity:
		var tmp PLMNIdentity
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SelectedPLMNIdentity = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
