package ie

import "ngap/aper"

type PathSwitchRequestAcknowledge struct {
	MessageType                                    MessageType                                    //`bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                                    AmfUeNgapId                                    //`bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                                    RanUeNgapId                                    //`bitstring:"sizeLB:0,sizeUB:150"`
	UeSecurityCapabilities                         UeSecurityCapabilities                         //`bitstring:"sizeLB:0,sizeUB:150"`
	SecurityContext                                SecurityContext                                //`bitstring:"sizeLB:0,sizeUB:150"`
	NewSecurityContextIndicator                    NewSecurityContextIndicator                    //`bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceSwitchedList                 []PduSessionResourceSwitchedItem               //`bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceReleasedList                 []PduSessionResourceReleasedItem               //`bitstring:"sizeLB:0,sizeUB:150"`
	AllowedNssai                                   AllowedNssai                                   //`bitstring:"sizeLB:0,sizeUB:150"`
	CoreNetworkAssistanceInformationForRrcInactive CoreNetworkAssistanceInformationForRrcInactive //`bitstring:"sizeLB:0,sizeUB:150"`
	RrcInactiveTransitionReportRequest             RrcInactiveTransitionReportRequest             //`bitstring:"sizeLB:0,sizeUB:150"`
	CriticalityDiagnostics                         CriticalityDiagnostics                         //`bitstring:"sizeLB:0,sizeUB:150"`
	RedirectionForVoiceEpsFallback                 RedirectionForVoiceEpsFallback                 //`bitstring:"sizeLB:0,sizeUB:150"`
	CnAssistedRanParametersTuning                  CnAssistedRanParametersTuning                  //`bitstring:"sizeLB:0,sizeUB:150"`
	SrvccOperationPossible                         SrvccOperationPossible                         //`bitstring:"sizeLB:0,sizeUB:150"`
	EnhancedCoverageRestriction                    EnhancedCoverageRestriction                    //`bitstring:"sizeLB:0,sizeUB:150"`
	ExtendedConnectedTime                          ExtendedConnectedTime                          //`bitstring:"sizeLB:0,sizeUB:150"`
	UeDifferentiationInformation                   UeDifferentiationInformation                   //`bitstring:"sizeLB:0,sizeUB:150"`
	NrV2XServicesAuthorized                        NrV2XServicesAuthorized                        //`bitstring:"sizeLB:0,sizeUB:150"`
	LteV2XServicesAuthorized                       LteV2XServicesAuthorized                       //`bitstring:"sizeLB:0,sizeUB:150"`
	NrUeSidelinkAggregateMaximumBitRate            NrUeSidelinkAggregateMaximumBitRate            //`bitstring:"sizeLB:0,sizeUB:150"`
	LteUeSidelinkAggregateMaximumBitRate           LteUeSidelinkAggregateMaximumBitRate           //`bitstring:"sizeLB:0,sizeUB:150"`
	Pc5QosParameters                               Pc5QosParameters                               //`bitstring:"sizeLB:0,sizeUB:150"`
	CeModeBRestricted                              CeModeBRestricted                              //`bitstring:"sizeLB:0,sizeUB:150"`
	UeUserPlaneCiotSupportIndicator                UeUserPlaneCiotSupportIndicator                //`bitstring:"sizeLB:0,sizeUB:150"`
	UeRadioCapabilityId                            UeRadioCapabilityId                            //`bitstring:"sizeLB:0,sizeUB:150"`
	ManagementBasedMdtPlmnList                     MdtPlmnList                                    //`bitstring:"sizeLB:0,sizeUB:150"`
	TimeSynchronisationAssistanceInformation       TimeSynchronisationAssistanceInformation       //`bitstring:"sizeLB:0,sizeUB:150"`
	Ie5GProseAuthorized                            Ie5GProseAuthorized                            //`bitstring:"sizeLB:0,sizeUB:150"`
	Ie5GProseUePc5AggregateMaximumBitRate          NrUeSidelinkAggregateMaximumBitRate            //`bitstring:"sizeLB:0,sizeUB:150"`
	Ie5GProsePc5QosParameters                      Ie5GProsePc5QosParameters                      //`bitstring:"sizeLB:0,sizeUB:150"`
	ManagementBasedMdtPlmnModificationList         MdtPlmnModificationList                        //`bitstring:"sizeLB:0,sizeUB:150"`
	IabAuthorized                                  IabAuthorized                                  //`bitstring:"sizeLB:0,sizeUB:150"`
}

type PduSessionResourceSwitchedItem struct {
	PduSessionId                          PduSessionId                //`bitstring:"sizeLB:0,sizeUB:150"`
	PathSwitchRequestAcknowledgeTransfer  aper.OctetString            //`octetstring:"sizeLB:0,sizeUB:150"`
	PduSessionExpectedUeActivityBehaviour ExpectedUeActivityBehaviour //`bitstring:"sizeLB:0,sizeUB:150"`
}
