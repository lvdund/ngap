package ie

type QosFlowPerTNLInformation struct {
	UPTransportLayerInformation UPTransportLayerInformation    `False,`
	AssociatedQosFlowList       AssociatedQosFlowList          `False,`
	IEExtensions                QosFlowPerTNLInformationExtIEs `False,OPTIONAL`
}
