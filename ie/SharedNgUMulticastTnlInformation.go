package ie

type SharedNgUMulticastTnlInformation struct {
IpMulticastAddress	TransportLayerAddress	//`bitstring:"sizeLB:0,sizeUB:150"`
IpSourceAddress	TransportLayerAddress	//`bitstring:"sizeLB:0,sizeUB:150"`
GtpTeidAt5Gc	GtpTeid	//`bitstring:"sizeLB:0,sizeUB:150"`
}
