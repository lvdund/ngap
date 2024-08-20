package ie

type RecommendedCellsForPaging struct {
RecommendedCellList	*[]RecommendedCellItem
}

type RecommendedCellItem struct {
NgRanCgi	*NgRanCgi
TimeStayedInCell	uint16	//`bitstring:"sizeLB:0,sizeUB:4095"`
}
