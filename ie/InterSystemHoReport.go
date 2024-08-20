package ie

type InterSystemHoReport struct {
ChoiceHandoverReportType	*ChoiceHandoverReportType
}

type ChoiceHandoverReportType struct {
TooEarlyInterSystemHo	*TooEarlyInterSystemHo
InterSystemUnnecessaryHo	*InterSystemUnnecessaryHo
}

type TooEarlyInterSystemHo struct {
SourceCellId	*EUtraCgi
FailureCellId	*NgRanCgi
UeRlfReportContainer	*UeRlfReportContainer
}

type InterSystemUnnecessaryHo struct {
SourceCellCgi	*NgRanCgi
TargetCellCgi	*EUtraCgi
EarlyIratHo	*[]byte
CandidateCellList	*[]CandidateCellItem
}

type CandidateCellItem struct {
ChoiceCandidateCellType	*ChoiceCandidateCellType
}

type ChoiceCandidateCellType struct {
CandidateCgi	*CandidateCgi
CandidatePci	*CandidatePci
}

type CandidateCgi struct {
CandidateCellId	*NrCgi
}

type CandidatePci struct {
CandidatePci	uint16	//`bitstring:"sizeLB:0,sizeUB:1007"`
CandidateNrArfcn	*int
}
