package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type InitialContextSetupRequest struct {
	AMFUENGAPID                                 int64
	RANUENGAPID                                 int64
	OldAMF                                      *[]byte
	UEAggregateMaximumBitRate                   UEAggregateMaximumBitRate
	CoreNetworkAssistanceInformationForInactive *CoreNetworkAssistanceInformationForInactive
	GUAMI                                       GUAMI
	PDUSessionResourceSetupListCxtReq           *[]PDUSessionResourceSetupItemCxtReq
	AllowedNSSAI                                []AllowedNSSAIItem
	UESecurityCapabilities                      UESecurityCapabilities
	SecurityKey                                 []byte
	TraceActivation                             *TraceActivation
	MobilityRestrictionList                     *MobilityRestrictionList
	UERadioCapability                           *[]byte
	IndexToRFSP                                 *int64
	MaskedIMEISV                                *[]byte
	NASPDU                                      *[]byte
	EmergencyFallbackIndicator                  *EmergencyFallbackIndicator
	RRCInactiveTransitionReportRequest          *RRCInactiveTransitionReportRequest
	UERadioCapabilityForPaging                  *UERadioCapabilityForPaging
	RedirectionVoiceFallback                    *RedirectionVoiceFallback
	LocationReportingRequestType                *LocationReportingRequestType
	CNAssistedRANTuning                         *CNAssistedRANTuning
}

func (msg *InitialContextSetupRequest) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_InitialContextSetup, Criticality_PresentReject, msg.toIes())
}
func (msg *InitialContextSetupRequest) toIes() (ies []NgapMessageIE) {
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
	if msg.OldAMF != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_OldAMF},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 1, Ub: 150},
				ext:   true,
				Value: *msg.OldAMF,
			}})
	}
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
		Id:          ProtocolIEID{Value: ProtocolIEID_GUAMI},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.GUAMI,
	})
	if msg.PDUSessionResourceSetupListCxtReq != nil {
		tmp_PDUSessionResourceSetupListCxtReq := Sequence[*PDUSessionResourceSetupItemCxtReq]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range *msg.PDUSessionResourceSetupListCxtReq {
			tmp_PDUSessionResourceSetupListCxtReq.Value = append(tmp_PDUSessionResourceSetupListCxtReq.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceSetupListCxtReq},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_PDUSessionResourceSetupListCxtReq,
		})
	}
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
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_UESecurityCapabilities},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.UESecurityCapabilities,
	})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_SecurityKey},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &BITSTRING{
			Value: aper.BitString{
				Bytes:   msg.SecurityKey,
				NumBits: uint64(len(msg.SecurityKey))},
		}})
	if msg.TraceActivation != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TraceActivation},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.TraceActivation,
		})
	}
	if msg.MobilityRestrictionList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_MobilityRestrictionList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.MobilityRestrictionList,
		})
	}
	if msg.UERadioCapability != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UERadioCapability},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 0, Ub: 0},
				ext:   false,
				Value: *msg.UERadioCapability,
			}})
	}
	if msg.IndexToRFSP != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_IndexToRFSP},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 1, Ub: 256},
				ext:   true,
				Value: aper.Integer(*msg.IndexToRFSP),
			}})
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
	if msg.NASPDU != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NASPDU},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 0, Ub: 0},
				ext:   false,
				Value: *msg.NASPDU,
			}})
	}
	if msg.EmergencyFallbackIndicator != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_EmergencyFallbackIndicator},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.EmergencyFallbackIndicator,
		})
	}
	if msg.RRCInactiveTransitionReportRequest != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RRCInactiveTransitionReportRequest},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RRCInactiveTransitionReportRequest,
		})
	}
	if msg.UERadioCapabilityForPaging != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UERadioCapabilityForPaging},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UERadioCapabilityForPaging,
		})
	}
	if msg.RedirectionVoiceFallback != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RedirectionVoiceFallback},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RedirectionVoiceFallback,
		})
	}
	if msg.LocationReportingRequestType != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_LocationReportingRequestType},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.LocationReportingRequestType,
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
	case ProtocolIEID_OldAMF:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 150},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.OldAMF = tmp.Value
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
	case ProtocolIEID_GUAMI:
		var tmp GUAMI
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.GUAMI = tmp
	case ProtocolIEID_PDUSessionResourceSetupListCxtReq:
		tmp := Sequence[*PDUSessionResourceSetupItemCxtReq]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceSetupListCxtReq = &[]PDUSessionResourceSetupItemCxtReq{}
		for _, i := range tmp.Value {
			*msg.PDUSessionResourceSetupListCxtReq = append(*msg.PDUSessionResourceSetupListCxtReq, *i)
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
	case ProtocolIEID_UESecurityCapabilities:
		var tmp UESecurityCapabilities
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UESecurityCapabilities = tmp
	case ProtocolIEID_SecurityKey:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 256, Ub: 256},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SecurityKey = tmp.Value.Bytes
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
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.UERadioCapability = tmp.Value
	case ProtocolIEID_IndexToRFSP:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 256},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.IndexToRFSP = int64(tmp.Value)
	case ProtocolIEID_MaskedIMEISV:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 64, Ub: 64},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.MaskedIMEISV = tmp.Value.Bytes
	case ProtocolIEID_NASPDU:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.NASPDU = tmp.Value
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
