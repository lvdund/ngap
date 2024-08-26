package ie

type InitialContextSetupFailure struct {
	MessageType                         MessageType                           `bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                         AmfUeNgapId                           `bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                         RanUeNgapId                           `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceFailedToSetupList []PduSessionResourceFailedToSetupItem `bitstring:"sizeLB:0,sizeUB:150"`
	Cause                               Cause                                 `bitstring:"sizeLB:0,sizeUB:150"`
	CriticalityDiagnostics              CriticalityDiagnostics                `bitstring:"sizeLB:0,sizeUB:150"`
}
