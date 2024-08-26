package ie

type SliceOverloadList struct {
	SliceOverloadItem SliceOverloadItem `bitstring:"sizeLB:0,sizeUB:150"`
}

type SliceOverloadItem struct {
	SNssai SNssai `bitstring:"sizeLB:0,sizeUB:150"`
}
