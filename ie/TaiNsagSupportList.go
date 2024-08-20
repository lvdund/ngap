package ie

type TaiNsagSupportList struct {
	TaiNsagSupportItem *TaiNsagSupportItem
}

type TaiNsagSupportItem struct {
	NsagId               uint8 //`bitstring:"sizeLB:0,sizeUB:255"`
	NsagSliceSupportList *ExtendedSliceSupportList
}
