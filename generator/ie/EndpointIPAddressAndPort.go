package ie

type EndpointIPAddressAndPort struct {
	EndpointIPAddress TransportLayerAddress          `False,`
	PortNumber        PortNumber                     `False,`
	IEExtensions      EndpointIPAddressAndPortExtIEs `False,OPTIONAL`
}
