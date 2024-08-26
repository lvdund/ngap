package ie

type NpnPagingAssistanceInformation struct {
	ChoiceNpnPagingAssistanceInformation ChoiceNpnPagingAssistanceInformation `bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceNpnPagingAssistanceInformation struct {
	PniNpnPagingAssistance PniNpnPagingAssistance `bitstring:"sizeLB:0,sizeUB:150"`
}

type PniNpnPagingAssistance struct {
	PniNpnPagingAssistance AllowedPniNpnList `bitstring:"sizeLB:0,sizeUB:150"`
}
