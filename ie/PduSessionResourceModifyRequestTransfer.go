package ie

type PduSessionResourceModifyRequestTransfer struct {
PduSessionAggregateMaximumBitRate	PduSessionAggregateMaximumBitRate	//`bitstring:"sizeLB:0,sizeUB:150"`
UlNgUUpTnlModifyList	[]UlNgUUpTnlModifyItem	//`bitstring:"sizeLB:0,sizeUB:150"`
NetworkInstance	NetworkInstance	//`bitstring:"sizeLB:0,sizeUB:150"`
QosFlowAddOrModifyRequestList	[]QosFlowAddOrModifyRequestItem	//`bitstring:"sizeLB:0,sizeUB:150"`
QosFlowToReleaseList	QosFlowListWithCause	//`bitstring:"sizeLB:0,sizeUB:150"`
AdditionalUlNgUUpTnlInformation	UpTransportLayerInformationList	//`bitstring:"sizeLB:0,sizeUB:150"`
CommonNetworkInstance	CommonNetworkInstance	//`bitstring:"sizeLB:0,sizeUB:150"`
AdditionalRedundantUlNgUUpTnlInformation	UpTransportLayerInformationList	//`bitstring:"sizeLB:0,sizeUB:150"`
RedundantCommonNetworkInstance	CommonNetworkInstance	//`bitstring:"sizeLB:0,sizeUB:150"`
RedundantUlNgUUpTnlInformation	UpTransportLayerInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
SecurityIndication	SecurityIndication	//`bitstring:"sizeLB:0,sizeUB:150"`
MbsSessionSetupOrModifyRequestList	MbsSessionSetupOrModifyRequestList	//`bitstring:"sizeLB:0,sizeUB:150"`
MbsSessionToReleaseList	MbsSessionToReleaseList	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type UlNgUUpTnlModifyItem struct {
UlNgUUpTnlInformation	UpTransportLayerInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
DlNgUUpTnlInformation	UpTransportLayerInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
RedundantUlNgUUpTnlInformation	UpTransportLayerInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
RedundantDlNgUUpTnlInformation	UpTransportLayerInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type QosFlowAddOrModifyRequestItem struct {
QosFlowIdentifier	QosFlowIdentifier	//`bitstring:"sizeLB:0,sizeUB:150"`
QosFlowLevelQosParameters	QosFlowLevelQosParameters	//`bitstring:"sizeLB:0,sizeUB:150"`
ERabId	ERabId	//`bitstring:"sizeLB:0,sizeUB:150"`
TscTrafficCharacteristics	TscTrafficCharacteristics	//`bitstring:"sizeLB:0,sizeUB:150"`
RedundantQosFlowIndicator	RedundantQosFlowIndicator	//`bitstring:"sizeLB:0,sizeUB:150"`
}
