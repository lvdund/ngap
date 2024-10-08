package ie

type AdditionalDLUPTNLInformationForHOItem struct {
	AdditionalDLNGUUPTNLInformation        UPTransportLayerInformation                 `False,`
	AdditionalQosFlowSetupResponseList     QosFlowListWithDataForwarding               `False,`
	AdditionalDLForwardingUPTNLInformation UPTransportLayerInformation                 `False,OPTIONAL`
	IEExtensions                           AdditionalDLUPTNLInformationForHOItemExtIEs `False,OPTIONAL`
}
