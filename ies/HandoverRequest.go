package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type HandoverRequest struct {
	AMFUENGAPID                                 int64                                        `lb:0,ub:1099511627775,mandatory,reject`
	HandoverType                                HandoverType                                 `mandatory,reject`
	Cause                                       Cause                                        `mandatory,ignore`
	UEAggregateMaximumBitRate                   UEAggregateMaximumBitRate                    `mandatory,reject`
	CoreNetworkAssistanceInformationForInactive *CoreNetworkAssistanceInformationForInactive `optional,ignore`
	UESecurityCapabilities                      UESecurityCapabilities                       `mandatory,reject`
	SecurityContext                             SecurityContext                              `mandatory,reject`
	NewSecurityContextInd                       *NewSecurityContextInd                       `optional,reject`
	NASC                                        []byte                                       `lb:0,ub:0,optional,reject`
	PDUSessionResourceSetupListHOReq            []PDUSessionResourceSetupItemHOReq           `lb:1,ub:maxnoofPDUSessions,mandatory,reject`
	AllowedNSSAI                                []AllowedNSSAIItem                           `lb:1,ub:maxnoofAllowedSNSSAIs,mandatory,reject`
	TraceActivation                             *TraceActivation                             `optional,ignore`
	MaskedIMEISV                                []byte                                       `lb:64,ub:64,optional,ignore`
	SourceToTargetTransparentContainer          []byte                                       `lb:0,ub:0,mandatory,reject`
	MobilityRestrictionList                     *MobilityRestrictionList                     `optional,ignore`
	LocationReportingRequestType                *LocationReportingRequestType                `optional,ignore`
	RRCInactiveTransitionReportRequest          *RRCInactiveTransitionReportRequest          `optional,ignore`
	GUAMI                                       GUAMI                                        `mandatory,reject`
	RedirectionVoiceFallback                    *RedirectionVoiceFallback                    `optional,ignore`
	CNAssistedRANTuning                         *CNAssistedRANTuning                         `optional,ignore`
}

