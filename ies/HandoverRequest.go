package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type HandoverRequest struct {
	AMFUENGAPID                                 int64
	HandoverType                                HandoverType
	Cause                                       Cause
	UEAggregateMaximumBitRate                   UEAggregateMaximumBitRate
	CoreNetworkAssistanceInformationForInactive *CoreNetworkAssistanceInformationForInactive
	UESecurityCapabilities                      UESecurityCapabilities
	SecurityContext                             SecurityContext
	NewSecurityContextInd                       *NewSecurityContextInd
	NASC                                        *[]byte
	PDUSessionResourceSetupListHOReq            []PDUSessionResourceSetupItemHOReq
	AllowedNSSAI                                []AllowedNSSAIItem
	TraceActivation                             *TraceActivation
	MaskedIMEISV                                *[]byte
	SourceToTargetTransparentContainer          []byte
	MobilityRestrictionList                     *MobilityRestrictionList
	LocationReportingRequestType                *LocationReportingRequestType
	RRCInactiveTransitionReportRequest          *RRCInactiveTransitionReportRequest
	GUAMI                                       GUAMI
	RedirectionVoiceFallback                    *RedirectionVoiceFallback
	CNAssistedRANTuning                         *CNAssistedRANTuning
}

func (msg *HandoverRequest) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_HandoverResourceAllocation, Criticality_PresentReject, msg.toIes())
}
func (msg *HandoverRequest) toIes() (ies []NgapMessageIE) {
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
		Id:          ProtocolIEID{Value: ProtocolIEID_UEAggregateMaximumBitRate},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.UEAggregateMaximumBitRate,
	})
	if msg.CoreNetworkAssistanceInformationForInactive != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CoreNetworkAssistanceInformationForInactive},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CoreNetworkAssistanceInformationForInactive,
		})
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_UESecurityCapabilities},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.UESecurityCapabilities,
	})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_SecurityContext},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.SecurityContext,
	})
	if msg.NewSecurityContextInd != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NewSecurityContextInd},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.NewSecurityContextInd,
		})
	}
	if msg.NASC != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NASC},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 0, Ub: 0},
				ext:   false,
				Value: *msg.NASC,
			}})
	}
	tmp_PDUSessionResourceSetupListHOReq := Sequence[*PDUSessionResourceSetupItemHOReq]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
		ext: false,
	}
	for _, i := range msg.PDUSessionResourceSetupListHOReq {
		tmp_PDUSessionResourceSetupListHOReq.Value = append(tmp_PDUSessionResourceSetupListHOReq.Value, &i)
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceSetupListHOReq},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &tmp_PDUSessionResourceSetupListHOReq,
	})
	tmp_AllowedNSSAI := Sequence[*AllowedNSSAIItem]{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	for _, i := range msg.AllowedNSSAI {
		tmp_AllowedNSSAI.Value = append(tmp_AllowedNSSAI.Value, &i)
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_AllowedNSSAI},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &tmp_AllowedNSSAI,
	})
	if msg.TraceActivation != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TraceActivation},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.TraceActivation,
		})
	}
	if msg.MaskedIMEISV != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_MaskedIMEISV},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &BITSTRING{
				Value: aper.BitString{
					Bytes:   *msg.MaskedIMEISV,
					NumBits: uint64(len(*msg.MaskedIMEISV))},
			}})
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_SourceToTargetTransparentContainer},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &OCTETSTRING{
			c:     aper.Constraint{Lb: 0, Ub: 0},
			ext:   false,
			Value: msg.SourceToTargetTransparentContainer,
		}})
	if msg.MobilityRestrictionList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_MobilityRestrictionList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.MobilityRestrictionList,
		})
	}
	if msg.LocationReportingRequestType != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_LocationReportingRequestType},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.LocationReportingRequestType,
		})
	}
	if msg.RRCInactiveTransitionReportRequest != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RRCInactiveTransitionReportRequest},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RRCInactiveTransitionReportRequest,
		})
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_GUAMI},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.GUAMI,
	})
	if msg.RedirectionVoiceFallback != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RedirectionVoiceFallback},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RedirectionVoiceFallback,
		})
	}
	if msg.CNAssistedRANTuning != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CNAssistedRANTuning},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CNAssistedRANTuning,
		})
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
	case ProtocolIEID_UEAggregateMaximumBitRate:
		var tmp UEAggregateMaximumBitRate
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UEAggregateMaximumBitRate = tmp
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
		msg.UESecurityCapabilities = tmp
	case ProtocolIEID_SecurityContext:
		var tmp SecurityContext
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SecurityContext = tmp
	case ProtocolIEID_NewSecurityContextInd:
		var tmp NewSecurityContextInd
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NewSecurityContextInd = &tmp
	case ProtocolIEID_NASC:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.NASC = tmp.Value
	case ProtocolIEID_PDUSessionResourceSetupListHOReq:
		tmp := Sequence[*PDUSessionResourceSetupItemHOReq]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceSetupListHOReq = []PDUSessionResourceSetupItemHOReq{}
		for _, i := range tmp.Value {
			msg.PDUSessionResourceSetupListHOReq = append(msg.PDUSessionResourceSetupListHOReq, *i)
		}
	case ProtocolIEID_AllowedNSSAI:
		tmp := Sequence[*AllowedNSSAIItem]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AllowedNSSAI = []AllowedNSSAIItem{}
		for _, i := range tmp.Value {
			msg.AllowedNSSAI = append(msg.AllowedNSSAI, *i)
		}
	case ProtocolIEID_TraceActivation:
		var tmp TraceActivation
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.TraceActivation = &tmp
	case ProtocolIEID_MaskedIMEISV:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 64, Ub: 64},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.MaskedIMEISV = tmp.Value.Bytes
	case ProtocolIEID_SourceToTargetTransparentContainer:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SourceToTargetTransparentContainer = tmp.Value
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
		msg.GUAMI = tmp
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
