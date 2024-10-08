package ie

type CellIDCancelledNRItem struct {
	NRCGI              NRCGI                       `True,`
	NumberOfBroadcasts NumberOfBroadcasts          `False,`
	IEExtensions       CellIDCancelledNRItemExtIEs `False,OPTIONAL`
}
