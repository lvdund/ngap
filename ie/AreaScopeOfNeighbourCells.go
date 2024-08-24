package ie

import "ngap/aper"

type AreaScopeOfNeighbourCells struct {
AreaScopeOfNeighbourCellsItem	AreaScopeOfNeighbourCellsItem	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type AreaScopeOfNeighbourCellsItem struct {
NrFrequencyInfo	NrFrequencyInfo	//`bitstring:"sizeLB:0,sizeUB:150"`
PciListForMdt	PciListForMdt	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type PciListForMdt struct {
NrPci	aper.Integer	//`Integer:"valueLB:0,valueUB:1007"`
}
