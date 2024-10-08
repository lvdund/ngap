package ie

type HandoverRequestAcknowledgeTransferExtIEs struct {
	AdditionalDLUPTNLInformationForHOList  AdditionalDLUPTNLInformationForHOList `,ignore,optional`
	ULForwardingUPTNLInformation           UPTransportLayerInformation           `,reject,optional`
	AdditionalULForwardingUPTNLInformation UPTransportLayerInformationList       `,reject,optional`
}
