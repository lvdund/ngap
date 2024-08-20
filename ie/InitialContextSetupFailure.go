package ie

type InitialContextSetupFailure struct {
	MessageType                         *MessageType
	AmfUeNgapId                         *AmfUeNgapId
	RanUeNgapId                         *RanUeNgapId
	PduSessionResourceFailedToSetupList *[]PduSessionResourceFailedToSetupItem
	Cause                               *Cause
	CriticalityDiagnostics              *CriticalityDiagnostics
}
