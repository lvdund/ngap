package ie

type QosFlowListWithCause struct {
QosFlowItem	*QosFlowItem
}

type QosFlowItem struct {
QosFlowIdentifier	*QosFlowIdentifier
Cause	*Cause
}
