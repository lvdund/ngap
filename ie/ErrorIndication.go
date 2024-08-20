package ie

type ErrorIndication struct {
	MessageType            *MessageType
	AmfUeNgapId            *AmfUeNgapId
	RanUeNgapId            *RanUeNgapId
	Cause                  *Cause
	CriticalityDiagnostics *CriticalityDiagnostics
	Ie5GSTmsi              *Ie5GSTmsi
}
