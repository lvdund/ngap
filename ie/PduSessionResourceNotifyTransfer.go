package ie

import "ngap/aper"

type PduSessionResourceNotifyTransfer struct {
	QosFlowNotifyList            []QosFlowNotifyItem          `bitstring:"sizeLB:0,sizeUB:150"`
	QosFlowReleasedList          QosFlowListWithCause         `bitstring:"sizeLB:0,sizeUB:150"`
	SecondaryRatUsageInformation SecondaryRatUsageInformation `bitstring:"sizeLB:0,sizeUB:150"`
	QosFlowFeedbackList          []QosFlowFeedbackItem        `bitstring:"sizeLB:0,sizeUB:150"`
}

type QosFlowNotifyItem struct {
	QosFlowIdentifier            QosFlowIdentifier                      `bitstring:"sizeLB:0,sizeUB:150"`
	NotificationCause            []byte                                 `bitstring:"sizeLB:0,sizeUB:150"`
	CurrentQosParametersSetIndex AlternativeQosParametersSetNotifyIndex `bitstring:"sizeLB:0,sizeUB:150"`
}

type QosFlowFeedbackItem struct {
	QosFlowIdentifier           QosFlowIdentifier         `bitstring:"sizeLB:0,sizeUB:150"`
	UpdateFeedback              aper.BitString            `bitstring:"sizeLB:0,sizeUB:150"`
	CnPacketDelayBudgetDownlink ExtendedPacketDelayBudget `bitstring:"sizeLB:0,sizeUB:150"`
	CnPacketDelayBudgetUplink   ExtendedPacketDelayBudget `bitstring:"sizeLB:0,sizeUB:150"`
}
