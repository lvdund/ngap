package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

func encodeTransferMessage(w io.Writer, ies []NgapMessageIE) (err error) {
	aw := aper.NewWriter(w)
	if err = aw.WriteBool(aper.Zero); err != nil {
		return
	}
	if len(ies) == 0 {
		err = fmt.Errorf("empty message")
		return
	}

	if err = aper.WriteSequenceOf[NgapMessageIE](ies, aw, &aper.Constraint{
		Lb: 0,
		Ub: int64(aper.POW_16 - 1),
	}, false); err != nil {
		return
	}

	err = aw.Close()
	return
}

func encodeMessage(w io.Writer, present uint8, procedureCode int64, criticality aper.Enumerated, ies []NgapMessageIE) (err error) {
	aw := aper.NewWriter(w)
	if err = aw.WriteBool(aper.Zero); err != nil {
		return
	}
	if err = aw.WriteChoice(uint64(present), 2, true); err != nil {
		return
	}
	pCode := ProcedureCode{
		Value: aper.Integer(procedureCode),
	}
	if err = pCode.Encode(aw); err != nil {
		return
	}
	cr := Criticality{
		Value: criticality,
	}
	if err = cr.Encode(aw); err != nil {
		return
	}
	if len(ies) == 0 {
		err = fmt.Errorf("empty message")
		return
	}

	var buf bytes.Buffer
	cW := aper.NewWriter(&buf) //container writer
	cW.WriteBool(aper.Zero)
	if err = aper.WriteSequenceOf[NgapMessageIE](ies, cW, &aper.Constraint{
		Lb: 0,
		Ub: int64(aper.POW_16 - 1),
	}, false); err != nil {
		return
	}

	if err = cW.Close(); err != nil { //finalize
		return
	}
	if err = aw.WriteOpenType(buf.Bytes()); err != nil {
		return
	}
	err = aw.Close()
	return
}

// represent an IE in Ngap messages
type NgapMessageIE struct {
	Id          ProtocolIEID //protocol IE identity
	Criticality Criticality
	Value       aper.AperMarshaller //open type
}

func (ie NgapMessageIE) Encode(w *aper.AperWriter) (err error) {
	//1. encode protocol Ie Id
	if err = ie.Id.Encode(w); err != nil {
		return
	}
	//2. encode criticality
	if err = ie.Criticality.Encode(w); err != nil {
		return
	}
	//3. encode NgapIE
	//encode IE into a byte array first
	var buf bytes.Buffer
	ieW := aper.NewWriter(&buf)
	if err = ie.Value.Encode(ieW); err != nil {
		return
	}
	ieW.Close()
	//then write the array as open type
	err = w.WriteOpenType(buf.Bytes())
	return
}

type ProcedureCode struct {
	Value aper.Integer `aper:"valueLB:0,valueUB:255"`
}

func (ie *ProcedureCode) Decode(r *aper.AperReader) error {
	if v, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return err
	} else {
		ie.Value = aper.Integer(v)
	}
	return nil
}
func (ie *ProcedureCode) Encode(r *aper.AperWriter) (err error) {
	if err = r.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return err
	}
	return nil
}

type TriggeringMessage struct {
	Value aper.Enumerated `aper:"valueLB:0,valueUB:2"`
}

func (ie *TriggeringMessage) Decode(r *aper.AperReader) error {
	if v, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return err
	} else {
		ie.Value = aper.Enumerated(v)
	}
	return nil
}
func (ie *TriggeringMessage) Encode(r *aper.AperWriter) (err error) {
	if err = r.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return err
	}
	return nil
}

type Criticality struct {
	Value aper.Enumerated `aper:"valueLB:0,valueUB:2"`
}

func (ie *Criticality) Decode(r *aper.AperReader) error {
	if v, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return err
	} else {
		ie.Value = aper.Enumerated(v)
	}
	return nil
}
func (ie *Criticality) Encode(r *aper.AperWriter) (err error) {
	if err = r.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return err
	}
	return nil
}

type ProtocolIEID struct {
	Value aper.Integer `aper:"valueLB:0,valueUB:65535"`
}

