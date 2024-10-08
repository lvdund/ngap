package ie

type QosFlowAddOrModifyRequestItem struct {
	QosFlowIdentifier         QosFlowIdentifier                   `False,`
	QosFlowLevelQosParameters QosFlowLevelQosParameters           `True,OPTIONAL`
	ERABID                    ERABID                              `False,OPTIONAL`
	IEExtensions              QosFlowAddOrModifyRequestItemExtIEs `False,OPTIONAL`
}
