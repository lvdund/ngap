package ie

import "ngap/aper"

type GlobalTngfId struct {
	PlmnIdentity PlmnIdentity `bitstring:"sizeLB:0,sizeUB:150"`
	ChoiceTngfId ChoiceTngfId `bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceTngfId struct {
	TngfId TngfId `bitstring:"sizeLB:0,sizeUB:150"`
}

type TngfId struct {
	TngfId aper.BitString `bitstring:"sizeLB:0,sizeUB:150"`
}
