package ie

type HandoverRequestAcknowledgeTransfer struct {
	DLNGUUPTNLInformation         UPTransportLayerInformation              `False,`
	DLForwardingUPTNLInformation  UPTransportLayerInformation              `False,OPTIONAL`
	SecurityResult                SecurityResult                           `True,OPTIONAL`
	QosFlowSetupResponseList      QosFlowListWithDataForwarding            `False,`
	QosFlowFailedToSetupList      QosFlowListWithCause                     `False,OPTIONAL`
	DataForwardingResponseDRBList DataForwardingResponseDRBList            `False,OPTIONAL`
	IEExtensions                  HandoverRequestAcknowledgeTransferExtIEs `False,OPTIONAL`
}
