package ie

type BroadcastSessionSetupResponse struct {
MessageType	*MessageType
MbsSessionId	*MbsSessionId
MbsSessionSetupResponseTransfer	*[]byte
CriticalityDiagnostics	*CriticalityDiagnostics
}
