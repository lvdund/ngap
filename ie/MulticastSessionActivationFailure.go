package ie

type MulticastSessionActivationFailure struct {
	MessageType            *MessageType
	MbsSessionId           *MbsSessionId
	Cause                  *Cause
	CriticalityDiagnostics *CriticalityDiagnostics
}
