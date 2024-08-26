package ie

type NpnAccessInformation struct {
	ChoiceNpnAccessInformation ChoiceNpnAccessInformation `bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceNpnAccessInformation struct {
	PniNpnAccessInformation PniNpnAccessInformation `bitstring:"sizeLB:0,sizeUB:150"`
}

type PniNpnAccessInformation struct {
	CellCagList CellCagList `bitstring:"sizeLB:0,sizeUB:150"`
}
