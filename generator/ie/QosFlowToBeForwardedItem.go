package ie

type QosFlowToBeForwardedItem struct {
	QosFlowIdentifier QosFlowIdentifier              `False,`
	IEExtensions      QosFlowToBeForwardedItemExtIEs `False,OPTIONAL`
}
