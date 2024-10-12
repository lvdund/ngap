package ie

type UplinkUEAssociatedNRPPaTransportIEs struct {
	AMFUENGAPID AMFUENGAPID `,reject,mandatory`
	RANUENGAPID RANUENGAPID `,reject,mandatory`
	RoutingID   RoutingID   `,reject,mandatory`
	NRPPaPDU    NRPPaPDU    `,reject,mandatory`
}
