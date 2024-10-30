package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type InitialContextSetupRequest struct {
	AMFUENGAPID                                 *AMFUENGAPID                                 `,reject,mandatory`
	RANUENGAPID                                 *RANUENGAPID                                 `,reject,mandatory`
	OldAMF                                      *AMFName                                     `,reject,optional`
	UEAggregateMaximumBitRate                   *UEAggregateMaximumBitRate                   `,reject,conditional`
	CoreNetworkAssistanceInformationForInactive *CoreNetworkAssistanceInformationForInactive `,ignore,optional`
	GUAMI                                       *GUAMI                                       `,reject,mandatory`
	PDUSessionResourceSetupListCxtReq           *PDUSessionResourceSetupListCxtReq           `,reject,optional`
	AllowedNSSAI                                *AllowedNSSAI                                `,reject,mandatory`
	UESecurityCapabilities                      *UESecurityCapabilities                      `,reject,mandatory`
	SecurityKey                                 *SecurityKey                                 `,reject,mandatory`
	TraceActivation                             *TraceActivation                             `,ignore,optional`
	MobilityRestrictionList                     *MobilityRestrictionList                     `,ignore,optional`
	UERadioCapability                           *UERadioCapability                           `,ignore,optional`
	IndexToRFSP                                 *IndexToRFSP                                 `,ignore,optional`
	MaskedIMEISV                                *MaskedIMEISV                                `,ignore,optional`
	NASPDU                                      *NASPDU                                      `,ignore,optional`
	EmergencyFallbackIndicator                  *EmergencyFallbackIndicator                  `,reject,optional`
	RRCInactiveTransitionReportRequest          *RRCInactiveTransitionReportRequest          `,ignore,optional`
	UERadioCapabilityForPaging                  *UERadioCapabilityForPaging                  `,ignore,optional`
	RedirectionVoiceFallback                    *RedirectionVoiceFallback                    `,ignore,optional`
	LocationReportingRequestType                *LocationReportingRequestType                `,ignore,optional`
	CNAssistedRANTuning                         *CNAssistedRANTuning                         `,ignore,optional`
}

func (msg *InitialContextSetupRequest) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_InitialContextSetup, Criticality_PresentReject, msg.toIes())
}
func (msg *InitialContextSetupRequest) toIes() (ies []NgapMessageIE) {
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
	if msg.OldAMF != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_OldAMF},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.OldAMF})
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
	if msg.GUAMI != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_GUAMI},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.GUAMI})
	}
	if msg.PDUSessionResourceSetupListCxtReq != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceSetupListCxtReq},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.PDUSessionResourceSetupListCxtReq})
	}
	if msg.AllowedNSSAI != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AllowedNSSAI},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.AllowedNSSAI})
	}
	if msg.UESecurityCapabilities != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UESecurityCapabilities},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.UESecurityCapabilities})
	}
	if msg.SecurityKey != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SecurityKey},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.SecurityKey})
	}
	if msg.TraceActivation != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TraceActivation},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.TraceActivation})
	}
	if msg.MobilityRestrictionList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_MobilityRestrictionList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.MobilityRestrictionList})
	}
	if msg.UERadioCapability != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UERadioCapability},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UERadioCapability})
	}
	if msg.IndexToRFSP != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_IndexToRFSP},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.IndexToRFSP})
	}
	if msg.MaskedIMEISV != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_MaskedIMEISV},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.MaskedIMEISV})
	}
	if msg.NASPDU != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NASPDU},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.NASPDU})
	}
	if msg.EmergencyFallbackIndicator != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_EmergencyFallbackIndicator},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.EmergencyFallbackIndicator})
	}
	if msg.RRCInactiveTransitionReportRequest != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RRCInactiveTransitionReportRequest},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RRCInactiveTransitionReportRequest})
	}
	if msg.UERadioCapabilityForPaging != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UERadioCapabilityForPaging},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UERadioCapabilityForPaging})
	}
	if msg.RedirectionVoiceFallback != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RedirectionVoiceFallback},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RedirectionVoiceFallback})
	}
	if msg.LocationReportingRequestType != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_LocationReportingRequestType},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.LocationReportingRequestType})
	}
	if msg.CNAssistedRANTuning != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CNAssistedRANTuning},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CNAssistedRANTuning})
	}
	return
}
func (msg *InitialContextSetupRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *InitialContextSetupRequest) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_OldAMF:
		var tmp AMFName
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.OldAMF = &tmp
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
	case ProtocolIEID_GUAMI:
		var tmp GUAMI
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.GUAMI = &tmp
	case ProtocolIEID_PDUSessionResourceSetupListCxtReq:
		var tmp PDUSessionResourceSetupListCxtReq
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceSetupListCxtReq = &tmp
	case ProtocolIEID_AllowedNSSAI:
		var tmp AllowedNSSAI
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AllowedNSSAI = &tmp
	case ProtocolIEID_UESecurityCapabilities:
		var tmp UESecurityCapabilities
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UESecurityCapabilities = &tmp
	case ProtocolIEID_SecurityKey:
		var tmp SecurityKey
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SecurityKey = &tmp
	case ProtocolIEID_TraceActivation:
		var tmp TraceActivation
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.TraceActivation = &tmp
	case ProtocolIEID_MobilityRestrictionList:
		var tmp MobilityRestrictionList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.MobilityRestrictionList = &tmp
	case ProtocolIEID_UERadioCapability:
		var tmp UERadioCapability
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UERadioCapability = &tmp
	case ProtocolIEID_IndexToRFSP:
		var tmp IndexToRFSP
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.IndexToRFSP = &tmp
	case ProtocolIEID_MaskedIMEISV:
		var tmp MaskedIMEISV
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.MaskedIMEISV = &tmp
	case ProtocolIEID_NASPDU:
		var tmp NASPDU
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NASPDU = &tmp
	case ProtocolIEID_EmergencyFallbackIndicator:
		var tmp EmergencyFallbackIndicator
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.EmergencyFallbackIndicator = &tmp
	case ProtocolIEID_RRCInactiveTransitionReportRequest:
		var tmp RRCInactiveTransitionReportRequest
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RRCInactiveTransitionReportRequest = &tmp
	case ProtocolIEID_UERadioCapabilityForPaging:
		var tmp UERadioCapabilityForPaging
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UERadioCapabilityForPaging = &tmp
	case ProtocolIEID_RedirectionVoiceFallback:
		var tmp RedirectionVoiceFallback
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RedirectionVoiceFallback = &tmp
	case ProtocolIEID_LocationReportingRequestType:
		var tmp LocationReportingRequestType
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.LocationReportingRequestType = &tmp
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
