package ie

type CellIDCancelledEUTRAItem struct {
	EUTRACGI           EUTRACGI                       `True,`
	NumberOfBroadcasts NumberOfBroadcasts             `False,`
	IEExtensions       CellIDCancelledEUTRAItemExtIEs `False,OPTIONAL`
}
