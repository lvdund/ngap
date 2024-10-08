package ie

import "ngap/aper"

type PagingAssistanceDataForCeCapableUe struct {
	GlobalCellId             EUtraCgi         `bitstring:"sizeLB:0,sizeUB:150"`
	CoverageEnhancementLevel aper.OctetString `octetstring:"sizeLB:0,sizeUB:150"`
}
