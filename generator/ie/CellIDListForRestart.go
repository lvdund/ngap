package ie

type CellIDListForRestart struct {
	EUTRACGIListforRestart EUTRACGIList               `False,`
	NRCGIListforRestart    NRCGIList                  `False,`
	ChoiceExtensions       CellIDListForRestartExtIEs `False,`
}
