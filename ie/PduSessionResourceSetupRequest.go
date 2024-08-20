package ie

type PduSessionResourceSetupRequest struct {
	MessageType                        *MessageType
	AmfUeNgapId                        *AmfUeNgapId
	RanUeNgapId                        *RanUeNgapId
	RanPagingPriority                  *RanPagingPriority
	NasPdu                             *NasPdu
	PduSessionResourceSetupRequestList *[]PduSessionResourceSetupRequestItem
	UeAggregateMaximumBitRate          *UeAggregateMaximumBitRate
	UeSliceMaximumBitRateList          *UeSliceMaximumBitRateList
}
