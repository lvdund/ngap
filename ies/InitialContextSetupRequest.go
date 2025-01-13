package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type InitialContextSetupRequest struct {
	AMFUENGAPID                                 int64                                        `lb:0,ub:1099511627775,mandatory,reject`
	RANUENGAPID                                 int64                                        `lb:0,ub:4294967295,mandatory,reject`
	OldAMF                                      []byte                                       `lb:1,ub:150,optional,reject,valueExt`
	UEAggregateMaximumBitRate                   UEAggregateMaximumBitRate                    `mandatory,reject`
	CoreNetworkAssistanceInformationForInactive *CoreNetworkAssistanceInformationForInactive `optional,ignore`
	GUAMI                                       GUAMI                                        `mandatory,reject`
	PDUSessionResourceSetupListCxtReq           []PDUSessionResourceSetupItemCxtReq          `lb:1,ub:maxnoofPDUSessions,optional,reject`
	AllowedNSSAI                                []AllowedNSSAIItem                           `lb:1,ub:maxnoofAllowedSNSSAIs,mandatory,reject`
	UESecurityCapabilities                      UESecurityCapabilities                       `mandatory,reject`
	SecurityKey                                 []byte                                       `lb:256,ub:256,mandatory,reject`
	TraceActivation                             *TraceActivation                             `optional,ignore`
	MobilityRestrictionList                     *MobilityRestrictionList                     `optional,ignore`
	UERadioCapability                           []byte                                       `lb:0,ub:0,optional,ignore`
	IndexToRFSP                                 *int64                                       `lb:1,ub:256,optional,ignore,valueExt`
	MaskedIMEISV                                []byte                                       `lb:64,ub:64,optional,ignore`
	NASPDU                                      []byte                                       `lb:0,ub:0,optional,ignore`
	EmergencyFallbackIndicator                  *EmergencyFallbackIndicator                  `optional,reject`
	RRCInactiveTransitionReportRequest          *RRCInactiveTransitionReportRequest          `optional,ignore`
	UERadioCapabilityForPaging                  *UERadioCapabilityForPaging                  `optional,ignore`
	RedirectionVoiceFallback                    *RedirectionVoiceFallback                    `optional,ignore`
	LocationReportingRequestType                *LocationReportingRequestType                `optional,ignore`
	CNAssistedRANTuning                         *CNAssistedRANTuning                         `optional,ignore`
}

