package ie

type UeContextSuspendResponse struct {
	MessageType            *MessageType
	AmfUeNgapId            *AmfUeNgapId
	RanUeNgapId            *RanUeNgapId
	SecurityContext        *SecurityContext
	CriticalityDiagnostics *CriticalityDiagnostics
}
