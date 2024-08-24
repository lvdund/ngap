package ie

import "ngap/aper"

type InterSystemHoReport struct {
	ChoiceHandoverReportType ChoiceHandoverReportType //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceHandoverReportType struct {
	TooEarlyInterSystemHo    TooEarlyInterSystemHo    //`bitstring:"sizeLB:0,sizeUB:150"`
	InterSystemUnnecessaryHo InterSystemUnnecessaryHo //`bitstring:"sizeLB:0,sizeUB:150"`
}

type TooEarlyInterSystemHo struct {
	SourceCellId         EUtraCgi             //`bitstring:"sizeLB:0,sizeUB:150"`
	FailureCellId        NgRanCgi             //`bitstring:"sizeLB:0,sizeUB:150"`
	UeRlfReportContainer UeRlfReportContainer //`bitstring:"sizeLB:0,sizeUB:150"`
}

type InterSystemUnnecessaryHo struct {
	SourceCellCgi     NgRanCgi            //`bitstring:"sizeLB:0,sizeUB:150"`
	TargetCellCgi     EUtraCgi            //`bitstring:"sizeLB:0,sizeUB:150"`
	EarlyIratHo       []byte              //`bitstring:"sizeLB:0,sizeUB:150"`
	CandidateCellList []CandidateCellItem //`bitstring:"sizeLB:0,sizeUB:150"`
}

type CandidateCellItem struct {
	ChoiceCandidateCellType ChoiceCandidateCellType //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceCandidateCellType struct {
	CandidateCgi CandidateCgi //`bitstring:"sizeLB:0,sizeUB:150"`
	CandidatePci CandidatePci //`bitstring:"sizeLB:0,sizeUB:150"`
}

type CandidateCgi struct {
	CandidateCellId NrCgi //`bitstring:"sizeLB:0,sizeUB:150"`
}

type CandidatePci struct {
	CandidatePci     aper.Integer //`Integer:"valueLB:0,valueUB:1007"`
	CandidateNrArfcn aper.Integer //`Integer:"valueLB:0,valueUB:150"`
}
