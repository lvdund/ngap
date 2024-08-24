package ie

type SecondaryRatUsageInformation struct {
	PduSessionUsageReport   PduSessionUsageReport    //`bitstring:"sizeLB:0,sizeUB:150"`
	QosFlowsUsageReportList []QosFlowUsageReportItem //`bitstring:"sizeLB:0,sizeUB:150"`
}

type PduSessionUsageReport struct {
	RatType                   []byte                //`bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionTimedReportList VolumeTimedReportList //`bitstring:"sizeLB:0,sizeUB:150"`
}

type QosFlowUsageReportItem struct {
	QosFlowIndicator        QosFlowIdentifier     //`bitstring:"sizeLB:0,sizeUB:150"`
	RatType                 []byte                //`bitstring:"sizeLB:0,sizeUB:150"`
	QosFlowsTimedReportList VolumeTimedReportList //`bitstring:"sizeLB:0,sizeUB:150"`
}
