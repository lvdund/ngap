package ie

type UPTransportLayerInformationItem struct {
	NGUUPTNLInformation UPTransportLayerInformation           `False,`
	IEExtensions        UPTransportLayerInformationItemExtIEs `False,OPTIONAL`
}
