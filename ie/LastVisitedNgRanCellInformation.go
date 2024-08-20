package ie

type LastVisitedNgRanCellInformation struct {
	GlobalCellId                          *NgRanCgi
	CellType                              *CellType
	TimeUeStayedInCell                    uint16 //`bitstring:"sizeLB:0,sizeUB:4095"`
	TimeUeStayedInCellEnhancedGranularity uint16 //`bitstring:"sizeLB:0,sizeUB:40950"`
	HoCauseValue                          *Cause
	LastVisitedPscellList                 *LastVisitedPscellList
}

type LastVisitedPscellList struct {
	LastVisitedPscellInformation *LastVisitedPscellInformation
}
