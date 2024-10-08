package ie

type UeHistoryInformationFromUe struct {
	ChoiceUeHistoryInformationFromUe ChoiceUeHistoryInformationFromUe `bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceUeHistoryInformationFromUe struct {
	Nr Nr `bitstring:"sizeLB:0,sizeUB:150"`
}
