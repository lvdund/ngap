package ie

type QosFlowListWithDataForwarding struct {
QosFlowItemWithDataForwarding	QosFlowItemWithDataForwarding	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type QosFlowItemWithDataForwarding struct {
QosFlowIdentifier	QosFlowIdentifier	//`bitstring:"sizeLB:0,sizeUB:150"`
DataForwardingAccepted	DataForwardingAccepted	//`bitstring:"sizeLB:0,sizeUB:150"`
CurrentQosParametersSetIndex	AlternativeQosParametersSetIndex	//`bitstring:"sizeLB:0,sizeUB:150"`
}
