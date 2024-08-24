package ie

type UeAggregateMaximumBitRate struct {
	UeAggregateMaximumBitRateDownlink BitRate //`bitstring:"sizeLB:0,sizeUB:150"`
	UeAggregateMaximumBitRateUplink   BitRate //`bitstring:"sizeLB:0,sizeUB:150"`
}
