package ie

type HandoverCommandTransfer struct {
	DlForwardingUpTnlInformation           UpTransportLayerInformation     `bitstring:"sizeLB:0,sizeUB:150"`
	QosFlowToBeForwardedList               []QosFlowToBeForwardedItem      `bitstring:"sizeLB:0,sizeUB:150"`
	DataForwardingResponseDrbList          DataForwardingResponseDrbList   `bitstring:"sizeLB:0,sizeUB:150"`
	AdditionalDlForwardingUpTnlInformation QosFlowPerTnlInformationList    `bitstring:"sizeLB:0,sizeUB:150"`
	UlForwardingUpTnlInformation           UpTransportLayerInformation     `bitstring:"sizeLB:0,sizeUB:150"`
	AdditionalUlForwardingUpTnlInformation UpTransportLayerInformationList `bitstring:"sizeLB:0,sizeUB:150"`
	DataForwardingResponseERabList         DataForwardingResponseERabList  `bitstring:"sizeLB:0,sizeUB:150"`
	QosFlowFailedToSetupList               QosFlowListWithCause            `bitstring:"sizeLB:0,sizeUB:150"`
}

type QosFlowToBeForwardedItem struct {
	QosFlowIdentifier QosFlowIdentifier `bitstring:"sizeLB:0,sizeUB:150"`
}
