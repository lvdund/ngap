package ie

import "ngap/aper"

type TimeSynchronisationAssistanceInformation struct {
	TimeDistributionIndication       []byte       //`bitstring:"sizeLB:0,sizeUB:150"`
	UuTimeSynchronisationErrorBudget aper.Integer //`Integer:"valueLB:1,valueUB:1000000"`
}
