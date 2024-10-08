package ie

type SecondaryRATUsageInformation struct {
	PDUSessionUsageReport   PDUSessionUsageReport              `True,OPTIONAL`
	QosFlowsUsageReportList QoSFlowsUsageReportList            `False,OPTIONAL`
	IEExtension             SecondaryRATUsageInformationExtIEs `False,OPTIONAL`
}
