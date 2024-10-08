package ie

type AreaOfInterestItem struct {
	AreaOfInterest               AreaOfInterest               `True,`
	LocationReportingReferenceID LocationReportingReferenceID `False,`
	IEExtensions                 AreaOfInterestItemExtIEs     `False,OPTIONAL`
}
