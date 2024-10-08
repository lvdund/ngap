package ie

type QosFlowItemWithDataForwarding struct {
	QosFlowIdentifier      QosFlowIdentifier                   `False,`
	DataForwardingAccepted DataForwardingAccepted              `False,OPTIONAL`
	IEExtensions           QosFlowItemWithDataForwardingExtIEs `False,OPTIONAL`
}
