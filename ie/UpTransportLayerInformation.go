package ie

type UpTransportLayerInformation struct {
	ChoiceUpTransportLayerInformation ChoiceUpTransportLayerInformation //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceUpTransportLayerInformation struct {
	GtpTunnel GtpTunnel //`bitstring:"sizeLB:0,sizeUB:150"`
}

type GtpTunnel struct {
	EndpointIpAddress TransportLayerAddress //`bitstring:"sizeLB:0,sizeUB:150"`
	GtpTeid           GtpTeid               //`bitstring:"sizeLB:0,sizeUB:150"`
}
