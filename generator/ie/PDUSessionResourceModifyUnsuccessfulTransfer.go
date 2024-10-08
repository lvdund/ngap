package ie

type PDUSessionResourceModifyUnsuccessfulTransfer struct {
	Cause                  Cause                                              `False,`
	CriticalityDiagnostics CriticalityDiagnostics                             `True,OPTIONAL`
	IEExtensions           PDUSessionResourceModifyUnsuccessfulTransferExtIEs `False,OPTIONAL`
}
