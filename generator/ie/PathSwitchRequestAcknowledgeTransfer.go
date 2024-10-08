package ie

type PathSwitchRequestAcknowledgeTransfer struct {
	ULNGUUPTNLInformation UPTransportLayerInformation                `False,OPTIONAL`
	SecurityIndication    SecurityIndication                         `True,OPTIONAL`
	IEExtensions          PathSwitchRequestAcknowledgeTransferExtIEs `False,OPTIONAL`
}
