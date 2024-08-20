package ie

type SecondaryRatUsageInformation struct {
PduSessionUsageReport	*PduSessionUsageReport
QosFlowsUsageReportList	*[]QosFlowUsageReportItem
}

type PduSessionUsageReport struct {
RatType	*[]byte
PduSessionTimedReportList	*VolumeTimedReportList
}

type QosFlowUsageReportItem struct {
QosFlowIndicator	*QosFlowIdentifier
RatType	*[]byte
QosFlowsTimedReportList	*VolumeTimedReportList
}
