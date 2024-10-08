package ie

type CPTransportLayerInformation struct {
	EndpointIPAddress TransportLayerAddress             `False,`
	ChoiceExtensions  CPTransportLayerInformationExtIEs `False,`
}
