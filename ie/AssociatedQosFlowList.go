package ie

type AssociatedQosFlowList struct {
	AssociatedQosFlowItem AssociatedQosFlowItem //`bitstring:"sizeLB:0,sizeUB:150"`
}

type AssociatedQosFlowItem struct {
	QosFlowIdentifier            QosFlowIdentifier                //`bitstring:"sizeLB:0,sizeUB:150"`
	QosFlowMappingIndication     []byte                           //`bitstring:"sizeLB:0,sizeUB:150"`
	CurrentQosParametersSetIndex AlternativeQosParametersSetIndex //`bitstring:"sizeLB:0,sizeUB:150"`
}
