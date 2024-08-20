package ie

type AreaOfInterest struct {
	AreaOfInterestTaiList     *[]AreaOfInterestTaiItem
	AreaOfInterestCellList    *[]AreaOfInterestCellItem
	AreaOfInterestRanNodeList *[]AreaOfInterestRanNodeItem
}

type AreaOfInterestTaiItem struct {
	Tai *Tai
}

type AreaOfInterestCellItem struct {
	NgRanCgi *NgRanCgi
}

type AreaOfInterestRanNodeItem struct {
	GlobalRanNodeId *GlobalRanNodeId
}
