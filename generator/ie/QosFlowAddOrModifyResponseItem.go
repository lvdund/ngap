package ie

type QosFlowAddOrModifyResponseItem struct {
	QosFlowIdentifier QosFlowIdentifier                    `False,`
	IEExtensions      QosFlowAddOrModifyResponseItemExtIEs `False,OPTIONAL`
}
