package ie

import "ngap/aper"

type UeRlfReportContainer struct {
ChoiceRlfType	ChoiceRlfType	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceRlfType struct {
Nr	Nr	//`bitstring:"sizeLB:0,sizeUB:150"`
Lte	Lte	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type Lte struct {
LteUeRlfReportContainer	aper.OctetString	//`octetstring:"sizeLB:0,sizeUB:150"`
}
