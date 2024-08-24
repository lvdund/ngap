package ie

type UpTransportLayerInformationList struct {
UpTransportLayerInformationItem	UpTransportLayerInformationItem	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type UpTransportLayerInformationItem struct {
NgUUpTnlInformation	UpTransportLayerInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
CommonNetworkInstance	CommonNetworkInstance	//`bitstring:"sizeLB:0,sizeUB:150"`
}
