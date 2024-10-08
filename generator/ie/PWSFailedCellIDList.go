package ie

type PWSFailedCellIDList struct {
	EUTRACGIPWSFailedList EUTRACGIList              `False,`
	NRCGIPWSFailedList    NRCGIList                 `False,`
	ChoiceExtensions      PWSFailedCellIDListExtIEs `False,`
}
