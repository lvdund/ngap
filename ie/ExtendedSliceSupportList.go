package ie

type ExtendedSliceSupportList struct {
	SliceSupportItem SliceSupportItem //`bitstring:"sizeLB:0,sizeUB:150"`
}

type SliceSupportItem struct {
	SNssai SNssai //`bitstring:"sizeLB:0,sizeUB:150"`
}
