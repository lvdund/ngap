package ie

type BroadcastSessionSetupFailure struct {
	MessageType                    *MessageType
	MbsSessionId                   *MbsSessionId
	MbsSessionSetupFailureTransfer *[]byte
	Cause                          *Cause
	CriticalityDiagnostics         *CriticalityDiagnostics
}
