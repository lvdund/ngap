package ie

type LocationReportingRequestType struct {
	EventType                                 EventType                          `False,`
	ReportArea                                ReportArea                         `False,`
	AreaOfInterestList                        AreaOfInterestList                 `False,OPTIONAL`
	LocationReportingReferenceIDToBeCancelled LocationReportingReferenceID       `False,OPTIONAL`
	IEExtensions                              LocationReportingRequestTypeExtIEs `False,OPTIONAL`
}
