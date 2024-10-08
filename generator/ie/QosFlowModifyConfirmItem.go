package ie

type QosFlowModifyConfirmItem struct {
	QosFlowIdentifier QosFlowIdentifier              `False,`
	IEExtensions      QosFlowModifyConfirmItemExtIEs `False,OPTIONAL`
}
