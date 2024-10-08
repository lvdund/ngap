package ie

type PDUSessionAggregateMaximumBitRate struct {
	PDUSessionAggregateMaximumBitRateDL BitRate                                 `False,`
	PDUSessionAggregateMaximumBitRateUL BitRate                                 `False,`
	IEExtensions                        PDUSessionAggregateMaximumBitRateExtIEs `False,OPTIONAL`
}
