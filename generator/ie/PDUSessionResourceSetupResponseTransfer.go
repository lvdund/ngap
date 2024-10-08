package ie

type PDUSessionResourceSetupResponseTransfer struct {
	DLQosFlowPerTNLInformation           QosFlowPerTNLInformation                      `True,`
	AdditionalDLQosFlowPerTNLInformation QosFlowPerTNLInformationList                  `False,OPTIONAL`
	SecurityResult                       SecurityResult                                `True,OPTIONAL`
	QosFlowFailedToSetupList             QosFlowListWithCause                          `False,OPTIONAL`
	IEExtensions                         PDUSessionResourceSetupResponseTransferExtIEs `False,OPTIONAL`
}
