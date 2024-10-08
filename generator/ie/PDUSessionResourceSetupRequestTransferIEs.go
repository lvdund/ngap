package ie

type PDUSessionResourceSetupRequestTransferIEs struct {
	PDUSessionAggregateMaximumBitRate PDUSessionAggregateMaximumBitRate `,reject,optional`
	ULNGUUPTNLInformation             UPTransportLayerInformation       `,reject,mandatory`
	AdditionalULNGUUPTNLInformation   UPTransportLayerInformationList   `,reject,optional`
	DataForwardingNotPossible         DataForwardingNotPossible         `,reject,optional`
	PDUSessionType                    PDUSessionType                    `,reject,mandatory`
	SecurityIndication                SecurityIndication                `,reject,optional`
	NetworkInstance                   NetworkInstance                   `,reject,optional`
	QosFlowSetupRequestList           QosFlowSetupRequestList           `,reject,mandatory`
	CommonNetworkInstance             CommonNetworkInstance             `,ignore,optional`
}
