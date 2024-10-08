package ie

type EmergencyAreaIDCancelledNRItem struct {
	EmergencyAreaID       EmergencyAreaID                      `False,`
	CancelledCellsInEAINR CancelledCellsInEAINR                `False,`
	IEExtensions          EmergencyAreaIDCancelledNRItemExtIEs `False,OPTIONAL`
}
