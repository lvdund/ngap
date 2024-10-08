package ie

type AmfConfigurationUpdateFailure struct {
	MessageType            MessageType            `bitstring:"sizeLB:0,sizeUB:150"`
	Cause                  Cause                  `bitstring:"sizeLB:0,sizeUB:150"`
	TimeToWait             TimeToWait             `bitstring:"sizeLB:0,sizeUB:150"`
	CriticalityDiagnostics CriticalityDiagnostics `bitstring:"sizeLB:0,sizeUB:150"`
}
