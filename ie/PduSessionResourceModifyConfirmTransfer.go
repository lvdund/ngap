package ie

type PduSessionResourceModifyConfirmTransfer struct {
	QosFlowModifyConfirmList               *[]QosFlowModifyConfirmItem
	UlNgUUpTnlInformation                  *UpTransportLayerInformation
	AdditionalNgUUpTnlInformation          *UpTransportLayerInformationPairList
	QosFlowFailedToModifyList              *QosFlowListWithCause
	RedundantUlNgUUpTnlInformation         *UpTransportLayerInformation
	AdditionalRedundantNgUUpTnlInformation *UpTransportLayerInformationPairList
}

type QosFlowModifyConfirmItem struct {
	QosFlowIdentifier *QosFlowIdentifier
}
