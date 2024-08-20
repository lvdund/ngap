package ie

type HandoverCancelAcknowledge struct {
	MessageType            *MessageType
	AmfUeNgapId            *AmfUeNgapId
	RanUeNgapId            *RanUeNgapId
	CriticalityDiagnostics *CriticalityDiagnostics
}