func (ie *ProtocolIEID) Decode(r *aper.AperReader) error {
	if v, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return err
	} else {
		ie.Value = aper.Integer(v)
	}
	return nil
}
func (ie *ProtocolIEID) Encode(r *aper.AperWriter) (err error) {
	if err = r.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return err
	}
	return nil
}

func BuildDiagnostics(present uint8, procedureCode ProcedureCode, criticality Criticality, diagnosticsItems []CriticalityDiagnostics) (diagnostics *CriticalityDiagnostics) {
	//TODO: build diagnostic content
	return
}

// NgapPdu - Present
const (
	NgapPresentNothing uint8 = iota
	NgapPduInitiatingMessage
	NgapPduSuccessfulOutcome
	NgapPduUnsuccessfulOutcome
)

// ProcedureCode
const (
	ProcedureCode_AMFConfigurationUpdate                = 0 // id-AMFConfigurationUpdate
	ProcedureCode_AMFStatusIndication                   = 1
	ProcedureCode_CellTrafficTrace                      = 2
	ProcedureCode_DeactivateTrace                       = 3
	ProcedureCode_DownlinkNASTransport                  = 4
	ProcedureCode_DownlinkNonUEAssociatedNRPPaTransport = 5
	ProcedureCode_DownlinkRANConfigurationTransfer      = 6
	ProcedureCode_DownlinkRANStatusTransfer             = 7
	ProcedureCode_DownlinkUEAssociatedNRPPaTransport    = 8
	ProcedureCode_ErrorIndication                       = 9
	ProcedureCode_HandoverCancel                        = 10
	ProcedureCode_HandoverNotification                  = 11
	ProcedureCode_HandoverPreparation                   = 12
	ProcedureCode_HandoverResourceAllocation            = 13
	ProcedureCode_InitialContextSetup                   = 14
	ProcedureCode_InitialUEMessage                      = 15
	ProcedureCode_LocationReportingControl              = 16
	ProcedureCode_LocationReportingFailureIndication    = 17
	ProcedureCode_LocationReport                        = 18
	ProcedureCode_NASNonDeliveryIndication              = 19
	ProcedureCode_NGReset                               = 20
	ProcedureCode_NGSetup                               = 21
	ProcedureCode_OverloadStart                         = 22
	ProcedureCode_OverloadStop                          = 23
	ProcedureCode_Paging                                = 24
	ProcedureCode_PathSwitchRequest                     = 25
	ProcedureCode_PDUSessionResourceModify              = 26
	ProcedureCode_PDUSessionResourceModifyIndication    = 27
	ProcedureCode_PDUSessionResourceRelease             = 28
	ProcedureCode_PDUSessionResourceSetup               = 29
	ProcedureCode_PDUSessionResourceNotify              = 30
	ProcedureCode_PrivateMessage                        = 31
	ProcedureCode_PWSCancel                             = 32
	ProcedureCode_PWSFailureIndication                  = 33
	ProcedureCode_PWSRestartIndication                  = 34
	ProcedureCode_RANConfigurationUpdate                = 35
	ProcedureCode_RerouteNASRequest                     = 36
	ProcedureCode_RRCInactiveTransitionReport           = 37
	ProcedureCode_TraceFailureIndication                = 38
	ProcedureCode_TraceStart                            = 39
	ProcedureCode_UEContextModification                 = 40
	ProcedureCode_UEContextRelease                      = 41
	ProcedureCode_UEContextReleaseRequest               = 42
	ProcedureCode_UERadioCapabilityCheck                = 43
	ProcedureCode_UERadioCapabilityInfoIndication       = 44
	ProcedureCode_UETNLABindingRelease                  = 45
	ProcedureCode_UplinkNASTransport                    = 46
	ProcedureCode_UplinkNonUEAssociatedNRPPaTransport   = 47
	ProcedureCode_UplinkRANConfigurationTransfer        = 48
	ProcedureCode_UplinkRANStatusTransfer               = 49
	ProcedureCode_UplinkUEAssociatedNRPPaTransport      = 50
	ProcedureCode_WriteReplaceWarning                   = 51
	ProcedureCode_SecondaryRATDataUsageReport           = 52
)

