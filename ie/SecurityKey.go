package ie

import "ngap/aper"

type SecurityKey struct {
	SecurityKey aper.BitString //`bitstring:"sizeLB:256,sizeUB:256"`
}
