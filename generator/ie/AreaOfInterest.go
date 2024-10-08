package ie

type AreaOfInterest struct {
	AreaOfInterestTAIList     AreaOfInterestTAIList     `False,OPTIONAL`
	AreaOfInterestCellList    AreaOfInterestCellList    `False,OPTIONAL`
	AreaOfInterestRANNodeList AreaOfInterestRANNodeList `False,OPTIONAL`
	IEExtensions              AreaOfInterestExtIEs      `False,OPTIONAL`
}
