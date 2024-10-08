package ie

type PathSwitchRequestTransfer struct {
	DLNGUUPTNLInformation        UPTransportLayerInformation     `False,`
	DLNGUTNLInformationReused    DLNGUTNLInformationReused       `False,OPTIONAL`
	UserPlaneSecurityInformation UserPlaneSecurityInformation    `True,OPTIONAL`
	QosFlowAcceptedList          QosFlowAcceptedList             `False,`
	IEExtensions                 PathSwitchRequestTransferExtIEs `False,OPTIONAL`
}
