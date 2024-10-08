package ie

type HandoverCommandTransfer struct {
	DLForwardingUPTNLInformation  UPTransportLayerInformation   `False,OPTIONAL`
	QosFlowToBeForwardedList      QosFlowToBeForwardedList      `False,OPTIONAL`
	DataForwardingResponseDRBList DataForwardingResponseDRBList `False,OPTIONAL`
	IEExtensions                  HandoverCommandTransferExtIEs `False,OPTIONAL`
}
