package ie

import "ngap/aper"

type SonInformationReport struct {
	ChoiceSonInformationReport ChoiceSonInformationReport `bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceSonInformationReport struct {
	FailureIndicationInformation  FailureIndicationInformation  `bitstring:"sizeLB:0,sizeUB:150"`
	HoReportInformation           HoReportInformation           `bitstring:"sizeLB:0,sizeUB:150"`
	SuccessfulHoReportInformation SuccessfulHoReportInformation `bitstring:"sizeLB:0,sizeUB:150"`
}

type FailureIndicationInformation struct {
	FailureIndication FailureIndication `bitstring:"sizeLB:0,sizeUB:150"`
}

type HoReportInformation struct {
	HoReport HoReport `bitstring:"sizeLB:0,sizeUB:150"`
}

type SuccessfulHoReportInformation struct {
	SuccessfulHoReportList []SuccessfulHoReportItem `bitstring:"sizeLB:0,sizeUB:150"`
}

type SuccessfulHoReportItem struct {
	SuccessfulHoReportContainer aper.OctetString `octetstring:"sizeLB:0,sizeUB:150"`
}
