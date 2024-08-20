package ie

type DownlinkRanEarlyStatusTransfer struct {
	MessageType                             *MessageType
	AmfUeNgapId                             *AmfUeNgapId
	RanUeNgapId                             *RanUeNgapId
	EarlyStatusTransferTransparentContainer *EarlyStatusTransferTransparentContainer
}
