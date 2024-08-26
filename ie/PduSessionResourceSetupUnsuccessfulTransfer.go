package ie

type PduSessionResourceSetupUnsuccessfulTransfer struct {
	Cause                  Cause                  `bitstring:"sizeLB:0,sizeUB:150"`
	CriticalityDiagnostics CriticalityDiagnostics `bitstring:"sizeLB:0,sizeUB:150"`
}
