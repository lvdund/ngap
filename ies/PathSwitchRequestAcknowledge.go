package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type PathSwitchRequestAcknowledge struct {
	AMFUENGAPID                                 int64
	RANUENGAPID                                 int64
	UESecurityCapabilities                      *UESecurityCapabilities
	SecurityContext                             SecurityContext
	NewSecurityContextInd                       *NewSecurityContextInd
	PDUSessionResourceSwitchedList              []PDUSessionResourceSwitchedItem
	PDUSessionResourceReleasedListPSAck         *[]PDUSessionResourceReleasedItemPSAck
	AllowedNSSAI                                []AllowedNSSAIItem
	CoreNetworkAssistanceInformationForInactive *CoreNetworkAssistanceInformationForInactive
	RRCInactiveTransitionReportRequest          *RRCInactiveTransitionReportRequest
	CriticalityDiagnostics                      *CriticalityDiagnostics
	RedirectionVoiceFallback                    *RedirectionVoiceFallback
	CNAssistedRANTuning                         *CNAssistedRANTuning
}

func (msg *PathSwitchRequestAcknowledge) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_PathSwitchRequest, Criticality_PresentReject, msg.toIes())
}
func (msg *PathSwitchRequestAcknowledge) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 1099511627775},
			ext:   false,
			Value: aper.Integer(msg.AMFUENGAPID),
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4294967295},
			ext:   false,
			Value: aper.Integer(msg.RANUENGAPID),
		}})
	if msg.UESecurityCapabilities != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UESecurityCapabilities},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.UESecurityCapabilities,
		})
	}
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
	tmp_PDUSessionResourceSwitchedList := Sequence[*PDUSessionResourceSwitchedItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
		ext: false,
	}
	for _, i := range msg.PDUSessionResourceSwitchedList {
		tmp_PDUSessionResourceSwitchedList.Value = append(tmp_PDUSessionResourceSwitchedList.Value, &i)
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceSwitchedList},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &tmp_PDUSessionResourceSwitchedList,
	})
	if msg.PDUSessionResourceReleasedListPSAck != nil {
		tmp_PDUSessionResourceReleasedListPSAck := Sequence[*PDUSessionResourceReleasedItemPSAck]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range *msg.PDUSessionResourceReleasedListPSAck {
			tmp_PDUSessionResourceReleasedListPSAck.Value = append(tmp_PDUSessionResourceReleasedListPSAck.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceReleasedListPSAck},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_PDUSessionResourceReleasedListPSAck,
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
	if msg.CoreNetworkAssistanceInformationForInactive != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CoreNetworkAssistanceInformationForInactive},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CoreNetworkAssistanceInformationForInactive,
		})
	}
	if msg.RRCInactiveTransitionReportRequest != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RRCInactiveTransitionReportRequest},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RRCInactiveTransitionReportRequest,
		})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
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
		msg.SecurityContext = tmp
	case ProtocolIEID_NewSecurityContextInd:
		var tmp NewSecurityContextInd
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NewSecurityContextInd = &tmp
	case ProtocolIEID_PDUSessionResourceSwitchedList:
		tmp := Sequence[*PDUSessionResourceSwitchedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceSwitchedList = []PDUSessionResourceSwitchedItem{}
		for _, i := range tmp.Value {
			msg.PDUSessionResourceSwitchedList = append(msg.PDUSessionResourceSwitchedList, *i)
		}
	case ProtocolIEID_PDUSessionResourceReleasedListPSAck:
		tmp := Sequence[*PDUSessionResourceReleasedItemPSAck]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceReleasedListPSAck = &[]PDUSessionResourceReleasedItemPSAck{}
		for _, i := range tmp.Value {
			*msg.PDUSessionResourceReleasedListPSAck = append(*msg.PDUSessionResourceReleasedListPSAck, *i)
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
