package ie

type UEPresenceInAreaOfInterestItem struct {
	LocationReportingReferenceID LocationReportingReferenceID         `False,`
	UEPresence                   UEPresence                           `False,`
	IEExtensions                 UEPresenceInAreaOfInterestItemExtIEs `False,OPTIONAL`
}
