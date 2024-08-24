package ie

import "ngap/aper"

type InterSystemResourceStatusRequest struct {
	ReportingSystem       ReportingSystem  //`bitstring:"sizeLB:0,sizeUB:150"`
	ReportCharacteristics []byte           //`bitstring:"sizeLB:32,sizeUB:32"`
	ChoiceReportType      ChoiceReportType //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceReportType struct {
	EventBasedReporting EventBasedReporting //`bitstring:"sizeLB:0,sizeUB:150"`
	PeriodicReporting   PeriodicReporting   //`bitstring:"sizeLB:0,sizeUB:150"`
}

type EventBasedReporting struct {
	InterSystemResourceThresholdLow    aper.Integer //`Integer:"valueLB:0,valueUB:100"`
	InterSystemResourceThresholdHigh   aper.Integer //`Integer:"valueLB:0,valueUB:100"`
	NumberOfMeasurementReportingLevels []byte       //`bitstring:"sizeLB:0,sizeUB:150"`
}

type PeriodicReporting struct {
	ReportingPeriodicity []byte //`bitstring:"sizeLB:0,sizeUB:150"`
}
