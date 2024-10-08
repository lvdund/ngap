package ie

type UPTransportLayerInformation struct {
	GTPTunnel        GTPTunnel                         `True,`
	ChoiceExtensions UPTransportLayerInformationExtIEs `False,`
}
