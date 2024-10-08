package ie

type QoSFlowsUsageReportItem struct {
	QosFlowIdentifier       QosFlowIdentifier             `False,`
	QoSFlowsTimedReportList VolumeTimedReportList         `False,`
	IEExtensions            QoSFlowsUsageReportItemExtIEs `False,OPTIONAL`
}
