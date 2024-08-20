package ie

type PduSessionResourceNotifyTransfer struct {
QosFlowNotifyList	*[]QosFlowNotifyItem
QosFlowReleasedList	*QosFlowListWithCause
SecondaryRatUsageInformation	*SecondaryRatUsageInformation
QosFlowFeedbackList	*[]QosFlowFeedbackItem
}

type QosFlowNotifyItem struct {
QosFlowIdentifier	*QosFlowIdentifier
NotificationCause	*[]byte
CurrentQosParametersSetIndex	*AlternativeQosParametersSetNotifyIndex
}

type QosFlowFeedbackItem struct {
QosFlowIdentifier	*QosFlowIdentifier
UpdateFeedback	*[]byte
CnPacketDelayBudgetDownlink	*ExtendedPacketDelayBudget
CnPacketDelayBudgetUplink	*ExtendedPacketDelayBudget
}