func (msg *InitialContextSetupRequest) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		return
	}
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_InitialContextSetup, Criticality_PresentReject, ies)
}
func (msg *InitialContextSetupRequest) toIes() (ies []NgapMessageIE, err error) {
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
				Value: msg.OldAMF,
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
	if len(msg.PDUSessionResourceSetupListCxtReq) > 0 {
		tmp_PDUSessionResourceSetupListCxtReq := Sequence[*PDUSessionResourceSetupItemCxtReq]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range msg.PDUSessionResourceSetupListCxtReq {
			tmp_PDUSessionResourceSetupListCxtReq.Value = append(tmp_PDUSessionResourceSetupListCxtReq.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceSetupListCxtReq},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_PDUSessionResourceSetupListCxtReq,
		})
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
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_UESecurityCapabilities},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.UESecurityCapabilities,
	})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_SecurityKey},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &BITSTRING{
			c:   aper.Constraint{Lb: 256, Ub: 256},
			ext: false,
			Value: aper.BitString{
				Bytes: msg.SecurityKey},
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
				Value: msg.UERadioCapability,
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
				c:   aper.Constraint{Lb: 64, Ub: 64},
				ext: false,
				Value: aper.BitString{
					Bytes: msg.MaskedIMEISV},
			}})
	}
	if msg.NASPDU != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NASPDU},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 0, Ub: 0},
				ext:   false,
				Value: msg.NASPDU,
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
func (msg *InitialContextSetupRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := InitialContextSetupRequestDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_RANUENGAPID]; !ok {
		err = fmt.Errorf("Mandatory field RANUENGAPID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
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
	if _, ok := decoder.list[ProtocolIEID_GUAMI]; !ok {
		err = fmt.Errorf("Mandatory field GUAMI is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_GUAMI},
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
	if _, ok := decoder.list[ProtocolIEID_UESecurityCapabilities]; !ok {
		err = fmt.Errorf("Mandatory field UESecurityCapabilities is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_UESecurityCapabilities},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_SecurityKey]; !ok {
		err = fmt.Errorf("Mandatory field SecurityKey is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_SecurityKey},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type InitialContextSetupRequestDecoder struct {
	msg      *InitialContextSetupRequest
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *InitialContextSetupRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_RANUENGAPID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4294967295},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RANUENGAPID", err)
			return
		}
		msg.RANUENGAPID = int64(tmp.Value)
	case ProtocolIEID_OldAMF:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 150},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read OldAMF", err)
			return
		}
		msg.OldAMF = tmp.Value
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
	case ProtocolIEID_GUAMI:
		var tmp GUAMI
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read GUAMI", err)
			return
		}
		msg.GUAMI = tmp
	case ProtocolIEID_PDUSessionResourceSetupListCxtReq:
		tmp := Sequence[*PDUSessionResourceSetupItemCxtReq]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		fn := func() *PDUSessionResourceSetupItemCxtReq { return new(PDUSessionResourceSetupItemCxtReq) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read PDUSessionResourceSetupListCxtReq", err)
			return
		}
		msg.PDUSessionResourceSetupListCxtReq = []PDUSessionResourceSetupItemCxtReq{}
		for _, i := range tmp.Value {
			msg.PDUSessionResourceSetupListCxtReq = append(msg.PDUSessionResourceSetupListCxtReq, *i)
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
	case ProtocolIEID_UESecurityCapabilities:
		var tmp UESecurityCapabilities
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read UESecurityCapabilities", err)
			return
		}
		msg.UESecurityCapabilities = tmp
	case ProtocolIEID_SecurityKey:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 256, Ub: 256},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read SecurityKey", err)
			return
		}
		msg.SecurityKey = tmp.Value.Bytes
	case ProtocolIEID_TraceActivation:
		var tmp TraceActivation
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read TraceActivation", err)
			return
		}
		msg.TraceActivation = &tmp
	case ProtocolIEID_MobilityRestrictionList:
		var tmp MobilityRestrictionList
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read MobilityRestrictionList", err)
			return
		}
		msg.MobilityRestrictionList = &tmp
	case ProtocolIEID_UERadioCapability:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read UERadioCapability", err)
			return
		}
		msg.UERadioCapability = tmp.Value
	case ProtocolIEID_IndexToRFSP:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 256},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read IndexToRFSP", err)
			return
		}
		*msg.IndexToRFSP = int64(tmp.Value)
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
	case ProtocolIEID_NASPDU:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read NASPDU", err)
			return
		}
		msg.NASPDU = tmp.Value
	case ProtocolIEID_EmergencyFallbackIndicator:
		var tmp EmergencyFallbackIndicator
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read EmergencyFallbackIndicator", err)
			return
		}
		msg.EmergencyFallbackIndicator = &tmp
	case ProtocolIEID_RRCInactiveTransitionReportRequest:
		var tmp RRCInactiveTransitionReportRequest
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RRCInactiveTransitionReportRequest", err)
			return
		}
		msg.RRCInactiveTransitionReportRequest = &tmp
	case ProtocolIEID_UERadioCapabilityForPaging:
		var tmp UERadioCapabilityForPaging
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read UERadioCapabilityForPaging", err)
			return
		}
		msg.UERadioCapabilityForPaging = &tmp
	case ProtocolIEID_RedirectionVoiceFallback:
		var tmp RedirectionVoiceFallback
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RedirectionVoiceFallback", err)
			return
		}
		msg.RedirectionVoiceFallback = &tmp
	case ProtocolIEID_LocationReportingRequestType:
		var tmp LocationReportingRequestType
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read LocationReportingRequestType", err)
			return
		}
		msg.LocationReportingRequestType = &tmp
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
