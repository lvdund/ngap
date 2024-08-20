package ie

type InterSystemResourceStatusReport struct {
ChoiceReportingSystem	*ChoiceReportingSystem
}

type ChoiceReportingSystem struct {
EUtran	*EUtran
NgRan	*NgRan
}

type EUtran struct {
EUtranCellReportList	*[]EUtranCellReportItem
}

type EUtranCellReportItem struct {
CellId	*EUtraCgi
EUtranCompositeAvailableCapacityGroup	*EUtranCompositeAvailableCapacityGroup
NumberOfActiveUes	uint32	//`bitstring:"sizeLB:0,sizeUB:16777215"`
RrcConnections	uint32	//`bitstring:"sizeLB:1,sizeUB:65536"`
EUtranRadioResourceStatus	*EUtranRadioResourceStatus
}

type NgRan struct {
NgRanCellReportList	*[]NgRanCellReportItem
}

type NgRanCellReportItem struct {
CellId	*NgRanCgi
NrCompositeAvailableCapacityGroup	*EUtranCompositeAvailableCapacityGroup
NumberOfActiveUes	uint32	//`bitstring:"sizeLB:0,sizeUB:16777215"`
RrcConnections	uint32	//`bitstring:"sizeLB:1,sizeUB:65536"`
NrRadioResourceStatus	*NrRadioResourceStatus
}
