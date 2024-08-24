package ie

type UplinkRanEarlyStatusTransfer struct {
MessageType	MessageType	//`bitstring:"sizeLB:0,sizeUB:150"`
AmfUeNgapId	AmfUeNgapId	//`bitstring:"sizeLB:0,sizeUB:150"`
RanUeNgapId	RanUeNgapId	//`bitstring:"sizeLB:0,sizeUB:150"`
EarlyStatusTransferTransparentContainer	EarlyStatusTransferTransparentContainer	//`bitstring:"sizeLB:0,sizeUB:150"`
}
