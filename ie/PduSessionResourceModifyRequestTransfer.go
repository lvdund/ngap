package ie

type PduSessionResourceModifyRequestTransfer struct {
	PduSessionAggregateMaximumBitRate        *PduSessionAggregateMaximumBitRate
	UlNgUUpTnlModifyList                     *[]UlNgUUpTnlModifyItem
	NetworkInstance                          *NetworkInstance
	QosFlowAddOrModifyRequestList            *[]QosFlowAddOrModifyRequestItem
	QosFlowToReleaseList                     *QosFlowListWithCause
	AdditionalUlNgUUpTnlInformation          *UpTransportLayerInformationList
	CommonNetworkInstance                    *CommonNetworkInstance
	AdditionalRedundantUlNgUUpTnlInformation *UpTransportLayerInformationList
	RedundantCommonNetworkInstance           *CommonNetworkInstance
	RedundantUlNgUUpTnlInformation           *UpTransportLayerInformation
	SecurityIndication                       *SecurityIndication
	MbsSessionSetupOrModifyRequestList       *MbsSessionSetupOrModifyRequestList
	MbsSessionToReleaseList                  *MbsSessionToReleaseList
}

type UlNgUUpTnlModifyItem struct {
	UlNgUUpTnlInformation          *UpTransportLayerInformation
	DlNgUUpTnlInformation          *UpTransportLayerInformation
	RedundantUlNgUUpTnlInformation *UpTransportLayerInformation
	RedundantDlNgUUpTnlInformation *UpTransportLayerInformation
}

type QosFlowAddOrModifyRequestItem struct {
	QosFlowIdentifier         *QosFlowIdentifier
	QosFlowLevelQosParameters *QosFlowLevelQosParameters
	ERabId                    *ERabId
	TscTrafficCharacteristics *TscTrafficCharacteristics
	RedundantQosFlowIndicator *RedundantQosFlowIndicator
}
