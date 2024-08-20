package ie

type RanConfigurationUpdateFailure struct {
	MessageType            *MessageType
	Cause                  *Cause
	TimeToWait             *TimeToWait
	CriticalityDiagnostics *CriticalityDiagnostics
}
