package ie

import "ngap/aper"

type GlobalGnbId struct {
	PlmnIdentity PlmnIdentity `bitstring:"sizeLB:0,sizeUB:150"`
	ChoiceGnbId  ChoiceGnbId  `bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceGnbId struct {
	GnbId GnbId `bitstring:"sizeLB:0,sizeUB:150"`
}

type GnbId struct {
	GnbId aper.BitString `bitstring:"sizeLB:22,sizeUB:32"`
}
