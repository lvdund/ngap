package ie

type UplinkRanStatusTransfer struct {
	MessageType                           MessageType                           //`bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                           AmfUeNgapId                           //`bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                           RanUeNgapId                           //`bitstring:"sizeLB:0,sizeUB:150"`
	RanStatusTransferTransparentContainer RanStatusTransferTransparentContainer //`bitstring:"sizeLB:0,sizeUB:150"`
}
