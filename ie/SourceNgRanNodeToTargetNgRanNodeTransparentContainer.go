package ie

import "ngap/aper"

type SourceNgRanNodeToTargetNgRanNodeTransparentContainer struct {
	RrcContainer                                  aper.OctetString                                //`octetstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceInformationList             []PduSessionResourceInformationItem             //`bitstring:"sizeLB:0,sizeUB:150"`
	ERabInformationList                           []ERabInformationItem                           //`bitstring:"sizeLB:0,sizeUB:150"`
	TargetCellId                                  NgRanCgi                                        //`bitstring:"sizeLB:0,sizeUB:150"`
	IndexToRatFrequencySelectionPriority          IndexToRatFrequencySelectionPriority            //`bitstring:"sizeLB:0,sizeUB:150"`
	UeHistoryInformation                          UeHistoryInformation                            //`bitstring:"sizeLB:0,sizeUB:150"`
	SgnbUeX2ApId                                  SgnbUeX2ApId                                    //`bitstring:"sizeLB:0,sizeUB:150"`
	UeHistoryInformationFromUe                    UeHistoryInformationFromUe                      //`bitstring:"sizeLB:0,sizeUB:150"`
	SourceNodeId                                  SourceNodeId                                    //`bitstring:"sizeLB:0,sizeUB:150"`
	UeContextReferenceAtSource                    RanUeNgapId                                     //`bitstring:"sizeLB:0,sizeUB:150"`
	MbsActiveSessionInformationSourceToTargetList []MbsActiveSessionInformationSourceToTargetItem //`bitstring:"sizeLB:0,sizeUB:150"`
	QmcConfigurationInformation                   QmcConfigurationInformation                     //`bitstring:"sizeLB:0,sizeUB:150"`
	NgapIeSupportInformationRequestList           []NgapIeSupportInformationRequestItem           //`bitstring:"sizeLB:0,sizeUB:150"`
}

type PduSessionResourceInformationItem struct {
	PduSessionId              PduSessionId              //`bitstring:"sizeLB:0,sizeUB:150"`
	QosFlowInformationList    []QosFlowInformationItem  //`bitstring:"sizeLB:0,sizeUB:150"`
	DrbsToQosFlowsMappingList DrbsToQosFlowsMappingList //`bitstring:"sizeLB:0,sizeUB:150"`
}

type QosFlowInformationItem struct {
	QosFlowIdentifier               QosFlowIdentifier     //`bitstring:"sizeLB:0,sizeUB:150"`
	DlForwarding                    DlForwarding          //`bitstring:"sizeLB:0,sizeUB:150"`
	UlForwarding                    UlForwarding          //`bitstring:"sizeLB:0,sizeUB:150"`
	SourceTransportLayerAddress     TransportLayerAddress //`bitstring:"sizeLB:0,sizeUB:150"`
	SourceNodeTransportLayerAddress TransportLayerAddress //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ERabInformationItem struct {
	ERabId                          ERabId                //`bitstring:"sizeLB:0,sizeUB:150"`
	DlForwarding                    DlForwarding          //`bitstring:"sizeLB:0,sizeUB:150"`
	SourceTransportLayerAddress     TransportLayerAddress //`bitstring:"sizeLB:0,sizeUB:150"`
	SourceNodeTransportLayerAddress TransportLayerAddress //`bitstring:"sizeLB:0,sizeUB:150"`
}

type MbsActiveSessionInformationSourceToTargetItem struct {
	MbsSessionId                           MbsSessionId                             //`bitstring:"sizeLB:0,sizeUB:150"`
	MbsAreaSessionId                       MbsAreaSessionId                         //`bitstring:"sizeLB:0,sizeUB:150"`
	MbsServiceArea                         MbsServiceArea                           //`bitstring:"sizeLB:0,sizeUB:150"`
	MbsQosFlowsToBeSetupList               MbsQosFlowsToBeSetupList                 //`bitstring:"sizeLB:0,sizeUB:150"`
	MbsMappingAndDataForwardingRequestList []MbsMappingAndDataForwardingRequestItem //`bitstring:"sizeLB:0,sizeUB:150"`
}

type MbsMappingAndDataForwardingRequestItem struct {
	MrbId                  MrbId                  //`bitstring:"sizeLB:0,sizeUB:150"`
	MbsQosFlowList         MbsQosFlowList         //`bitstring:"sizeLB:0,sizeUB:150"`
	MrbProgressInformation MrbProgressInformation //`bitstring:"sizeLB:0,sizeUB:150"`
}

type MbsQosFlowList struct {
	MbsQosFlowIdentifier QosFlowIdentifier //`bitstring:"sizeLB:0,sizeUB:150"`
}

type NgapIeSupportInformationRequestItem struct {
	NgapProtocolIeId NgapProtocolIeId //`bitstring:"sizeLB:0,sizeUB:150"`
}
