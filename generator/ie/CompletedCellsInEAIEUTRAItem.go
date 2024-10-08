package ie

type CompletedCellsInEAIEUTRAItem struct {
	EUTRACGI     EUTRACGI                           `True,`
	IEExtensions CompletedCellsInEAIEUTRAItemExtIEs `False,OPTIONAL`
}
