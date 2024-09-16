package ie

import "ngap/aper"

type NgapProtocolIeId struct {
	NgapProtocolIeId aper.Integer `Integer:"valueLB:0,valueUB:65535"`
}

func (ie *NgapProtocolIeId) Decode(r aper.AperReader) error {

	return nil
}

func (ie *NgapProtocolIeId) Encode(r aper.AperWriter) (err error) {
	if err = r.WriteInteger(int64(ie.NgapProtocolIeId), &aper.Constrain{Lb: 10, Ub: 10}, false); err != nil {
		return err
	}
	return nil
}

const (
	ProtocolIEIDAllowedNSSAI                               aper.Integer = 0
	ProtocolIEIDAMFName                                    aper.Integer = 1
	ProtocolIEIDAMFOverloadResponse                        aper.Integer = 2
	ProtocolIEIDAMFSetID                                   aper.Integer = 3
	ProtocolIEIDAMFTNLAssociationFailedToSetupList         aper.Integer = 4
	ProtocolIEIDAMFTNLAssociationSetupList                 aper.Integer = 5
	ProtocolIEIDAMFTNLAssociationToAddList                 aper.Integer = 6
	ProtocolIEIDAMFTNLAssociationToRemoveList              aper.Integer = 7
	ProtocolIEIDAMFTNLAssociationToUpdateList              aper.Integer = 8
	ProtocolIEIDAMFTrafficLoadReductionIndication          aper.Integer = 9
	ProtocolIEIDAMFUENGAPID                                aper.Integer = 10
	ProtocolIEIDAssistanceDataForPaging                    aper.Integer = 11
	ProtocolIEIDBroadcastCancelledAreaList                 aper.Integer = 12
	ProtocolIEIDBroadcastCompletedAreaList                 aper.Integer = 13
	ProtocolIEIDCancelAllWarningMessages                   aper.Integer = 14
	ProtocolIEIDCause                                      aper.Integer = 15
	ProtocolIEIDCellIDListForRestart                       aper.Integer = 16
	ProtocolIEIDConcurrentWarningMessageInd                aper.Integer = 17
	ProtocolIEIDCoreNetworkAssistanceInformation           aper.Integer = 18
	ProtocolIEIDCriticalityDiagnostics                     aper.Integer = 19
	ProtocolIEIDDataCodingScheme                           aper.Integer = 20
	ProtocolIEIDDefaultPagingDRX                           aper.Integer = 21
	ProtocolIEIDDirectForwardingPathAvailability           aper.Integer = 22
	ProtocolIEIDEmergencyAreaIDListForRestart              aper.Integer = 23
	ProtocolIEIDEmergencyFallbackIndicator                 aper.Integer = 24
	ProtocolIEIDEUTRACGI                                   aper.Integer = 25
	ProtocolIEIDFiveGSTMSI                                 aper.Integer = 26
	ProtocolIEIDGlobalRANNodeID                            aper.Integer = 27
	ProtocolIEIDGUAMI                                      aper.Integer = 28
	ProtocolIEIDHandoverType                               aper.Integer = 29
	ProtocolIEIDIMSVoiceSupportIndicator                   aper.Integer = 30
	ProtocolIEIDIndexToRFSP                                aper.Integer = 31
	ProtocolIEIDInfoOnRecommendedCellsAndRANNodesForPaging aper.Integer = 32
	ProtocolIEIDLocationReportingRequestType               aper.Integer = 33
	ProtocolIEIDMaskedIMEISV                               aper.Integer = 34
	ProtocolIEIDMessageIdentifier                          aper.Integer = 35
	ProtocolIEIDMobilityRestrictionList                    aper.Integer = 36
	ProtocolIEIDNASC                                       aper.Integer = 37
	ProtocolIEIDNASPDU                                     aper.Integer = 38
	ProtocolIEIDNASSecurityParametersFromNGRAN             aper.Integer = 39
	ProtocolIEIDNewAMFUENGAPID                             aper.Integer = 40
	ProtocolIEIDNewSecurityContextInd                      aper.Integer = 41
	ProtocolIEIDNGAPMessage                                aper.Integer = 42
	ProtocolIEIDNGRANCGI                                   aper.Integer = 43
	ProtocolIEIDNGRANTraceID                               aper.Integer = 44
	ProtocolIEIDNRCGI                                      aper.Integer = 45
	ProtocolIEIDNRPPaPDU                                   aper.Integer = 46
	ProtocolIEIDNumberOfBroadcastsRequested                aper.Integer = 47
	ProtocolIEIDOldAMF                                     aper.Integer = 48
	ProtocolIEIDOverloadStartNSSAIList                     aper.Integer = 49
	ProtocolIEIDPagingDRX                                  aper.Integer = 50
	ProtocolIEIDPagingOrigin                               aper.Integer = 51
	ProtocolIEIDPagingPriority                             aper.Integer = 52
	ProtocolIEIDPDUSessionResourceAdmittedList             aper.Integer = 53
	ProtocolIEIDPDUSessionResourceFailedToModifyListModRes aper.Integer = 54
	ProtocolIEIDPDUSessionResourceFailedToSetupListCxtRes  aper.Integer = 55
	ProtocolIEIDPDUSessionResourceFailedToSetupListHOAck   aper.Integer = 56
	ProtocolIEIDPDUSessionResourceFailedToSetupListPSReq   aper.Integer = 57
	ProtocolIEIDPDUSessionResourceFailedToSetupListSURes   aper.Integer = 58
	ProtocolIEIDPDUSessionResourceHandoverList             aper.Integer = 59
	ProtocolIEIDPDUSessionResourceListCxtRelCpl            aper.Integer = 60
	ProtocolIEIDPDUSessionResourceListHORqd                aper.Integer = 61
	ProtocolIEIDPDUSessionResourceModifyListModCfm         aper.Integer = 62
	ProtocolIEIDPDUSessionResourceModifyListModInd         aper.Integer = 63
	ProtocolIEIDPDUSessionResourceModifyListModReq         aper.Integer = 64
	ProtocolIEIDPDUSessionResourceModifyListModRes         aper.Integer = 65
	ProtocolIEIDPDUSessionResourceNotifyList               aper.Integer = 66
	ProtocolIEIDPDUSessionResourceReleasedListNot          aper.Integer = 67
	ProtocolIEIDPDUSessionResourceReleasedListPSAck        aper.Integer = 68
	ProtocolIEIDPDUSessionResourceReleasedListPSFail       aper.Integer = 69
	ProtocolIEIDPDUSessionResourceReleasedListRelRes       aper.Integer = 70
	ProtocolIEIDPDUSessionResourceSetupListCxtReq          aper.Integer = 71
	ProtocolIEIDPDUSessionResourceSetupListCxtRes          aper.Integer = 72
	ProtocolIEIDPDUSessionResourceSetupListHOReq           aper.Integer = 73
	ProtocolIEIDPDUSessionResourceSetupListSUReq           aper.Integer = 74
	ProtocolIEIDPDUSessionResourceSetupListSURes           aper.Integer = 75
	ProtocolIEIDPDUSessionResourceToBeSwitchedDLList       aper.Integer = 76
	ProtocolIEIDPDUSessionResourceSwitchedList             aper.Integer = 77
	ProtocolIEIDPDUSessionResourceToReleaseListHOCmd       aper.Integer = 78
	ProtocolIEIDPDUSessionResourceToReleaseListRelCmd      aper.Integer = 79
	ProtocolIEIDPLMNSupportList                            aper.Integer = 80
	ProtocolIEIDPWSFailedCellIDList                        aper.Integer = 81
	ProtocolIEIDRANNodeName                                aper.Integer = 82
	ProtocolIEIDRANPagingPriority                          aper.Integer = 83
	ProtocolIEIDRANStatusTransferTransparentContainer      aper.Integer = 84
	ProtocolIEIDRANUENGAPID                                aper.Integer = 85
	ProtocolIEIDRelativeAMFCapacity                        aper.Integer = 86
	ProtocolIEIDRepetitionPeriod                           aper.Integer = 87
	ProtocolIEIDResetType                                  aper.Integer = 88
	ProtocolIEIDRoutingID                                  aper.Integer = 89
	ProtocolIEIDRRCEstablishmentCause                      aper.Integer = 90
	ProtocolIEIDRRCInactiveTransitionReportRequest         aper.Integer = 91
	ProtocolIEIDRRCState                                   aper.Integer = 92
	ProtocolIEIDSecurityContext                            aper.Integer = 93
	ProtocolIEIDSecurityKey                                aper.Integer = 94
	ProtocolIEIDSerialNumber                               aper.Integer = 95
	ProtocolIEIDServedGUAMIList                            aper.Integer = 96
	ProtocolIEIDSliceSupportList                           aper.Integer = 97
	ProtocolIEIDSONConfigurationTransferDL                 aper.Integer = 98
	ProtocolIEIDSONConfigurationTransferUL                 aper.Integer = 99
	ProtocolIEIDSourceAMFUENGAPID                          aper.Integer = 100
	ProtocolIEIDSourceToTargetTransparentContainer         aper.Integer = 101
	ProtocolIEIDSupportedTAList                            aper.Integer = 102
	ProtocolIEIDTAIListForPaging                           aper.Integer = 103
	ProtocolIEIDTAIListForRestart                          aper.Integer = 104
	ProtocolIEIDTargetID                                   aper.Integer = 105
	ProtocolIEIDTargetToSourceTransparentContainer         aper.Integer = 106
	ProtocolIEIDTimeToWait                                 aper.Integer = 107
	ProtocolIEIDTraceActivation                            aper.Integer = 108
	ProtocolIEIDTraceCollectionEntityIPAddress             aper.Integer = 109
	ProtocolIEIDUEAggregateMaximumBitRate                  aper.Integer = 110
	ProtocolIEIDUEAssociatedLogicalNGConnectionList        aper.Integer = 111
	ProtocolIEIDUEContextRequest                           aper.Integer = 112
	ProtocolIEIDUENGAPIDs                                  aper.Integer = 114
	ProtocolIEIDUEPagingIdentity                           aper.Integer = 115
	ProtocolIEIDUEPresenceInAreaOfInterestList             aper.Integer = 116
	ProtocolIEIDUERadioCapability                          aper.Integer = 117
	ProtocolIEIDUERadioCapabilityForPaging                 aper.Integer = 118
	ProtocolIEIDUESecurityCapabilities                     aper.Integer = 119
	ProtocolIEIDUnavailableGUAMIList                       aper.Integer = 120
	ProtocolIEIDUserLocationInformation                    aper.Integer = 121
	ProtocolIEIDWarningAreaList                            aper.Integer = 122
	ProtocolIEIDWarningMessageContents                     aper.Integer = 123
	ProtocolIEIDWarningSecurityInfo                        aper.Integer = 124
	ProtocolIEIDWarningType                                aper.Integer = 125
	ProtocolIEIDAdditionalULNGUUPTNLInformation            aper.Integer = 126
	ProtocolIEIDDataForwardingNotPossible                  aper.Integer = 127
	ProtocolIEIDDLNGUUPTNLInformation                      aper.Integer = 128
	ProtocolIEIDNetworkInstance                            aper.Integer = 129
	ProtocolIEIDPDUSessionAggregateMaximumBitRate          aper.Integer = 130
	ProtocolIEIDPDUSessionResourceFailedToModifyListModCfm aper.Integer = 131
	ProtocolIEIDPDUSessionResourceFailedToSetupListCxtFail aper.Integer = 132
	ProtocolIEIDPDUSessionResourceListCxtRelReq            aper.Integer = 133
	ProtocolIEIDPDUSessionType                             aper.Integer = 134
	ProtocolIEIDQosFlowAddOrModifyRequestList              aper.Integer = 135
	ProtocolIEIDQosFlowSetupRequestList                    aper.Integer = 136
	ProtocolIEIDQosFlowToReleaseList                       aper.Integer = 137
	ProtocolIEIDSecurityIndication                         aper.Integer = 138
	ProtocolIEIDULNGUUPTNLInformation                      aper.Integer = 139
	ProtocolIEIDULNGUUPTNLModifyList                       aper.Integer = 140
	ProtocolIEIDWarningAreaCoordinates                     aper.Integer = 141
	ProtocolIEIDPDUSessionResourceSecondaryRATUsageList    aper.Integer = 142
	ProtocolIEIDHandoverFlag                               aper.Integer = 143
	ProtocolIEIDSecondaryRATUsageInformation               aper.Integer = 144
	ProtocolIEIDPDUSessionResourceReleaseResponseTransfer  aper.Integer = 145
	ProtocolIEIDRedirectionVoiceFallback                   aper.Integer = 146
	ProtocolIEIDUERetentionInformation                     aper.Integer = 147
	ProtocolIEIDSNSSAI                                     aper.Integer = 148
	ProtocolIEIDPSCellInformation                          aper.Integer = 149
	ProtocolIEIDLastEUTRANPLMNIdentity                     aper.Integer = 150
	ProtocolIEIDMaximumIntegrityProtectedDataRateDL        aper.Integer = 151
	ProtocolIEIDAdditionalDLForwardingUPTNLInformation     aper.Integer = 152
	ProtocolIEIDAdditionalDLUPTNLInformationForHOList      aper.Integer = 153
	ProtocolIEIDAdditionalNGUUPTNLInformation              aper.Integer = 154
	ProtocolIEIDAdditionalDLQosFlowPerTNLInformation       aper.Integer = 155
	ProtocolIEIDSecurityResult                             aper.Integer = 156
	ProtocolIEIDENDCSONConfigurationTransferDL             aper.Integer = 157
	ProtocolIEIDENDCSONConfigurationTransferUL             aper.Integer = 158
	ProtocolIEIDGlobalTNGFID                               aper.Integer = 240
	ProtocolIEIDGlobalTWIFID                               aper.Integer = 241
	ProtocolIEIDGlobalWAGFID                               aper.Integer = 242
	ProtocolIEIDUserLocationInformationWAGF                aper.Integer = 243
	ProtocolIEIDUserLocationInformationTNGF                aper.Integer = 244
	ProtocolIEIDUserLocationInformationTWIF                aper.Integer = 248
	ProtocolIEIDNPNSupport                                 aper.Integer = 258
	ProtocolIEIDNPNAccessInformation                       aper.Integer = 259
	ProtocolIEIDNPNPagingAssistanceInformation             aper.Integer = 260
	ProtocolIEIDNPNMobilityInformation                     aper.Integer = 261
	ProtocolIEIDNID                                        aper.Integer = 263
	ProtocolIEIDExtendedSliceSupportList                   aper.Integer = 270
	ProtocolIEIDExtendedTAISliceSupportList                aper.Integer = 271
)
