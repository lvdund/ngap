package ie

type PDUSessionResourceModifyRequestTransferIEs struct {
	PDUSessionAggregateMaximumBitRate PDUSessionAggregateMaximumBitRate `,reject,optional`
	ULNGUUPTNLModifyList              ULNGUUPTNLModifyList              `,reject,optional`
	NetworkInstance                   NetworkInstance                   `,reject,optional`
	QosFlowAddOrModifyRequestList     QosFlowAddOrModifyRequestList     `,reject,optional`
	QosFlowToReleaseList              QosFlowListWithCause              `,reject,optional`
	AdditionalULNGUUPTNLInformation   UPTransportLayerInformationList   `,reject,optional`
	CommonNetworkInstance             CommonNetworkInstance             `,ignore,optional`
}
