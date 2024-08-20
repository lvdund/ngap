package ie

type BroadcastSessionModificationFailure struct {
	MessageType                           *MessageType
	MbsSessionId                          *MbsSessionId
	MbsSessionModificationFailureTransfer *[]byte
	Cause                                 *Cause
	CriticalityDiagnostics                *CriticalityDiagnostics
}
