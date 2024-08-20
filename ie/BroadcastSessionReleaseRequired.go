package ie

type BroadcastSessionReleaseRequired struct {
	MessageType  *MessageType
	MbsSessionId *MbsSessionId
	Cause        *Cause
}
