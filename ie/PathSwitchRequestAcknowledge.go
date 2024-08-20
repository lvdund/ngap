package ie

type PathSwitchRequestAcknowledge struct {
	MessageType                                    *MessageType
	AmfUeNgapId                                    *AmfUeNgapId
	RanUeNgapId                                    *RanUeNgapId
	UeSecurityCapabilities                         *UeSecurityCapabilities
	SecurityContext                                *SecurityContext
	NewSecurityContextIndicator                    *NewSecurityContextIndicator
	PduSessionResourceSwitchedList                 *[]PduSessionResourceSwitchedItem
	PduSessionResourceReleasedList                 *[]PduSessionResourceReleasedItem
	AllowedNssai                                   *AllowedNssai
	CoreNetworkAssistanceInformationForRrcInactive *CoreNetworkAssistanceInformationForRrcInactive
	RrcInactiveTransitionReportRequest             *RrcInactiveTransitionReportRequest
	CriticalityDiagnostics                         *CriticalityDiagnostics
	RedirectionForVoiceEpsFallback                 *RedirectionForVoiceEpsFallback
	CnAssistedRanParametersTuning                  *CnAssistedRanParametersTuning
	SrvccOperationPossible                         *SrvccOperationPossible
	EnhancedCoverageRestriction                    *EnhancedCoverageRestriction
	ExtendedConnectedTime                          *ExtendedConnectedTime
	UeDifferentiationInformation                   *UeDifferentiationInformation
	NrV2XServicesAuthorized                        *NrV2XServicesAuthorized
	LteV2XServicesAuthorized                       *LteV2XServicesAuthorized
	NrUeSidelinkAggregateMaximumBitRate            *NrUeSidelinkAggregateMaximumBitRate
	LteUeSidelinkAggregateMaximumBitRate           *LteUeSidelinkAggregateMaximumBitRate
	Pc5QosParameters                               *Pc5QosParameters
	CeModeBRestricted                              *CeModeBRestricted
	UeUserPlaneCiotSupportIndicator                *UeUserPlaneCiotSupportIndicator
	UeRadioCapabilityId                            *UeRadioCapabilityId
	ManagementBasedMdtPlmnList                     *MdtPlmnList
	TimeSynchronisationAssistanceInformation       *TimeSynchronisationAssistanceInformation
	Ie5GProseAuthorized                            *Ie5GProseAuthorized
	Ie5GProseUePc5AggregateMaximumBitRate          *NrUeSidelinkAggregateMaximumBitRate
	Ie5GProsePc5QosParameters                      *Ie5GProsePc5QosParameters
	ManagementBasedMdtPlmnModificationList         *MdtPlmnModificationList
	IabAuthorized                                  *IabAuthorized
}

type PduSessionResourceSwitchedItem struct {
	PduSessionId                          *PduSessionId
	PathSwitchRequestAcknowledgeTransfer  *[]byte
	PduSessionExpectedUeActivityBehaviour *ExpectedUeActivityBehaviour
}
