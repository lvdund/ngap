package ie

type DataForwardingResponseDrbList struct {
	DataForwardingResponseDrbItem *DataForwardingResponseDrbItem
}

type DataForwardingResponseDrbItem struct {
	DrbId                        *DrbId
	DlForwardingUpTnlInformation *UpTransportLayerInformation
	UlForwardingUpTnlInformation *UpTransportLayerInformation
}
