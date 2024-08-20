package ie

type UplinkRanStatusTransfer struct {
	MessageType                           *MessageType
	AmfUeNgapId                           *AmfUeNgapId
	RanUeNgapId                           *RanUeNgapId
	RanStatusTransferTransparentContainer *RanStatusTransferTransparentContainer
}
