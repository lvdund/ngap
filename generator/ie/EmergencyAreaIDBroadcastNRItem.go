package ie

type EmergencyAreaIDBroadcastNRItem struct {
	EmergencyAreaID       EmergencyAreaID                      `False,`
	CompletedCellsInEAINR CompletedCellsInEAINR                `False,`
	IEExtensions          EmergencyAreaIDBroadcastNRItemExtIEs `False,OPTIONAL`
}
