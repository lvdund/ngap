package ie

type PDUSessionResourceSetupUnsuccessfulTransfer struct {
	Cause                  Cause                                             `False,`
	CriticalityDiagnostics CriticalityDiagnostics                            `True,OPTIONAL`
	IEExtensions           PDUSessionResourceSetupUnsuccessfulTransferExtIEs `False,OPTIONAL`
}
