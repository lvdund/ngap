package ie

type PDUSessionUsageReport struct {
	PDUSessionTimedReportList VolumeTimedReportList       `False,`
	IEExtensions              PDUSessionUsageReportExtIEs `False,OPTIONAL`
}
