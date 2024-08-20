package ie

type HandoverRequired struct {
	MessageType                        *MessageType
	AmfUeNgapId                        *AmfUeNgapId
	RanUeNgapId                        *RanUeNgapId
	HandoverType                       *HandoverType
	Cause                              *Cause
	TargetId                           *TargetId
	DirectForwardingPathAvailability   *DirectForwardingPathAvailability
	PduSessionResourceList             *[]PduSessionResourceItem
	SourceToTargetTransparentContainer *SourceToTargetTransparentContainer
}
