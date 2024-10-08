package ie

type PDUSessionResourceModifyIndicationTransfer struct {
	DLQosFlowPerTNLInformation           QosFlowPerTNLInformation                         `True,`
	AdditionalDLQosFlowPerTNLInformation QosFlowPerTNLInformationList                     `False,OPTIONAL`
	IEExtensions                         PDUSessionResourceModifyIndicationTransferExtIEs `False,OPTIONAL`
}
