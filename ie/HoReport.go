package ie

import "ngap/aper"

type HoReport struct {
HandoverReportType	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
HandoverCause	Cause	//`bitstring:"sizeLB:0,sizeUB:150"`
SourceCellCgi	NgRanCgi	//`bitstring:"sizeLB:0,sizeUB:150"`
TargetCellCgi	NgRanCgi	//`bitstring:"sizeLB:0,sizeUB:150"`
ReEstablishmentCellCgi	NgRanCgi	//`bitstring:"sizeLB:0,sizeUB:150"`
SourceCellCRnti	aper.BitString	//`bitstring:"sizeLB:0,sizeUB:150"`
TargetCellInEUtran	EUtraCgi	//`bitstring:"sizeLB:0,sizeUB:150"`
MobilityInformation	aper.BitString	//`bitstring:"sizeLB:0,sizeUB:150"`
UeRlfReportContainer	UeRlfReportContainer	//`bitstring:"sizeLB:0,sizeUB:150"`
ExtendedMobilityInformation	aper.BitString	//`bitstring:"sizeLB:0,sizeUB:150"`
}
