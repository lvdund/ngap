package ie

type DownlinkUEAssociatedNRPPaTransportIEs struct {
	AMFUENGAPID AMFUENGAPID `,reject,mandatory`
	RANUENGAPID RANUENGAPID `,reject,mandatory`
	RoutingID   RoutingID   `,reject,mandatory`
	NRPPaPDU    NRPPaPDU    `,reject,mandatory`
}
