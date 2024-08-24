package ie

import "ngap/aper"

type UlCpSecurityInformation struct {
	UlNasMac   aper.BitString //`bitstring:"sizeLB:16,sizeUB:16"`
	UlNasCount aper.BitString //`bitstring:"sizeLB:5,sizeUB:5"`
}
