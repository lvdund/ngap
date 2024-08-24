package ie

import "ngap/aper"

type CpTransportLayerInformation struct {
ChoiceCpTransportLayerInformation	ChoiceCpTransportLayerInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceCpTransportLayerInformation struct {
EndpointIpAddress	EndpointIpAddress	//`bitstring:"sizeLB:0,sizeUB:150"`
EndpointIpAddressAndPort	EndpointIpAddressAndPort	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type EndpointIpAddress struct {
EndpointIpAddress	TransportLayerAddress	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type EndpointIpAddressAndPort struct {
EndpointIpAddress	TransportLayerAddress	//`bitstring:"sizeLB:0,sizeUB:150"`
PortNumber	aper.OctetString	//`octetstring:"sizeLB:2,sizeUB:2"`
}
