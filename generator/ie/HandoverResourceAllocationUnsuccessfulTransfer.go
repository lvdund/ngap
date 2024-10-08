package ie

type HandoverResourceAllocationUnsuccessfulTransfer struct {
	Cause                  Cause                                                `False,`
	CriticalityDiagnostics CriticalityDiagnostics                               `True,OPTIONAL`
	IEExtensions           HandoverResourceAllocationUnsuccessfulTransferExtIEs `False,OPTIONAL`
}
