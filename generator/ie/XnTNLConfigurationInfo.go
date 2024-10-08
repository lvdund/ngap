package ie

type XnTNLConfigurationInfo struct {
	XnTransportLayerAddresses         XnTLAs                       `False,`
	XnExtendedTransportLayerAddresses XnExtTLAs                    `False,OPTIONAL`
	IEExtensions                      XnTNLConfigurationInfoExtIEs `False,OPTIONAL`
}
