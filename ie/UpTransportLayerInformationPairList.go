package ie

type UpTransportLayerInformationPairList struct {
	UpTransportLayerInformationPairItem UpTransportLayerInformationPairItem `bitstring:"sizeLB:0,sizeUB:150"`
}

type UpTransportLayerInformationPairItem struct {
	UlNgUUpTnlInformation UpTransportLayerInformation `bitstring:"sizeLB:0,sizeUB:150"`
	DlNgUUpTnlInformation UpTransportLayerInformation `bitstring:"sizeLB:0,sizeUB:150"`
}
