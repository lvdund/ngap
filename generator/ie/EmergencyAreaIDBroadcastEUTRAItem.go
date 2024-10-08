package ie

type EmergencyAreaIDBroadcastEUTRAItem struct {
	EmergencyAreaID          EmergencyAreaID                         `False,`
	CompletedCellsInEAIEUTRA CompletedCellsInEAIEUTRA                `False,`
	IEExtensions             EmergencyAreaIDBroadcastEUTRAItemExtIEs `False,OPTIONAL`
}
