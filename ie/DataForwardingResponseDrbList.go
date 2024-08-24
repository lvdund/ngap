package ie

type DataForwardingResponseDrbList struct {
DataForwardingResponseDrbItem	DataForwardingResponseDrbItem	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type DataForwardingResponseDrbItem struct {
DrbId	DrbId	//`bitstring:"sizeLB:0,sizeUB:150"`
DlForwardingUpTnlInformation	UpTransportLayerInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
UlForwardingUpTnlInformation	UpTransportLayerInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
}
