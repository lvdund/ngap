package ie

type AreaOfInterest struct {
	AreaOfInterestTaiList     []AreaOfInterestTaiItem     //`bitstring:"sizeLB:0,sizeUB:150"`
	AreaOfInterestCellList    []AreaOfInterestCellItem    //`bitstring:"sizeLB:0,sizeUB:150"`
	AreaOfInterestRanNodeList []AreaOfInterestRanNodeItem //`bitstring:"sizeLB:0,sizeUB:150"`
}

type AreaOfInterestTaiItem struct {
	Tai Tai //`bitstring:"sizeLB:0,sizeUB:150"`
}

type AreaOfInterestCellItem struct {
	NgRanCgi NgRanCgi //`bitstring:"sizeLB:0,sizeUB:150"`
}

type AreaOfInterestRanNodeItem struct {
	GlobalRanNodeId GlobalRanNodeId //`bitstring:"sizeLB:0,sizeUB:150"`
}
