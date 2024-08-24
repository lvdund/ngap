package ie

type ReportingSystem struct {
ChoiceReportingSystem	ChoiceReportingSystem	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type EUtranCellToReportItem struct {
CellId	EUtraCgi	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type NgRanCellToReportItem struct {
CellId	NgRanCgi	//`bitstring:"sizeLB:0,sizeUB:150"`
}
