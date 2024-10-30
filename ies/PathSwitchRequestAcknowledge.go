package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type PathSwitchRequestAcknowledge struct {
	AMFUENGAPID                                 *AMFUENGAPID                                 `,ignore,mandatory`
	RANUENGAPID                                 *RANUENGAPID                                 `,ignore,mandatory`
	UESecurityCapabilities                      *UESecurityCapabilities                      `,reject,optional`
	SecurityContext                             *SecurityContext                             `,reject,mandatory`
	NewSecurityContextInd                       *NewSecurityContextInd                       `,reject,optional`
	PDUSessionResourceSwitchedList              *PDUSessionResourceSwitchedList              `,ignore,mandatory`
	PDUSessionResourceReleasedListPSAck         *PDUSessionResourceReleasedListPSAck         `,ignore,optional`
	AllowedNSSAI                                *AllowedNSSAI                                `,reject,mandatory`
	CoreNetworkAssistanceInformationForInactive *CoreNetworkAssistanceInformationForInactive `,ignore,optional`
	RRCInactiveTransitionReportRequest          *RRCInactiveTransitionReportRequest          `,ignore,optional`
	CriticalityDiagnostics                      *CriticalityDiagnostics                      `,ignore,optional`
	RedirectionVoiceFallback                    *RedirectionVoiceFallback                    `,ignore,optional`
	CNAssistedRANTuning                         *CNAssistedRANTuning                         `,ignore,optional`
}

func (msg *PathSwitchRequestAcknowledge) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_PathSwitchRequest, Criticality_PresentReject, msg.toIes())
}
func (msg *PathSwitchRequestAcknowledge) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.AMFUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.AMFUENGAPID})
	}
	if msg.RANUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RANUENGAPID})
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
	if msg.PDUSessionResourceSwitchedList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceSwitchedList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.PDUSessionResourceSwitchedList})
	}
	if msg.PDUSessionResourceReleasedListPSAck != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceReleasedListPSAck},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.PDUSessionResourceReleasedListPSAck})
	}
	if msg.AllowedNSSAI != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AllowedNSSAI},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.AllowedNSSAI})
	}
	if msg.CoreNetworkAssistanceInformationForInactive != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CoreNetworkAssistanceInformationForInactive},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CoreNetworkAssistanceInformationForInactive})
	}
	if msg.RRCInactiveTransitionReportRequest != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RRCInactiveTransitionReportRequest},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RRCInactiveTransitionReportRequest})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics})
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
func (msg *PathSwitchRequestAcknowledge) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *PathSwitchRequestAcknowledge) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_PDUSessionResourceSwitchedList:
		var tmp PDUSessionResourceSwitchedList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceSwitchedList = &tmp
	case ProtocolIEID_PDUSessionResourceReleasedListPSAck:
		var tmp PDUSessionResourceReleasedListPSAck
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceReleasedListPSAck = &tmp
	case ProtocolIEID_AllowedNSSAI:
		var tmp AllowedNSSAI
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AllowedNSSAI = &tmp
	case ProtocolIEID_CoreNetworkAssistanceInformationForInactive:
		var tmp CoreNetworkAssistanceInformationForInactive
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.CoreNetworkAssistanceInformationForInactive = &tmp
	case ProtocolIEID_RRCInactiveTransitionReportRequest:
		var tmp RRCInactiveTransitionReportRequest
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RRCInactiveTransitionReportRequest = &tmp
	case ProtocolIEID_CriticalityDiagnostics:
		var tmp CriticalityDiagnostics
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.CriticalityDiagnostics = &tmp
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
