package utils

import (
	"github.com/lvdund/ngap/aper"
	"github.com/lvdund/ngap/ies"
)

type RatType string

// TS 38.413 9.3.1.85
func RATRestrictionInformationToNgap(ratType RatType) (ratResInfo ies.RATRestrictionInformation) {
	// Only support EUTRA & NR in version15.2.0
	switch ratType {
	case "EUTRA":
		ratResInfo.Value = aper.BitString{
			Bytes:   []byte{0x80},
			NumBits: 8,
		}
	case "NR":
		ratResInfo.Value = aper.BitString{
			Bytes:   []byte{0x40},
			NumBits: 8,
		}
	}
	return
}
