package ie

type AreaScopeOfNeighbourCells struct {
	AreaScopeOfNeighbourCellsItem *AreaScopeOfNeighbourCellsItem
}

type AreaScopeOfNeighbourCellsItem struct {
	NrFrequencyInfo *NrFrequencyInfo
	PciListForMdt   *PciListForMdt
}

type PciListForMdt struct {
	NrPci uint16 //`bitstring:"sizeLB:0,sizeUB:1007"`
}
