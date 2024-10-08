package ie

type PDUSessionResourceModifyConfirmTransfer struct {
	QosFlowModifyConfirmList      QosFlowModifyConfirmList                      `False,`
	ULNGUUPTNLInformation         UPTransportLayerInformation                   `False,`
	AdditionalNGUUPTNLInformation UPTransportLayerInformationPairList           `False,OPTIONAL`
	QosFlowFailedToModifyList     QosFlowListWithCause                          `False,OPTIONAL`
	IEExtensions                  PDUSessionResourceModifyConfirmTransferExtIEs `False,OPTIONAL`
}
