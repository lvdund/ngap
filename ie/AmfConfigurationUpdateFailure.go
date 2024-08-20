package ie

type AmfConfigurationUpdateFailure struct {
	MessageType            *MessageType
	Cause                  *Cause
	TimeToWait             *TimeToWait
	CriticalityDiagnostics *CriticalityDiagnostics
}
