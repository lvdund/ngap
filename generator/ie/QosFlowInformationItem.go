package ie

type QosFlowInformationItem struct {
	QosFlowIdentifier QosFlowIdentifier            `False,`
	DLForwarding      DLForwarding                 `False,OPTIONAL`
	IEExtensions      QosFlowInformationItemExtIEs `False,OPTIONAL`
}
