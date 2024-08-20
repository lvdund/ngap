package ie

type UeContextModificationFailure struct {
	MessageType            *MessageType
	AmfUeNgapId            *AmfUeNgapId
	RanUeNgapId            *RanUeNgapId
	Cause                  *Cause
	CriticalityDiagnostics *CriticalityDiagnostics
}
