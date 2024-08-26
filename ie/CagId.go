package ie

import "ngap/aper"

type CagId struct {
	CagId aper.BitString `bitstring:"sizeLB:32,sizeUB:32"`
}
