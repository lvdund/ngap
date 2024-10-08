package ie

type TAICancelledNRItem struct {
	TAI                   TAI                      `True,`
	CancelledCellsInTAINR CancelledCellsInTAINR    `False,`
	IEExtensions          TAICancelledNRItemExtIEs `False,OPTIONAL`
}
