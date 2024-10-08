package ie

type PDUSessionResourceNotifyReleasedTransfer struct {
	Cause        Cause                                          `False,`
	IEExtensions PDUSessionResourceNotifyReleasedTransferExtIEs `False,OPTIONAL`
}
