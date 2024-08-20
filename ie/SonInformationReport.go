package ie

type SonInformationReport struct {
	ChoiceSonInformationReport *ChoiceSonInformationReport
}

type ChoiceSonInformationReport struct {
	FailureIndicationInformation  *FailureIndicationInformation
	HoReportInformation           *HoReportInformation
	SuccessfulHoReportInformation *SuccessfulHoReportInformation
}

type FailureIndicationInformation struct {
	FailureIndication *FailureIndication
}

type HoReportInformation struct {
	HoReport *HoReport
}

type SuccessfulHoReportInformation struct {
	SuccessfulHoReportList *[]SuccessfulHoReportItem
}

type SuccessfulHoReportItem struct {
	SuccessfulHoReportContainer *[]byte
}
