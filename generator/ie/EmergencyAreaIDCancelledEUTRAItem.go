package ie

type EmergencyAreaIDCancelledEUTRAItem struct {
	EmergencyAreaID          EmergencyAreaID                         `False,`
	CancelledCellsInEAIEUTRA CancelledCellsInEAIEUTRA                `False,`
	IEExtensions             EmergencyAreaIDCancelledEUTRAItemExtIEs `False,OPTIONAL`
}
