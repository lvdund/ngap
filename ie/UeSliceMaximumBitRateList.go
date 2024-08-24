package ie

type UeSliceMaximumBitRateList struct {
	UeSliceMaximumBitRateItem UeSliceMaximumBitRateItem //`bitstring:"sizeLB:0,sizeUB:150"`
}

type UeSliceMaximumBitRateItem struct {
	SNssai                        SNssai  //`bitstring:"sizeLB:0,sizeUB:150"`
	UeSliceMaximumBitRateDownlink BitRate //`bitstring:"sizeLB:0,sizeUB:150"`
	UeSliceMaximumBitRateUplink   BitRate //`bitstring:"sizeLB:0,sizeUB:150"`
}
