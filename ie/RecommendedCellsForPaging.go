package ie

import "ngap/aper"

type RecommendedCellsForPaging struct {
RecommendedCellList	[]RecommendedCellItem	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type RecommendedCellItem struct {
NgRanCgi	NgRanCgi	//`bitstring:"sizeLB:0,sizeUB:150"`
TimeStayedInCell	aper.Integer	//`Integer:"valueLB:0,valueUB:4095"`
}
