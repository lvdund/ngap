package ie

import "ngap/aper"

type LastVisitedCellInformation struct {
	ChoiceLastVisitedCellInformation ChoiceLastVisitedCellInformation //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceLastVisitedCellInformation struct {
	NgRanCell  NgRanCell  //`bitstring:"sizeLB:0,sizeUB:150"`
	EUtranCell EUtranCell //`bitstring:"sizeLB:0,sizeUB:150"`
	UtranCell  UtranCell  //`bitstring:"sizeLB:0,sizeUB:150"`
	GeranCell  GeranCell  //`bitstring:"sizeLB:0,sizeUB:150"`
}

type NgRanCell struct {
	LastVisitedNgRanCellInformation LastVisitedNgRanCellInformation //`bitstring:"sizeLB:0,sizeUB:150"`
}

type EUtranCell struct {
	LastVisitedEUtranCellInformation aper.OctetString //`octetstring:"sizeLB:0,sizeUB:150"`
}

type UtranCell struct {
	LastVisitedUtranCellInformation aper.OctetString //`octetstring:"sizeLB:0,sizeUB:150"`
}

type GeranCell struct {
	LastVisitedGeranCellInformation aper.OctetString //`octetstring:"sizeLB:0,sizeUB:150"`
}
