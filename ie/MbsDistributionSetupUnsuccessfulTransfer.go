package ie

type MbsDistributionSetupUnsuccessfulTransfer struct {
	MbsSessionId           MbsSessionId           `bitstring:"sizeLB:0,sizeUB:150"`
	MbsAreaSessionId       MbsAreaSessionId       `bitstring:"sizeLB:0,sizeUB:150"`
	Cause                  Cause                  `bitstring:"sizeLB:0,sizeUB:150"`
	CriticalityDiagnostics CriticalityDiagnostics `bitstring:"sizeLB:0,sizeUB:150"`
}
