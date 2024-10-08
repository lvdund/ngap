package ie

type TAIBroadcastNRItem struct {
	TAI                   TAI                      `True,`
	CompletedCellsInTAINR CompletedCellsInTAINR    `False,`
	IEExtensions          TAIBroadcastNRItemExtIEs `False,OPTIONAL`
}
