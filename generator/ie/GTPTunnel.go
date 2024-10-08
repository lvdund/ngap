package ie

type GTPTunnel struct {
	TransportLayerAddress TransportLayerAddress `False,`
	GTPTEID               GTPTEID               `False,`
	IEExtensions          GTPTunnelExtIEs       `False,OPTIONAL`
}
