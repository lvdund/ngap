package ie

type XnExtTLAItem struct {
	IPsecTLA     TransportLayerAddress `False,OPTIONAL`
	GTPTLAs      XnGTPTLAs             `False,OPTIONAL`
	IEExtensions XnExtTLAItemExtIEs    `False,OPTIONAL`
}
