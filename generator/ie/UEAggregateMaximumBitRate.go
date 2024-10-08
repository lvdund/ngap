package ie

type UEAggregateMaximumBitRate struct {
	UEAggregateMaximumBitRateDL BitRate                         `False,`
	UEAggregateMaximumBitRateUL BitRate                         `False,`
	IEExtensions                UEAggregateMaximumBitRateExtIEs `False,OPTIONAL`
}
