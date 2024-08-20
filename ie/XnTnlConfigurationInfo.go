package ie

type XnTnlConfigurationInfo struct {
	XnTransportLayerAddresses         *XnTransportLayerAddresses
	XnExtendedTransportLayerAddresses *XnExtendedTransportLayerAddresses
}

type XnTransportLayerAddresses struct {
	TransportLayerAddress *TransportLayerAddress
}

type XnExtendedTransportLayerAddresses struct {
	IpSecTransportLayerAddress    *TransportLayerAddress
	XnGtpTransportLayerAddresses  *XnGtpTransportLayerAddresses
	XnSctpTransportLayerAddresses *XnSctpTransportLayerAddresses
}

type XnGtpTransportLayerAddresses struct {
	GtpTransportLayerAddress *TransportLayerAddress
}

type XnSctpTransportLayerAddresses struct {
	TransportLayerAddressSctp *TransportLayerAddress
}
