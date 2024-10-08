package ie

type HandoverRequestAcknowledgeTransfer struct {
	DlNgUUpTnlInformation                  UpTransportLayerInformation             `bitstring:"sizeLB:0,sizeUB:150"`
	DlForwardingUpTnlInformation           UpTransportLayerInformation             `bitstring:"sizeLB:0,sizeUB:150"`
	SecurityResult                         SecurityResult                          `bitstring:"sizeLB:0,sizeUB:150"`
	QosFlowSetupResponseList               QosFlowListWithDataForwarding           `bitstring:"sizeLB:0,sizeUB:150"`
	QosFlowFailedToSetupList               QosFlowListWithCause                    `bitstring:"sizeLB:0,sizeUB:150"`
	DataForwardingResponseDrbList          DataForwardingResponseDrbList           `bitstring:"sizeLB:0,sizeUB:150"`
	AdditionalDlUpTnlInformationForHoList  []AdditionalDlUpTnlInformationForHoItem `bitstring:"sizeLB:0,sizeUB:150"`
	UlForwardingUpTnlInformation           UpTransportLayerInformation             `bitstring:"sizeLB:0,sizeUB:150"`
	AdditionalUlForwardingUpTnlInformation UpTransportLayerInformationList         `bitstring:"sizeLB:0,sizeUB:150"`
	DataForwardingResponseERabList         DataForwardingResponseERabList          `bitstring:"sizeLB:0,sizeUB:150"`
	RedundantDlNgUUpTnlInformation         UpTransportLayerInformation             `bitstring:"sizeLB:0,sizeUB:150"`
	UsedRsnInformation                     RedundantPduSessionInformation          `bitstring:"sizeLB:0,sizeUB:150"`
	GlobalRanNodeIdOfSecondaryNgRanNode    GlobalRanNodeId                         `bitstring:"sizeLB:0,sizeUB:150"`
	MbsSupportIndicator                    MbsSupportIndicator                     `bitstring:"sizeLB:0,sizeUB:150"`
}

type AdditionalDlUpTnlInformationForHoItem struct {
	AdditionalDlNgUUpTnlInformation          UpTransportLayerInformation   `bitstring:"sizeLB:0,sizeUB:150"`
	AdditionalQosFlowSetupResponseList       QosFlowListWithDataForwarding `bitstring:"sizeLB:0,sizeUB:150"`
	AdditionalDlForwardingUpTnlInformation   UpTransportLayerInformation   `bitstring:"sizeLB:0,sizeUB:150"`
	AdditionalRedundantDlNgUUpTnlInformation UpTransportLayerInformation   `bitstring:"sizeLB:0,sizeUB:150"`
}
