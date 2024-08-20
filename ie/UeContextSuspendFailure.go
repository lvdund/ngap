package ie

type UeContextSuspendFailure struct {
	MessageType            *MessageType
	AmfUeNgapId            *AmfUeNgapId
	RanUeNgapId            *RanUeNgapId
	Cause                  *Cause
	CriticalityDiagnostics *CriticalityDiagnostics
}
