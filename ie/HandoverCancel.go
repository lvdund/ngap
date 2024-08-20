package ie

type HandoverCancel struct {
	MessageType *MessageType
	AmfUeNgapId *AmfUeNgapId
	RanUeNgapId *RanUeNgapId
	Cause       *Cause
}
