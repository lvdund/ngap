package ie

import "ngap/aper"

type InterSystemCellActivationRequest struct {
	ActivationId        aper.Integer        //`Integer:"valueLB:0,valueUB:16384"`
	CellsToActivateList CellsToActivateList //`bitstring:"sizeLB:0,sizeUB:150"`
}

type CellsToActivateList struct {
	NgRanCgi NgRanCgi //`bitstring:"sizeLB:0,sizeUB:150"`
}
