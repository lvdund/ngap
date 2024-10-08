package ie

type PDUSessionResourceModifyResponseTransfer struct {
	DLNGUUPTNLInformation                UPTransportLayerInformation                    `False,OPTIONAL`
	ULNGUUPTNLInformation                UPTransportLayerInformation                    `False,OPTIONAL`
	QosFlowAddOrModifyResponseList       QosFlowAddOrModifyResponseList                 `False,OPTIONAL`
	AdditionalDLQosFlowPerTNLInformation QosFlowPerTNLInformationList                   `False,OPTIONAL`
	QosFlowFailedToAddOrModifyList       QosFlowListWithCause                           `False,OPTIONAL`
	IEExtensions                         PDUSessionResourceModifyResponseTransferExtIEs `False,OPTIONAL`
}
