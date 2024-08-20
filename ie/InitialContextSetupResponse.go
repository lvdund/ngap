package ie

type InitialContextSetupResponse struct {
	MessageType                         *MessageType
	AmfUeNgapId                         *AmfUeNgapId
	RanUeNgapId                         *RanUeNgapId
	PduSessionResourceSetupResponseList *[]PduSessionResourceSetupResponseItem
	PduSessionResourceFailedToSetupList *[]PduSessionResourceFailedToSetupItem
	CriticalityDiagnostics              *CriticalityDiagnostics
}

type PduSessionResourceSetupResponseItem struct {
	PduSessionId                            *PduSessionId
	PduSessionResourceSetupResponseTransfer *[]byte
}
