package ie

type PduSessionResourceModifyConfirmTransfer struct {
QosFlowModifyConfirmList	[]QosFlowModifyConfirmItem	//`bitstring:"sizeLB:0,sizeUB:150"`
UlNgUUpTnlInformation	UpTransportLayerInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
AdditionalNgUUpTnlInformation	UpTransportLayerInformationPairList	//`bitstring:"sizeLB:0,sizeUB:150"`
QosFlowFailedToModifyList	QosFlowListWithCause	//`bitstring:"sizeLB:0,sizeUB:150"`
RedundantUlNgUUpTnlInformation	UpTransportLayerInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
AdditionalRedundantNgUUpTnlInformation	UpTransportLayerInformationPairList	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type QosFlowModifyConfirmItem struct {
QosFlowIdentifier	QosFlowIdentifier	//`bitstring:"sizeLB:0,sizeUB:150"`
}
