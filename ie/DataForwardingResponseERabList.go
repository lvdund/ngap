package ie

type DataForwardingResponseERabList struct {
	ERabId                       ERabId                      //`bitstring:"sizeLB:0,sizeUB:150"`
	DlForwardingUpTnlInformation UpTransportLayerInformation //`bitstring:"sizeLB:0,sizeUB:150"`
}
