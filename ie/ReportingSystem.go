package ie

type ReportingSystem struct {
	ChoiceReportingSystem *ChoiceReportingSystem
}

type EUtranCellToReportItem struct {
	CellId *EUtraCgi
}

type NgRanCellToReportItem struct {
	CellId *NgRanCgi
}
