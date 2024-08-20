package ie

type UePresenceInAreaOfInterestList struct {
UePresenceInAreaOfInterestItem	*UePresenceInAreaOfInterestItem
}

type UePresenceInAreaOfInterestItem struct {
LocationReportingReferenceId	*LocationReportingReferenceId
UePresence	*[]byte
}
