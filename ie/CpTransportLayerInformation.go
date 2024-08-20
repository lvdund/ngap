package ie

type CpTransportLayerInformation struct {
	ChoiceCpTransportLayerInformation *ChoiceCpTransportLayerInformation
}

type ChoiceCpTransportLayerInformation struct {
	EndpointIpAddress        *EndpointIpAddress
	EndpointIpAddressAndPort *EndpointIpAddressAndPort
}

type EndpointIpAddress struct {
	EndpointIpAddress *TransportLayerAddress
}

type EndpointIpAddressAndPort struct {
	EndpointIpAddress *TransportLayerAddress
	PortNumber        []byte //`bitstring:"sizeLB:2,sizeUB:2"`
}
