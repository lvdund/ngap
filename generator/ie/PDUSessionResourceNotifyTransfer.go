package ie

type PDUSessionResourceNotifyTransfer struct {
	QosFlowNotifyList   QosFlowNotifyList                      `False,OPTIONAL`
	QosFlowReleasedList QosFlowListWithCause                   `False,OPTIONAL`
	IEExtensions        PDUSessionResourceNotifyTransferExtIEs `False,OPTIONAL`
}