// Criticality
const (
	Criticality_PresentReject aper.Enumerated = 0
	Criticality_PresentIgnore aper.Enumerated = 1
	Criticality_PresentNotify aper.Enumerated = 2
)

// NgapProtocolIeId
const (
	ProtocolIEID_AllowedNSSAI                                = 0 // id-AllowedNSSAI
	ProtocolIEID_AMFName                                     = 1
	ProtocolIEID_AMFOverloadResponse                         = 2
	ProtocolIEID_AMFSetID                                    = 3
	ProtocolIEID_AMFTNLAssociationFailedToSetupList          = 4
	ProtocolIEID_AMFTNLAssociationSetupList                  = 5
	ProtocolIEID_AMFTNLAssociationToAddList                  = 6
	ProtocolIEID_AMFTNLAssociationToRemoveList               = 7
	ProtocolIEID_AMFTNLAssociationToUpdateList               = 8
	ProtocolIEID_AMFTrafficLoadReductionIndication           = 9
	ProtocolIEID_AMFUENGAPID                                 = 10
	ProtocolIEID_AssistanceDataForPaging                     = 11
	ProtocolIEID_BroadcastCancelledAreaList                  = 12
	ProtocolIEID_BroadcastCompletedAreaList                  = 13
	ProtocolIEID_CancelAllWarningMessages                    = 14
	ProtocolIEID_Cause                                       = 15
	ProtocolIEID_CellIDListForRestart                        = 16
	ProtocolIEID_ConcurrentWarningMessageInd                 = 17
	ProtocolIEID_CoreNetworkAssistanceInformationForInactive = 18
	ProtocolIEID_CriticalityDiagnostics                      = 19
	ProtocolIEID_DataCodingScheme                            = 20
	ProtocolIEID_DefaultPagingDRX                            = 21
	ProtocolIEID_DirectForwardingPathAvailability            = 22
	ProtocolIEID_EmergencyAreaIDListForRestart               = 23
	ProtocolIEID_EmergencyFallbackIndicator                  = 24
	ProtocolIEID_EUTRACGI                                    = 25
	ProtocolIEID_FiveGSTMSI                                  = 26
	ProtocolIEID_GlobalRANNodeID                             = 27
	ProtocolIEID_GUAMI                                       = 28
	ProtocolIEID_HandoverType                                = 29
	ProtocolIEID_IMSVoiceSupportIndicator                    = 30
	ProtocolIEID_IndexToRFSP                                 = 31
	ProtocolIEID_InfoOnRecommendedCellsAndRANNodesForPaging  = 32
	ProtocolIEID_LocationReportingRequestType                = 33
	ProtocolIEID_MaskedIMEISV                                = 34
	ProtocolIEID_MessageIdentifier                           = 35
	ProtocolIEID_MobilityRestrictionList                     = 36
	ProtocolIEID_NASC                                        = 37
	ProtocolIEID_NASPDU                                      = 38
	ProtocolIEID_NASSecurityParametersFromNGRAN              = 39
	ProtocolIEID_NewAMFUENGAPID                              = 40
	ProtocolIEID_NewSecurityContextInd                       = 41
	ProtocolIEID_NGAPMessage                                 = 42
	ProtocolIEID_NGRANCGI                                    = 43
	ProtocolIEID_NGRANTraceID                                = 44
	ProtocolIEID_NRCGI                                       = 45
	ProtocolIEID_NRPPaPDU                                    = 46
	ProtocolIEID_NumberOfBroadcastsRequested                 = 47
	ProtocolIEID_OldAMF                                      = 48
	ProtocolIEID_OverloadStartNSSAIList                      = 49
	ProtocolIEID_PagingDRX                                   = 50
	ProtocolIEID_PagingOrigin                                = 51
	ProtocolIEID_PagingPriority                              = 52
	ProtocolIEID_PDUSessionResourceAdmittedList              = 53
	ProtocolIEID_PDUSessionResourceFailedToModifyListModRes  = 54
	ProtocolIEID_PDUSessionResourceFailedToSetupListCxtRes   = 55
	ProtocolIEID_PDUSessionResourceFailedToSetupListHOAck    = 56
	ProtocolIEID_PDUSessionResourceFailedToSetupListPSReq    = 57
	ProtocolIEID_PDUSessionResourceFailedToSetupListSURes    = 58
	ProtocolIEID_PDUSessionResourceHandoverList              = 59
	ProtocolIEID_PDUSessionResourceListCxtRelCpl             = 60
	ProtocolIEID_PDUSessionResourceListHORqd                 = 61
	ProtocolIEID_PDUSessionResourceModifyListModCfm          = 62
	ProtocolIEID_PDUSessionResourceModifyListModInd          = 63
	ProtocolIEID_PDUSessionResourceModifyListModReq          = 64
	ProtocolIEID_PDUSessionResourceModifyListModRes          = 65
	ProtocolIEID_PDUSessionResourceNotifyList                = 66
	ProtocolIEID_PDUSessionResourceReleasedListNot           = 67
	ProtocolIEID_PDUSessionResourceReleasedListPSAck         = 68
	ProtocolIEID_PDUSessionResourceReleasedListPSFail        = 69
	ProtocolIEID_PDUSessionResourceReleasedListRelRes        = 70
	ProtocolIEID_PDUSessionResourceSetupListCxtReq           = 71
	ProtocolIEID_PDUSessionResourceSetupListCxtRes           = 72
	ProtocolIEID_PDUSessionResourceSetupListHOReq            = 73
	ProtocolIEID_PDUSessionResourceSetupListSUReq            = 74
	ProtocolIEID_PDUSessionResourceSetupListSURes            = 75
	ProtocolIEID_PDUSessionResourceToBeSwitchedDLList        = 76
	ProtocolIEID_PDUSessionResourceSwitchedList              = 77
	ProtocolIEID_PDUSessionResourceToReleaseListHOCmd        = 78
	ProtocolIEID_PDUSessionResourceToReleaseListRelCmd       = 79
	ProtocolIEID_PLMNSupportList                             = 80
	ProtocolIEID_PWSFailedCellIDList                         = 81
	ProtocolIEID_RANNodeName                                 = 82
	ProtocolIEID_RANPagingPriority                           = 83
	ProtocolIEID_RANStatusTransferTransparentContainer       = 84
	ProtocolIEID_RANUENGAPID                                 = 85
	ProtocolIEID_RelativeAMFCapacity                         = 86
	ProtocolIEID_RepetitionPeriod                            = 87
	ProtocolIEID_ResetType                                   = 88
	ProtocolIEID_RoutingID                                   = 89
	ProtocolIEID_RRCEstablishmentCause                       = 90
	ProtocolIEID_RRCInactiveTransitionReportRequest          = 91
	ProtocolIEID_RRCState                                    = 92
	ProtocolIEID_SecurityContext                             = 93
	ProtocolIEID_SecurityKey                                 = 94
	ProtocolIEID_SerialNumber                                = 95
	ProtocolIEID_ServedGUAMIList                             = 96
	ProtocolIEID_SliceSupportList                            = 97
	ProtocolIEID_SONConfigurationTransferDL                  = 98
	ProtocolIEID_SONConfigurationTransferUL                  = 99
	ProtocolIEID_SourceAMFUENGAPID                           = 100
	ProtocolIEID_SourceToTargetTransparentContainer          = 101
	ProtocolIEID_SupportedTAList                             = 102
	ProtocolIEID_TAIListForPaging                            = 103
	ProtocolIEID_TAIListForRestart                           = 104
	ProtocolIEID_TargetID                                    = 105
	ProtocolIEID_TargetToSourceTransparentContainer          = 106
	ProtocolIEID_TimeToWait                                  = 107
	ProtocolIEID_TraceActivation                             = 108
	ProtocolIEID_TraceCollectionEntityIPAddress              = 109
	ProtocolIEID_UEAggregateMaximumBitRate                   = 110
	ProtocolIEID_UEassociatedLogicalNGconnectionList         = 111
	ProtocolIEID_UEContextRequest                            = 112
	ProtocolIEID_UENGAPIDs                                   = 114
	ProtocolIEID_UEPagingIdentity                            = 115
	ProtocolIEID_UEPresenceInAreaOfInterestList              = 116
	ProtocolIEID_UERadioCapability                           = 117
	ProtocolIEID_UERadioCapabilityForPaging                  = 118
	ProtocolIEID_UESecurityCapabilities                      = 119
	ProtocolIEID_UnavailableGUAMIList                        = 120
	ProtocolIEID_UserLocationInformation                     = 121
	ProtocolIEID_WarningAreaList                             = 122
	ProtocolIEID_WarningMessageContents                      = 123
	ProtocolIEID_WarningSecurityInfo                         = 124
	ProtocolIEID_WarningType                                 = 125
	ProtocolIEID_AdditionalULNGUUPTNLInformation             = 126
	ProtocolIEID_DataForwardingNotPossible                   = 127
	ProtocolIEID_DLNGUUPTNLInformation                       = 128
	ProtocolIEID_NetworkInstance                             = 129
	ProtocolIEID_PDUSessionAggregateMaximumBitRate           = 130
	ProtocolIEID_PDUSessionResourceFailedToModifyListModCfm  = 131
	ProtocolIEID_PDUSessionResourceFailedToSetupListCxtFail  = 132
	ProtocolIEID_PDUSessionResourceListCxtRelReq             = 133
	ProtocolIEID_PDUSessionType                              = 134
	ProtocolIEID_QosFlowAddOrModifyRequestList               = 135
	ProtocolIEID_QosFlowSetupRequestList                     = 136
	ProtocolIEID_QosFlowToReleaseList                        = 137
	ProtocolIEID_SecurityIndication                          = 138
	ProtocolIEID_ULNGUUPTNLInformation                       = 139
	ProtocolIEID_ULNGUUPTNLModifyList                        = 140
	ProtocolIEID_WarningAreaCoordinates                      = 141
	ProtocolIEID_PDUSessionResourceSecondaryRATUsageList     = 142
	ProtocolIEID_HandoverFlag                                = 143
	ProtocolIEID_SecondaryRATUsageInformation                = 144
	ProtocolIEID_PDUSessionResourceReleaseResponseTransfer   = 145
	ProtocolIEID_RedirectionVoiceFallback                    = 146
	ProtocolIEID_UERetentionInformation                      = 147
	ProtocolIEID_SNSSAI                                      = 148
	ProtocolIEID_PSCellInformation                           = 149
	ProtocolIEID_LastEUTRANPLMNIdentity                      = 150
	ProtocolIEID_MaximumIntegrityProtectedDataRateDL         = 151
	ProtocolIEID_AdditionalDLForwardingUPTNLInformation      = 152
	ProtocolIEID_AdditionalDLUPTNLInformationForHOList       = 153
	ProtocolIEID_AdditionalNGUUPTNLInformation               = 154
	ProtocolIEID_AdditionalDLQosFlowPerTNLInformation        = 155
	ProtocolIEID_SecurityResult                              = 156
	ProtocolIEID_ENDCSONConfigurationTransferDL              = 157
	ProtocolIEID_ENDCSONConfigurationTransferUL              = 158
	ProtocolIEID_OldAssociatedQosFlowListULendmarkerexpected = 159
	ProtocolIEID_CNTypeRestrictionsForEquivalent             = 160
	ProtocolIEID_CNTypeRestrictionsForServing                = 161
	ProtocolIEID_NewGUAMI                                    = 162
	ProtocolIEID_ULForwarding                                = 163
	ProtocolIEID_ULForwardingUPTNLInformation                = 164
	ProtocolIEID_CNAssistedRANTuning                         = 165
	ProtocolIEID_CommonNetworkInstance                       = 166
	ProtocolIEID_NGRANTNLAssociationToRemoveList             = 167
	ProtocolIEID_TNLAssociationTransportLayerAddressNGRAN    = 168
	ProtocolIEID_EndpointIPAddressAndPort                    = 169
	ProtocolIEID_LocationReportingAdditionalInfo             = 170
	ProtocolIEID_SourceToTargetAMFInformationReroute         = 171
	ProtocolIEID_AdditionalULForwardingUPTNLInformation      = 172
	ProtocolIEID_SCTPTLAs                                    = 173
	ProtocolIEID_SelectedPLMNIdentity                        = 174
)
