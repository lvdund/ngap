package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type HandoverRequest struct {
	AMFUENGAPID                                 *AMFUENGAPID                                 `,reject,mandatory`
	HandoverType                                *HandoverType                                `,reject,mandatory`
	Cause                                       *Cause                                       `,ignore,mandatory`
	UEAggregateMaximumBitRate                   *UEAggregateMaximumBitRate                   `,reject,mandatory`
	CoreNetworkAssistanceInformationForInactive *CoreNetworkAssistanceInformationForInactive `,ignore,optional`
	UESecurityCapabilities                      *UESecurityCapabilities                      `,reject,mandatory`
	SecurityContext                             *SecurityContext                             `,reject,mandatory`
	NewSecurityContextInd                       *NewSecurityContextInd                       `,reject,optional`
	NASC                                        *NASPDU                                      `,reject,optional`
	PDUSessionResourceSetupListHOReq            *PDUSessionResourceSetupListHOReq            `,reject,mandatory`
	AllowedNSSAI                                *AllowedNSSAI                                `,reject,mandatory`
	TraceActivation                             *TraceActivation                             `,ignore,optional`
	MaskedIMEISV                                *MaskedIMEISV                                `,ignore,optional`
	SourceToTargetTransparentContainer          *SourceToTargetTransparentContainer          `,reject,mandatory`
	MobilityRestrictionList                     *MobilityRestrictionList                     `,ignore,optional`
	LocationReportingRequestType                *LocationReportingRequestType                `,ignore,optional`
	RRCInactiveTransitionReportRequest          *RRCInactiveTransitionReportRequest          `,ignore,optional`
	GUAMI                                       *GUAMI                                       `,reject,mandatory`
	RedirectionVoiceFallback                    *RedirectionVoiceFallback                    `,ignore,optional`
	CNAssistedRANTuning                         *CNAssistedRANTuning                         `,ignore,optional`
}

func (msg *HandoverRequest) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_HandoverResourceAllocation, Criticality_PresentReject, msg.toIes())
}
func (msg *HandoverRequest) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.AMFUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.AMFUENGAPID})
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
	if msg.UEAggregateMaximumBitRate != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UEAggregateMaximumBitRate},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.UEAggregateMaximumBitRate})
	}
	if msg.CoreNetworkAssistanceInformationForInactive != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CoreNetworkAssistanceInformationForInactive},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CoreNetworkAssistanceInformationForInactive})
	}
	if msg.UESecurityCapabilities != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UESecurityCapabilities},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.UESecurityCapabilities})
	}
	if msg.SecurityContext != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SecurityContext},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.SecurityContext})
	}
	if msg.NewSecurityContextInd != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NewSecurityContextInd},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.NewSecurityContextInd})
	}
	if msg.NASC != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NASC},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.NASC})
	}
	if msg.PDUSessionResourceSetupListHOReq != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceSetupListHOReq},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.PDUSessionResourceSetupListHOReq})
	}
	if msg.AllowedNSSAI != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AllowedNSSAI},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.AllowedNSSAI})
	}
	if msg.TraceActivation != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TraceActivation},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.TraceActivation})
	}
	if msg.MaskedIMEISV != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_MaskedIMEISV},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.MaskedIMEISV})
	}
	if msg.SourceToTargetTransparentContainer != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SourceToTargetTransparentContainer},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.SourceToTargetTransparentContainer})
	}
	if msg.MobilityRestrictionList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_MobilityRestrictionList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.MobilityRestrictionList})
	}
	if msg.LocationReportingRequestType != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_LocationReportingRequestType},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.LocationReportingRequestType})
	}
	if msg.RRCInactiveTransitionReportRequest != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RRCInactiveTransitionReportRequest},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RRCInactiveTransitionReportRequest})
	}
	if msg.GUAMI != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_GUAMI},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.GUAMI})
	}
	if msg.RedirectionVoiceFallback != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RedirectionVoiceFallback},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RedirectionVoiceFallback})
	}
	if msg.CNAssistedRANTuning != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CNAssistedRANTuning},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CNAssistedRANTuning})
	}
	return
}
func (msg *HandoverRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *HandoverRequest) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_UEAggregateMaximumBitRate:
		var tmp UEAggregateMaximumBitRate
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UEAggregateMaximumBitRate = &tmp
	case ProtocolIEID_CoreNetworkAssistanceInformationForInactive:
		var tmp CoreNetworkAssistanceInformationForInactive
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.CoreNetworkAssistanceInformationForInactive = &tmp
	case ProtocolIEID_UESecurityCapabilities:
		var tmp UESecurityCapabilities
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UESecurityCapabilities = &tmp
	case ProtocolIEID_SecurityContext:
		var tmp SecurityContext
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SecurityContext = &tmp
	case ProtocolIEID_NewSecurityContextInd:
		var tmp NewSecurityContextInd
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NewSecurityContextInd = &tmp
	case ProtocolIEID_NASC:
		var tmp NASPDU
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NASC = &tmp
	case ProtocolIEID_PDUSessionResourceSetupListHOReq:
		var tmp PDUSessionResourceSetupListHOReq
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceSetupListHOReq = &tmp
	case ProtocolIEID_AllowedNSSAI:
		var tmp AllowedNSSAI
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AllowedNSSAI = &tmp
	case ProtocolIEID_TraceActivation:
		var tmp TraceActivation
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.TraceActivation = &tmp
	case ProtocolIEID_MaskedIMEISV:
		var tmp MaskedIMEISV
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.MaskedIMEISV = &tmp
	case ProtocolIEID_SourceToTargetTransparentContainer:
		var tmp SourceToTargetTransparentContainer
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SourceToTargetTransparentContainer = &tmp
	case ProtocolIEID_MobilityRestrictionList:
		var tmp MobilityRestrictionList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.MobilityRestrictionList = &tmp
	case ProtocolIEID_LocationReportingRequestType:
		var tmp LocationReportingRequestType
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.LocationReportingRequestType = &tmp
	case ProtocolIEID_RRCInactiveTransitionReportRequest:
		var tmp RRCInactiveTransitionReportRequest
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RRCInactiveTransitionReportRequest = &tmp
	case ProtocolIEID_GUAMI:
		var tmp GUAMI
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.GUAMI = &tmp
	case ProtocolIEID_RedirectionVoiceFallback:
		var tmp RedirectionVoiceFallback
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RedirectionVoiceFallback = &tmp
	case ProtocolIEID_CNAssistedRANTuning:
		var tmp CNAssistedRANTuning
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.CNAssistedRANTuning = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
