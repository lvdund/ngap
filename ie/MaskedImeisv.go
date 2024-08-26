package ie

import "ngap/aper"

type MaskedImeisv struct {
	MaskedImeisv aper.BitString `bitstring:"sizeLB:64,sizeUB:64"`
}
