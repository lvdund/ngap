package ie

type NrFrequencyInfo struct {
	NrArfcn             *int
	NrFrequencyBandList *[]NrFrequencyBandItem
}

type NrFrequencyBandItem struct {
	NrFrequencyBand *int
}
