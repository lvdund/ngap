package ie

type PduSessionResourceModifyConfirm struct {
	MessageType                          *MessageType
	AmfUeNgapId                          *AmfUeNgapId
	RanUeNgapId                          *RanUeNgapId
	PduSessionResourceModifyConfirmList  *[]PduSessionResourceModifyConfirmItem
	PduSessionResourceFailedToModifyList *[]PduSessionResourceFailedToModifyItem
	CriticalityDiagnostics               *CriticalityDiagnostics
}

type PduSessionResourceModifyConfirmItem struct {
	PduSessionId                            *PduSessionId
	PduSessionResourceModifyConfirmTransfer *[]byte
}
