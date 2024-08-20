package ie

type BroadcastSessionReleaseResponse struct {
	MessageType                       *MessageType
	MbsSessionId                      *MbsSessionId
	MbsSessionReleaseResponseTransfer *[]byte
	CriticalityDiagnostics            *CriticalityDiagnostics
}
