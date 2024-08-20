package ie

type NpnMobilityInformation struct {
ChoiceNpnMobilityInformation	*ChoiceNpnMobilityInformation
}

type ChoiceNpnMobilityInformation struct {
SnpnMobilityInformation	*SnpnMobilityInformation
PniNpnMobilityInformation	*PniNpnMobilityInformation
}

type SnpnMobilityInformation struct {
ServingNid	*Nid
}

type PniNpnMobilityInformation struct {
AllowedPniNpnList	*AllowedPniNpnList
}
