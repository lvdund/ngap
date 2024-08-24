package ie

type UePresenceInAreaOfInterestList struct {
UePresenceInAreaOfInterestItem	UePresenceInAreaOfInterestItem	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type UePresenceInAreaOfInterestItem struct {
LocationReportingReferenceId	LocationReportingReferenceId	//`bitstring:"sizeLB:0,sizeUB:150"`
UePresence	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
}
