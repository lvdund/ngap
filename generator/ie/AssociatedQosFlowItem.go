package ie

type AssociatedQosFlowItem struct {
	QosFlowIdentifier QosFlowIdentifier           `False,`
	IEExtensions      AssociatedQosFlowItemExtIEs `False,OPTIONAL`
}
