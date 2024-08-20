package ie

type AssociatedQosFlowList struct {
AssociatedQosFlowItem	*AssociatedQosFlowItem
}

type AssociatedQosFlowItem struct {
QosFlowIdentifier	*QosFlowIdentifier
QosFlowMappingIndication	*[]byte
CurrentQosParametersSetIndex	*AlternativeQosParametersSetIndex
}
