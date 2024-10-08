package ie

type HandoverCommandTransferExtIEs struct {
	AdditionalDLForwardingUPTNLInformation QosFlowPerTNLInformationList    `,ignore,optional`
	ULForwardingUPTNLInformation           UPTransportLayerInformation     `,reject,optional`
	AdditionalULForwardingUPTNLInformation UPTransportLayerInformationList `,reject,optional`
}
