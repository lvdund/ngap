package ie

type TAIBroadcastEUTRAItem struct {
	TAI                      TAI                         `True,`
	CompletedCellsInTAIEUTRA CompletedCellsInTAIEUTRA    `False,`
	IEExtensions             TAIBroadcastEUTRAItemExtIEs `False,OPTIONAL`
}
