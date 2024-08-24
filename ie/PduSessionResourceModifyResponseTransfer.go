package ie

type PduSessionResourceModifyResponseTransfer struct {
DlNgUUpTnlInformation	UpTransportLayerInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
UlNgUUpTnlInformation	UpTransportLayerInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
QosFlowAddOrModifyResponseList	[]QosFlowAddOrModifyResponseItem	//`bitstring:"sizeLB:0,sizeUB:150"`
AdditionalDlQosFlowPerTnlInformation	QosFlowPerTnlInformationList	//`bitstring:"sizeLB:0,sizeUB:150"`
QosFlowFailedToAddOrModifyList	QosFlowListWithCause	//`bitstring:"sizeLB:0,sizeUB:150"`
AdditionalNgUUpTnlInformation	UpTransportLayerInformationPairList	//`bitstring:"sizeLB:0,sizeUB:150"`
RedundantDlNgUUpTnlInformation	UpTransportLayerInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
RedundantUlNgUUpTnlInformation	UpTransportLayerInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
AdditionalRedundantDlQosFlowPerTnlInformation	QosFlowPerTnlInformationList	//`bitstring:"sizeLB:0,sizeUB:150"`
AdditionalRedundantNgUUpTnlInformation	UpTransportLayerInformationPairList	//`bitstring:"sizeLB:0,sizeUB:150"`
SecondaryRatUsageInformation	SecondaryRatUsageInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
MbsSupportIndicator	MbsSupportIndicator	//`bitstring:"sizeLB:0,sizeUB:150"`
MbsSessionSetupOrModifyResponseList	MbsSessionSetupResponseList	//`bitstring:"sizeLB:0,sizeUB:150"`
MbsSessionFailedToSetupOrModifyList	MbsSessionFailedToSetupList	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type QosFlowAddOrModifyResponseItem struct {
QosFlowIdentifier	QosFlowIdentifier	//`bitstring:"sizeLB:0,sizeUB:150"`
CurrentQosParametersSetIndex	AlternativeQosParametersSetIndex	//`bitstring:"sizeLB:0,sizeUB:150"`
}
