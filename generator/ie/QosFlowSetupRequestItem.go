package ie

type QosFlowSetupRequestItem struct {
	QosFlowIdentifier         QosFlowIdentifier             `False,`
	QosFlowLevelQosParameters QosFlowLevelQosParameters     `True,`
	ERABID                    ERABID                        `False,OPTIONAL`
	IEExtensions              QosFlowSetupRequestItemExtIEs `False,OPTIONAL`
}
