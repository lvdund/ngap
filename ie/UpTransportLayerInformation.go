package ie

type UpTransportLayerInformation struct {
ChoiceUpTransportLayerInformation	*ChoiceUpTransportLayerInformation
}

type ChoiceUpTransportLayerInformation struct {
GtpTunnel	*GtpTunnel
}

type GtpTunnel struct {
EndpointIpAddress	*TransportLayerAddress
GtpTeid	*GtpTeid
}
