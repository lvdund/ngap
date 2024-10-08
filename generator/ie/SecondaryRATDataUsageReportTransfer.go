package ie

type SecondaryRATDataUsageReportTransfer struct {
	SecondaryRATUsageInformation SecondaryRATUsageInformation              `False,OPTIONAL`
	IEExtensions                 SecondaryRATDataUsageReportTransferExtIEs `False,OPTIONAL`
}
