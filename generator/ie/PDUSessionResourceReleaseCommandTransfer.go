package ie

type PDUSessionResourceReleaseCommandTransfer struct {
	Cause        Cause                                          `False,`
	IEExtensions PDUSessionResourceReleaseCommandTransferExtIEs `False,OPTIONAL`
}
