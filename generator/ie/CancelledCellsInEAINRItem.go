package ie

type CancelledCellsInEAINRItem struct {
	NRCGI              NRCGI                           `True,`
	NumberOfBroadcasts NumberOfBroadcasts              `False,`
	IEExtensions       CancelledCellsInEAINRItemExtIEs `False,OPTIONAL`
}