func (msg *HandoverRequest) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		return
	}
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_HandoverResourceAllocation, Criticality_PresentReject, ies)
}
func (msg *HandoverRequest) toIes() (ies []NgapMessageIE, err error) {
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
				Value: msg.NASC,
			}})
	}
	if len(msg.PDUSessionResourceSetupListHOReq) > 0 {
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
	} else {
		err = utils.WrapError("PDUSessionResourceSetupListHOReq is nil", err)
		return
	}
	if len(msg.AllowedNSSAI) > 0 {
		tmp_AllowedNSSAI := Sequence[*AllowedNSSAIItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofAllowedSNSSAIs},
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
	} else {
		err = utils.WrapError("AllowedNSSAI is nil", err)
		return
	}
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
				c:   aper.Constraint{Lb: 64, Ub: 64},
				ext: false,
				Value: aper.BitString{
					Bytes: msg.MaskedIMEISV},
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
func (msg *HandoverRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := HandoverRequestDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*NgapMessageIE),
	}
	if _, err = aper.ReadSequenceOf[NgapMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	if _, ok := decoder.list[ProtocolIEID_AMFUENGAPID]; !ok {
		err = fmt.Errorf("Mandatory field AMFUENGAPID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_HandoverType]; !ok {
		err = fmt.Errorf("Mandatory field HandoverType is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_HandoverType},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_Cause]; !ok {
		err = fmt.Errorf("Mandatory field Cause is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_Cause},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_UEAggregateMaximumBitRate]; !ok {
		err = fmt.Errorf("Mandatory field UEAggregateMaximumBitRate is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_UEAggregateMaximumBitRate},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_UESecurityCapabilities]; !ok {
		err = fmt.Errorf("Mandatory field UESecurityCapabilities is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_UESecurityCapabilities},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_SecurityContext]; !ok {
		err = fmt.Errorf("Mandatory field SecurityContext is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_SecurityContext},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_PDUSessionResourceSetupListHOReq]; !ok {
		err = fmt.Errorf("Mandatory field PDUSessionResourceSetupListHOReq is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceSetupListHOReq},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_AllowedNSSAI]; !ok {
		err = fmt.Errorf("Mandatory field AllowedNSSAI is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_AllowedNSSAI},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_SourceToTargetTransparentContainer]; !ok {
		err = fmt.Errorf("Mandatory field SourceToTargetTransparentContainer is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_SourceToTargetTransparentContainer},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_GUAMI]; !ok {
		err = fmt.Errorf("Mandatory field GUAMI is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_GUAMI},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type HandoverRequestDecoder struct {
	msg      *HandoverRequest
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *HandoverRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	ieId := msgIe.Id.Value
	if _, ok := decoder.list[ieId]; ok {
		err = fmt.Errorf("Duplicated protocol IEID[%d] found", ieId)
		return
	}
	decoder.list[ieId] = msgIe
	ieR := aper.NewReader(bytes.NewReader(buf))
	msg := decoder.msg
	switch msgIe.Id.Value {
	case ProtocolIEID_AMFUENGAPID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 1099511627775},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read AMFUENGAPID", err)
			return
		}
		msg.AMFUENGAPID = int64(tmp.Value)
	case ProtocolIEID_HandoverType:
		var tmp HandoverType
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read HandoverType", err)
			return
		}
		msg.HandoverType = tmp
	case ProtocolIEID_Cause:
		var tmp Cause
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read Cause", err)
			return
		}
		msg.Cause = tmp
	case ProtocolIEID_UEAggregateMaximumBitRate:
		var tmp UEAggregateMaximumBitRate
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read UEAggregateMaximumBitRate", err)
			return
		}
		msg.UEAggregateMaximumBitRate = tmp
	case ProtocolIEID_CoreNetworkAssistanceInformationForInactive:
		var tmp CoreNetworkAssistanceInformationForInactive
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read CoreNetworkAssistanceInformationForInactive", err)
			return
		}
		msg.CoreNetworkAssistanceInformationForInactive = &tmp
	case ProtocolIEID_UESecurityCapabilities:
		var tmp UESecurityCapabilities
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read UESecurityCapabilities", err)
			return
		}
		msg.UESecurityCapabilities = tmp
	case ProtocolIEID_SecurityContext:
		var tmp SecurityContext
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read SecurityContext", err)
			return
		}
		msg.SecurityContext = tmp
	case ProtocolIEID_NewSecurityContextInd:
		var tmp NewSecurityContextInd
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read NewSecurityContextInd", err)
			return
		}
		msg.NewSecurityContextInd = &tmp
	case ProtocolIEID_NASC:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read NASC", err)
			return
		}
		msg.NASC = tmp.Value
	case ProtocolIEID_PDUSessionResourceSetupListHOReq:
		tmp := Sequence[*PDUSessionResourceSetupItemHOReq]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		fn := func() *PDUSessionResourceSetupItemHOReq { return new(PDUSessionResourceSetupItemHOReq) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read PDUSessionResourceSetupListHOReq", err)
			return
		}
		msg.PDUSessionResourceSetupListHOReq = []PDUSessionResourceSetupItemHOReq{}
		for _, i := range tmp.Value {
			msg.PDUSessionResourceSetupListHOReq = append(msg.PDUSessionResourceSetupListHOReq, *i)
		}
	case ProtocolIEID_AllowedNSSAI:
		tmp := Sequence[*AllowedNSSAIItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofAllowedSNSSAIs},
			ext: false,
		}
		fn := func() *AllowedNSSAIItem { return new(AllowedNSSAIItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read AllowedNSSAI", err)
			return
		}
		msg.AllowedNSSAI = []AllowedNSSAIItem{}
		for _, i := range tmp.Value {
			msg.AllowedNSSAI = append(msg.AllowedNSSAI, *i)
		}
	case ProtocolIEID_TraceActivation:
		var tmp TraceActivation
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read TraceActivation", err)
			return
		}
		msg.TraceActivation = &tmp
	case ProtocolIEID_MaskedIMEISV:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 64, Ub: 64},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read MaskedIMEISV", err)
			return
		}
		msg.MaskedIMEISV = tmp.Value.Bytes
	case ProtocolIEID_SourceToTargetTransparentContainer:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read SourceToTargetTransparentContainer", err)
			return
		}
		msg.SourceToTargetTransparentContainer = tmp.Value
	case ProtocolIEID_MobilityRestrictionList:
		var tmp MobilityRestrictionList
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read MobilityRestrictionList", err)
			return
		}
		msg.MobilityRestrictionList = &tmp
	case ProtocolIEID_LocationReportingRequestType:
		var tmp LocationReportingRequestType
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read LocationReportingRequestType", err)
			return
		}
		msg.LocationReportingRequestType = &tmp
	case ProtocolIEID_RRCInactiveTransitionReportRequest:
		var tmp RRCInactiveTransitionReportRequest
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RRCInactiveTransitionReportRequest", err)
			return
		}
		msg.RRCInactiveTransitionReportRequest = &tmp
	case ProtocolIEID_GUAMI:
		var tmp GUAMI
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read GUAMI", err)
			return
		}
		msg.GUAMI = tmp
	case ProtocolIEID_RedirectionVoiceFallback:
		var tmp RedirectionVoiceFallback
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RedirectionVoiceFallback", err)
			return
		}
		msg.RedirectionVoiceFallback = &tmp
	case ProtocolIEID_CNAssistedRANTuning:
		var tmp CNAssistedRANTuning
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read CNAssistedRANTuning", err)
			return
		}
		msg.CNAssistedRANTuning = &tmp
	default:
		switch msgIe.Criticality.Value {
		case Criticality_PresentReject:
			fmt.Errorf("Not comprehended IE ID 0x%04x (criticality: reject)", msgIe.Id.Value)
		case Criticality_PresentIgnore:
			fmt.Errorf("Not comprehended IE ID 0x%04x (criticality: ignore)", msgIe.Id.Value)
		case Criticality_PresentNotify:
			fmt.Errorf("Not comprehended IE ID 0x%04x (criticality: notify)", msgIe.Id.Value)
		}
		if msgIe.Criticality.Value != Criticality_PresentIgnore {
			decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
				IECriticality: msgIe.Criticality,
				IEID:          msgIe.Id,
				TypeOfError:   TypeOfError{Value: TypeOfErrorNotunderstood},
			})
		}
	}
	return
}
