package ie

type QosFlowListWithCause struct {
QosFlowItem	QosFlowItem	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type QosFlowItem struct {
QosFlowIdentifier	QosFlowIdentifier	//`bitstring:"sizeLB:0,sizeUB:150"`
Cause	Cause	//`bitstring:"sizeLB:0,sizeUB:150"`
}
