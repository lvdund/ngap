package ie

type NpnMobilityInformation struct {
	ChoiceNpnMobilityInformation ChoiceNpnMobilityInformation //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceNpnMobilityInformation struct {
	SnpnMobilityInformation   SnpnMobilityInformation   //`bitstring:"sizeLB:0,sizeUB:150"`
	PniNpnMobilityInformation PniNpnMobilityInformation //`bitstring:"sizeLB:0,sizeUB:150"`
}

type SnpnMobilityInformation struct {
	ServingNid Nid //`bitstring:"sizeLB:0,sizeUB:150"`
}

type PniNpnMobilityInformation struct {
	AllowedPniNpnList AllowedPniNpnList //`bitstring:"sizeLB:0,sizeUB:150"`
}
