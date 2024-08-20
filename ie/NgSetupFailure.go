package ie

type NgSetupFailure struct {
	MessageType            *MessageType
	Cause                  *Cause
	TimeToWait             *TimeToWait
	CriticalityDiagnostics *CriticalityDiagnostics
}
