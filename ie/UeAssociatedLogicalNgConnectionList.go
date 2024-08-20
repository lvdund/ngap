package ie

type UeAssociatedLogicalNgConnectionList struct {
	UeAssociatedLogicalNgConnectionItem *UeAssociatedLogicalNgConnectionItem
}

type UeAssociatedLogicalNgConnectionItem struct {
	AmfUeNgapId *AmfUeNgapId
	RanUeNgapId *RanUeNgapId
}
