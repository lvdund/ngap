package ie

type ULNGUUPTNLModifyItem struct {
	ULNGUUPTNLInformation UPTransportLayerInformation `False,`
	DLNGUUPTNLInformation UPTransportLayerInformation `False,`
	IEExtensions          ULNGUUPTNLModifyItemExtIEs  `False,OPTIONAL`
}
