package ie

type BroadcastSessionModificationResponse struct {
MessageType	*MessageType
MbsSessionId	*MbsSessionId
MbsSessionModificationResponseTransfer	*[]byte
CriticalityDiagnostics	*CriticalityDiagnostics
}
