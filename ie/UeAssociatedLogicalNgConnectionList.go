package ie

type UeAssociatedLogicalNgConnectionList struct {
	UeAssociatedLogicalNgConnectionItem UeAssociatedLogicalNgConnectionItem `bitstring:"sizeLB:0,sizeUB:150"`
}

type UeAssociatedLogicalNgConnectionItem struct {
	AmfUeNgapId AmfUeNgapId `bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId RanUeNgapId `bitstring:"sizeLB:0,sizeUB:150"`
}
