package ie

type BroadcastSessionReleaseRequest struct {
	MessageType  MessageType  `bitstring:"sizeLB:0,sizeUB:150"`
	MbsSessionId MbsSessionId `bitstring:"sizeLB:0,sizeUB:150"`
	Cause        Cause        `bitstring:"sizeLB:0,sizeUB:150"`
}
