package ie

type NpnPagingAssistanceInformation struct {
	ChoiceNpnPagingAssistanceInformation *ChoiceNpnPagingAssistanceInformation
}

type ChoiceNpnPagingAssistanceInformation struct {
	PniNpnPagingAssistance *PniNpnPagingAssistance
}

type PniNpnPagingAssistance struct {
	PniNpnPagingAssistance *AllowedPniNpnList
}
