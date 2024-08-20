package ie

type NpnAccessInformation struct {
	ChoiceNpnAccessInformation *ChoiceNpnAccessInformation
}

type ChoiceNpnAccessInformation struct {
	PniNpnAccessInformation *PniNpnAccessInformation
}

type PniNpnAccessInformation struct {
	CellCagList *CellCagList
}
