package ie

type HandoverFailure struct {
	MessageType                               MessageType                               `bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                               AmfUeNgapId                               `bitstring:"sizeLB:0,sizeUB:150"`
	Cause                                     Cause                                     `bitstring:"sizeLB:0,sizeUB:150"`
	CriticalityDiagnostics                    CriticalityDiagnostics                    `bitstring:"sizeLB:0,sizeUB:150"`
	TargetToSourceFailureTransparentContainer TargetToSourceFailureTransparentContainer `bitstring:"sizeLB:0,sizeUB:150"`
}
