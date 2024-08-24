package ie

import "ngap/aper"

type UeContextModificationRequest struct {
MessageType	MessageType	//`bitstring:"sizeLB:0,sizeUB:150"`
AmfUeNgapId	AmfUeNgapId	//`bitstring:"sizeLB:0,sizeUB:150"`
RanUeNgapId	RanUeNgapId	//`bitstring:"sizeLB:0,sizeUB:150"`
RanPagingPriority	RanPagingPriority	//`bitstring:"sizeLB:0,sizeUB:150"`
SecurityKey	SecurityKey	//`bitstring:"sizeLB:0,sizeUB:150"`
IndexToRatFrequencySelectionPriority	IndexToRatFrequencySelectionPriority	//`bitstring:"sizeLB:0,sizeUB:150"`
UeAggregateMaximumBitRate	UeAggregateMaximumBitRate	//`bitstring:"sizeLB:0,sizeUB:150"`
UeSecurityCapabilities	UeSecurityCapabilities	//`bitstring:"sizeLB:0,sizeUB:150"`
CoreNetworkAssistanceInformationForRrcInactive	CoreNetworkAssistanceInformationForRrcInactive	//`bitstring:"sizeLB:0,sizeUB:150"`
EmergencyFallbackIndicator	EmergencyFallbackIndicator	//`bitstring:"sizeLB:0,sizeUB:150"`
NewAmfUeNgapId	AmfUeNgapId	//`bitstring:"sizeLB:0,sizeUB:150"`
RrcInactiveTransitionReportRequest	RrcInactiveTransitionReportRequest	//`bitstring:"sizeLB:0,sizeUB:150"`
NewGuami	Guami	//`bitstring:"sizeLB:0,sizeUB:150"`
CnAssistedRanParametersTuning	CnAssistedRanParametersTuning	//`bitstring:"sizeLB:0,sizeUB:150"`
SrvccOperationPossible	SrvccOperationPossible	//`bitstring:"sizeLB:0,sizeUB:150"`
IabAuthorized	IabAuthorized	//`bitstring:"sizeLB:0,sizeUB:150"`
NrV2XServicesAuthorized	NrV2XServicesAuthorized	//`bitstring:"sizeLB:0,sizeUB:150"`
LteV2XServicesAuthorized	LteV2XServicesAuthorized	//`bitstring:"sizeLB:0,sizeUB:150"`
NrUeSidelinkAggregateMaximumBitRate	NrUeSidelinkAggregateMaximumBitRate	//`bitstring:"sizeLB:0,sizeUB:150"`
LteUeSidelinkAggregateMaximumBitRate	LteUeSidelinkAggregateMaximumBitRate	//`bitstring:"sizeLB:0,sizeUB:150"`
Pc5QosParameters	Pc5QosParameters	//`bitstring:"sizeLB:0,sizeUB:150"`
UeRadioCapabilityId	UeRadioCapabilityId	//`bitstring:"sizeLB:0,sizeUB:150"`
RgLevelWirelineAccessCharacteristics	aper.OctetString	//`octetstring:"sizeLB:0,sizeUB:150"`
TimeSynchronisationAssistanceInformation	TimeSynchronisationAssistanceInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
QmcConfigurationInformation	QmcConfigurationInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
QmcDeactivation	QmcDeactivation	//`bitstring:"sizeLB:0,sizeUB:150"`
UeSliceMaximumBitRateList	UeSliceMaximumBitRateList	//`bitstring:"sizeLB:0,sizeUB:150"`
ManagementBasedMdtPlmnModificationList	MdtPlmnModificationList	//`bitstring:"sizeLB:0,sizeUB:150"`
Ie5GProseAuthorized	Ie5GProseAuthorized	//`bitstring:"sizeLB:0,sizeUB:150"`
Ie5GProseUePc5AggregateMaximumBitRate	NrUeSidelinkAggregateMaximumBitRate	//`bitstring:"sizeLB:0,sizeUB:150"`
Ie5GProsePc5QosParameters	Ie5GProsePc5QosParameters	//`bitstring:"sizeLB:0,sizeUB:150"`
}
