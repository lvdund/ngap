package ie

type InterSystemFailureIndication struct {
	UeRlfReportContainer UeRlfReportContainer `bitstring:"sizeLB:0,sizeUB:150"`
}
