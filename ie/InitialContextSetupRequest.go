package ie

type InitialContextSetupRequest struct {
	MessageType                                    *MessageType
	AmfUeNgapId                                    *AmfUeNgapId
	RanUeNgapId                                    *RanUeNgapId
	OldAmf                                         *AmfName
	UeAggregateMaximumBitRate                      *UeAggregateMaximumBitRate
	CoreNetworkAssistanceInformationForRrcInactive *CoreNetworkAssistanceInformationForRrcInactive
	Guami                                          *Guami
	PduSessionResourceSetupRequestList             *[]PduSessionResourceSetupRequestItem
	AllowedNssai                                   *AllowedNssai
	UeSecurityCapabilities                         *UeSecurityCapabilities
	SecurityKey                                    *SecurityKey
	TraceActivation                                *TraceActivation
	MobilityRestrictionList                        *MobilityRestrictionList
	UeRadioCapability                              *UeRadioCapability
	IndexToRatFrequencySelectionPriority           *IndexToRatFrequencySelectionPriority
	MaskedImeisv                                   *MaskedImeisv
	NasPdu                                         *NasPdu
	EmergencyFallbackIndicator                     *EmergencyFallbackIndicator
	RrcInactiveTransitionReportRequest             *RrcInactiveTransitionReportRequest
	UeRadioCapabilityForPaging                     *UeRadioCapabilityForPaging
	RedirectionForVoiceEpsFallback                 *RedirectionForVoiceEpsFallback
	LocationReportingRequestType                   *LocationReportingRequestType
	CnAssistedRanParametersTuning                  *CnAssistedRanParametersTuning
	SrvccOperationPossible                         *SrvccOperationPossible
	IabAuthorized                                  *IabAuthorized
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
	RgLevelWirelineAccessCharacteristics           *[]byte
	ManagementBasedMdtPlmnList                     *MdtPlmnList
	UeRadioCapabilityId                            *UeRadioCapabilityId
	TimeSynchronisationAssistanceInformation       *TimeSynchronisationAssistanceInformation
	QmcConfigurationInformation                    *QmcConfigurationInformation
	TargetNssaiInformation                         *TargetNssaiInformation
	UeSliceMaximumBitRateList                      *UeSliceMaximumBitRateList
	Ie5GProseAuthorized                            *Ie5GProseAuthorized
	Ie5GProseUePc5AggregateMaximumBitRate          *NrUeSidelinkAggregateMaximumBitRate
	Ie5GProsePc5QosParameters                      *Ie5GProsePc5QosParameters
}

type PduSessionResourceSetupRequestItem struct {
	PduSessionId                           *PduSessionId
	PduSessionNasPdu                       *NasPdu
	SNssai                                 *SNssai
	PduSessionResourceSetupRequestTransfer *[]byte
	PduSessionExpectedUeActivityBehaviour  *ExpectedUeActivityBehaviour
}
