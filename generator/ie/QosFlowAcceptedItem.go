package ie

type QosFlowAcceptedItem struct {
	QosFlowIdentifier QosFlowIdentifier         `False,`
	IEExtensions      QosFlowAcceptedItemExtIEs `False,OPTIONAL`
}
