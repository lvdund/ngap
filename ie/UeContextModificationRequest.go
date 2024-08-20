package ie

type UeContextModificationRequest struct {
	MessageType                                    *MessageType
	AmfUeNgapId                                    *AmfUeNgapId
	RanUeNgapId                                    *RanUeNgapId
	RanPagingPriority                              *RanPagingPriority
	SecurityKey                                    *SecurityKey
	IndexToRatFrequencySelectionPriority           *IndexToRatFrequencySelectionPriority
	UeAggregateMaximumBitRate                      *UeAggregateMaximumBitRate
	UeSecurityCapabilities                         *UeSecurityCapabilities
	CoreNetworkAssistanceInformationForRrcInactive *CoreNetworkAssistanceInformationForRrcInactive
	EmergencyFallbackIndicator                     *EmergencyFallbackIndicator
	NewAmfUeNgapId                                 *AmfUeNgapId
	RrcInactiveTransitionReportRequest             *RrcInactiveTransitionReportRequest
	NewGuami                                       *Guami
	CnAssistedRanParametersTuning                  *CnAssistedRanParametersTuning
	SrvccOperationPossible                         *SrvccOperationPossible
	IabAuthorized                                  *IabAuthorized
	NrV2XServicesAuthorized                        *NrV2XServicesAuthorized
	LteV2XServicesAuthorized                       *LteV2XServicesAuthorized
	NrUeSidelinkAggregateMaximumBitRate            *NrUeSidelinkAggregateMaximumBitRate
	LteUeSidelinkAggregateMaximumBitRate           *LteUeSidelinkAggregateMaximumBitRate
	Pc5QosParameters                               *Pc5QosParameters
	UeRadioCapabilityId                            *UeRadioCapabilityId
	RgLevelWirelineAccessCharacteristics           *[]byte
	TimeSynchronisationAssistanceInformation       *TimeSynchronisationAssistanceInformation
	QmcConfigurationInformation                    *QmcConfigurationInformation
	QmcDeactivation                                *QmcDeactivation
	UeSliceMaximumBitRateList                      *UeSliceMaximumBitRateList
	ManagementBasedMdtPlmnModificationList         *MdtPlmnModificationList
	Ie5GProseAuthorized                            *Ie5GProseAuthorized
	Ie5GProseUePc5AggregateMaximumBitRate          *NrUeSidelinkAggregateMaximumBitRate
	Ie5GProsePc5QosParameters                      *Ie5GProsePc5QosParameters
}
