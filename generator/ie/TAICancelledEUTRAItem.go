package ie

type TAICancelledEUTRAItem struct {
	TAI                      TAI                         `True,`
	CancelledCellsInTAIEUTRA CancelledCellsInTAIEUTRA    `False,`
	IEExtensions             TAICancelledEUTRAItemExtIEs `False,OPTIONAL`
}
