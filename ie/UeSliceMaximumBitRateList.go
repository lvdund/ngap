package ie

type UeSliceMaximumBitRateList struct {
	UeSliceMaximumBitRateItem *UeSliceMaximumBitRateItem
}

type UeSliceMaximumBitRateItem struct {
	SNssai                        *SNssai
	UeSliceMaximumBitRateDownlink *BitRate
	UeSliceMaximumBitRateUplink   *BitRate
}
