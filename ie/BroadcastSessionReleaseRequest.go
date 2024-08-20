package ie

type BroadcastSessionReleaseRequest struct {
	MessageType  *MessageType
	MbsSessionId *MbsSessionId
	Cause        *Cause
}
