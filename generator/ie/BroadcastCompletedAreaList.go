package ie

type BroadcastCompletedAreaList struct {
	CellIDBroadcastEUTRA          CellIDBroadcastEUTRA             `False,`
	TAIBroadcastEUTRA             TAIBroadcastEUTRA                `False,`
	EmergencyAreaIDBroadcastEUTRA EmergencyAreaIDBroadcastEUTRA    `False,`
	CellIDBroadcastNR             CellIDBroadcastNR                `False,`
	TAIBroadcastNR                TAIBroadcastNR                   `False,`
	EmergencyAreaIDBroadcastNR    EmergencyAreaIDBroadcastNR       `False,`
	ChoiceExtensions              BroadcastCompletedAreaListExtIEs `False,`
}
