package ie

import "ngap/aper"

type InterSystemResourceStatusReport struct {
	ChoiceReportingSystem ChoiceReportingSystem //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceReportingSystem struct {
	EUtran EUtran //`bitstring:"sizeLB:0,sizeUB:150"`
	NgRan  NgRan  //`bitstring:"sizeLB:0,sizeUB:150"`
}

type EUtran struct {
	EUtranCellReportList []EUtranCellReportItem //`bitstring:"sizeLB:0,sizeUB:150"`
}

type EUtranCellReportItem struct {
	CellId                                EUtraCgi                              //`bitstring:"sizeLB:0,sizeUB:150"`
	EUtranCompositeAvailableCapacityGroup EUtranCompositeAvailableCapacityGroup //`bitstring:"sizeLB:0,sizeUB:150"`
	NumberOfActiveUes                     aper.Integer                          //`Integer:"valueLB:0,valueUB:16777215"`
	RrcConnections                        aper.Integer                          //`Integer:"valueLB:1,valueUB:65536"`
	EUtranRadioResourceStatus             EUtranRadioResourceStatus             //`bitstring:"sizeLB:0,sizeUB:150"`
}

type NgRan struct {
	NgRanCellReportList []NgRanCellReportItem //`bitstring:"sizeLB:0,sizeUB:150"`
}

type NgRanCellReportItem struct {
	CellId                            NgRanCgi                              //`bitstring:"sizeLB:0,sizeUB:150"`
	NrCompositeAvailableCapacityGroup EUtranCompositeAvailableCapacityGroup //`bitstring:"sizeLB:0,sizeUB:150"`
	NumberOfActiveUes                 aper.Integer                          //`Integer:"valueLB:0,valueUB:16777215"`
	RrcConnections                    aper.Integer                          //`Integer:"valueLB:1,valueUB:65536"`
	NrRadioResourceStatus             NrRadioResourceStatus                 //`bitstring:"sizeLB:0,sizeUB:150"`
}
