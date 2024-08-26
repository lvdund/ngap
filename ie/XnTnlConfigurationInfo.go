package ie

type XnTnlConfigurationInfo struct {
	XnTransportLayerAddresses         XnTransportLayerAddresses         `bitstring:"sizeLB:0,sizeUB:150"`
	XnExtendedTransportLayerAddresses XnExtendedTransportLayerAddresses `bitstring:"sizeLB:0,sizeUB:150"`
}

type XnTransportLayerAddresses struct {
	TransportLayerAddress TransportLayerAddress `bitstring:"sizeLB:0,sizeUB:150"`
}

type XnExtendedTransportLayerAddresses struct {
	IpSecTransportLayerAddress    TransportLayerAddress         `bitstring:"sizeLB:0,sizeUB:150"`
	XnGtpTransportLayerAddresses  XnGtpTransportLayerAddresses  `bitstring:"sizeLB:0,sizeUB:150"`
	XnSctpTransportLayerAddresses XnSctpTransportLayerAddresses `bitstring:"sizeLB:0,sizeUB:150"`
}

type XnGtpTransportLayerAddresses struct {
	GtpTransportLayerAddress TransportLayerAddress `bitstring:"sizeLB:0,sizeUB:150"`
}

type XnSctpTransportLayerAddresses struct {
	TransportLayerAddressSctp TransportLayerAddress `bitstring:"sizeLB:0,sizeUB:150"`
}
