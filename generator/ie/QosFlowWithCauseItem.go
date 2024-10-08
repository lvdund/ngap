package ie

type QosFlowWithCauseItem struct {
	QosFlowIdentifier QosFlowIdentifier          `False,`
	Cause             Cause                      `False,`
	IEExtensions      QosFlowWithCauseItemExtIEs `False,OPTIONAL`
}
