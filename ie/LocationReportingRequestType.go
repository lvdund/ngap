package ie

type LocationReportingRequestType struct {
	EventType                                 []byte                       //`bitstring:"sizeLB:0,sizeUB:150"`
	ReportArea                                []byte                       //`bitstring:"sizeLB:0,sizeUB:150"`
	AreaOfInterestList                        []AreaOfInterestItem         //`bitstring:"sizeLB:0,sizeUB:150"`
	LocationReportingReferenceIdToBeCancelled LocationReportingReferenceId //`bitstring:"sizeLB:0,sizeUB:150"`
	AdditionalLocationInformation             []byte                       //`bitstring:"sizeLB:0,sizeUB:150"`
}

type AreaOfInterestItem struct {
	AreaOfInterest               AreaOfInterest               //`bitstring:"sizeLB:0,sizeUB:150"`
	LocationReportingReferenceId LocationReportingReferenceId //`bitstring:"sizeLB:0,sizeUB:150"`
}
