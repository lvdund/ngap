package ie

import "ngap/aper"

type LastVisitedNgRanCellInformation struct {
	GlobalCellId                          NgRanCgi              //`bitstring:"sizeLB:0,sizeUB:150"`
	CellType                              CellType              //`bitstring:"sizeLB:0,sizeUB:150"`
	TimeUeStayedInCell                    aper.Integer          //`Integer:"valueLB:0,valueUB:4095"`
	TimeUeStayedInCellEnhancedGranularity aper.Integer          //`Integer:"valueLB:0,valueUB:40950"`
	HoCauseValue                          Cause                 //`bitstring:"sizeLB:0,sizeUB:150"`
	LastVisitedPscellList                 LastVisitedPscellList //`bitstring:"sizeLB:0,sizeUB:150"`
}

type LastVisitedPscellList struct {
	LastVisitedPscellInformation LastVisitedPscellInformation //`bitstring:"sizeLB:0,sizeUB:150"`
}
