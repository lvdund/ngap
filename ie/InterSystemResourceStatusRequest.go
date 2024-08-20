package ie

type InterSystemResourceStatusRequest struct {
	ReportingSystem       *ReportingSystem
	ReportCharacteristics []byte //`bitstring:"sizeLB:32,sizeUB:32"`
	ChoiceReportType      *ChoiceReportType
}

type ChoiceReportType struct {
	EventBasedReporting *EventBasedReporting
	PeriodicReporting   *PeriodicReporting
}

type EventBasedReporting struct {
	InterSystemResourceThresholdLow    uint8 //`bitstring:"sizeLB:0,sizeUB:100"`
	InterSystemResourceThresholdHigh   uint8 //`bitstring:"sizeLB:0,sizeUB:100"`
	NumberOfMeasurementReportingLevels *[]byte
}

type PeriodicReporting struct {
	ReportingPeriodicity *[]byte
}
