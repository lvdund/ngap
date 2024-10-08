package ie

type BroadcastCancelledAreaList struct {
	CellIDCancelledEUTRA          CellIDCancelledEUTRA             `False,`
	TAICancelledEUTRA             TAICancelledEUTRA                `False,`
	EmergencyAreaIDCancelledEUTRA EmergencyAreaIDCancelledEUTRA    `False,`
	CellIDCancelledNR             CellIDCancelledNR                `False,`
	TAICancelledNR                TAICancelledNR                   `False,`
	EmergencyAreaIDCancelledNR    EmergencyAreaIDCancelledNR       `False,`
	ChoiceExtensions              BroadcastCancelledAreaListExtIEs `False,`
}
