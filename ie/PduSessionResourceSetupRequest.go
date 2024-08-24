package ie

type PduSessionResourceSetupRequest struct {
	MessageType                        MessageType                          //`bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                        AmfUeNgapId                          //`bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                        RanUeNgapId                          //`bitstring:"sizeLB:0,sizeUB:150"`
	RanPagingPriority                  RanPagingPriority                    //`bitstring:"sizeLB:0,sizeUB:150"`
	NasPdu                             NasPdu                               //`bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceSetupRequestList []PduSessionResourceSetupRequestItem //`bitstring:"sizeLB:0,sizeUB:150"`
	UeAggregateMaximumBitRate          UeAggregateMaximumBitRate            //`bitstring:"sizeLB:0,sizeUB:150"`
	UeSliceMaximumBitRateList          UeSliceMaximumBitRateList            //`bitstring:"sizeLB:0,sizeUB:150"`
}
