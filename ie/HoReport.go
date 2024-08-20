package ie

type HoReport struct {
	HandoverReportType          *[]byte
	HandoverCause               *Cause
	SourceCellCgi               *NgRanCgi
	TargetCellCgi               *NgRanCgi
	ReEstablishmentCellCgi      *NgRanCgi
	SourceCellCRnti             *[]byte
	TargetCellInEUtran          *EUtraCgi
	MobilityInformation         *[]byte
	UeRlfReportContainer        *UeRlfReportContainer
	ExtendedMobilityInformation *[]byte
}
