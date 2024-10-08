package ie

type CancelledCellsInTAIEUTRAItem struct {
	EUTRACGI           EUTRACGI                           `True,`
	NumberOfBroadcasts NumberOfBroadcasts                 `False,`
	IEExtensions       CancelledCellsInTAIEUTRAItemExtIEs `False,OPTIONAL`
}
