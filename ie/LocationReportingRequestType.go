package ie

type LocationReportingRequestType struct {
EventType	*[]byte
ReportArea	*[]byte
AreaOfInterestList	*[]AreaOfInterestItem
LocationReportingReferenceIdToBeCancelled	*LocationReportingReferenceId
AdditionalLocationInformation	*[]byte
}

type AreaOfInterestItem struct {
AreaOfInterest	*AreaOfInterest
LocationReportingReferenceId	*LocationReportingReferenceId
}
