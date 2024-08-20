package ie

type UpTransportLayerInformationPairList struct {
	UpTransportLayerInformationPairItem *UpTransportLayerInformationPairItem
}

type UpTransportLayerInformationPairItem struct {
	UlNgUUpTnlInformation *UpTransportLayerInformation
	DlNgUUpTnlInformation *UpTransportLayerInformation
}
