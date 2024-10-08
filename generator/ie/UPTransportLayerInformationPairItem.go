package ie

type UPTransportLayerInformationPairItem struct {
	ULNGUUPTNLInformation UPTransportLayerInformation               `False,`
	DLNGUUPTNLInformation UPTransportLayerInformation               `False,`
	IEExtensions          UPTransportLayerInformationPairItemExtIEs `False,OPTIONAL`
}
