package ie

type InterSystemCellActivationRequest struct {
	ActivationId        uint16 //`bitstring:"sizeLB:0,sizeUB:16384"`
	CellsToActivateList *CellsToActivateList
}

type CellsToActivateList struct {
	NgRanCgi *NgRanCgi
}
