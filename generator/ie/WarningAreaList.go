package ie

type WarningAreaList struct {
	EUTRACGIListForWarning EUTRACGIListForWarning `False,`
	NRCGIListForWarning    NRCGIListForWarning    `False,`
	TAIListForWarning      TAIListForWarning      `False,`
	EmergencyAreaIDList    EmergencyAreaIDList    `False,`
	ChoiceExtensions       WarningAreaListExtIEs  `False,`
}
