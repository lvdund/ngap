package ie

type UeContextReleaseRequest struct {
	MessageType            *MessageType
	AmfUeNgapId            *AmfUeNgapId
	RanUeNgapId            *RanUeNgapId
	PduSessionResourceList *[]PduSessionResourceItem
	Cause                  *Cause
}

type PduSessionResourceItem struct {
	PduSessionId *PduSessionId
}
