package ie

type MulticastSessionDeactivationResponse struct {
MessageType	MessageType	//`bitstring:"sizeLB:0,sizeUB:150"`
MbsSessionId	MbsSessionId	//`bitstring:"sizeLB:0,sizeUB:150"`
CriticalityDiagnostics	CriticalityDiagnostics	//`bitstring:"sizeLB:0,sizeUB:150"`
}
