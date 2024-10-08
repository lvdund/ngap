package ie

type CancelledCellsInTAINRItem struct {
	NRCGI              NRCGI                           `True,`
	NumberOfBroadcasts NumberOfBroadcasts              `False,`
	IEExtensions       CancelledCellsInTAINRItemExtIEs `False,OPTIONAL`
}
