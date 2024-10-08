package ie

type CancelledCellsInEAIEUTRAItem struct {
	EUTRACGI           EUTRACGI                           `True,`
	NumberOfBroadcasts NumberOfBroadcasts                 `False,`
	IEExtensions       CancelledCellsInEAIEUTRAItemExtIEs `False,OPTIONAL`
}
