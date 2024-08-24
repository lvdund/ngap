package ie

import "ngap/aper"

type GlobalTwifId struct {
	PlmnIdentity PlmnIdentity //`bitstring:"sizeLB:0,sizeUB:150"`
	ChoiceTwifId ChoiceTwifId //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceTwifId struct {
	TwifId TwifId //`bitstring:"sizeLB:0,sizeUB:150"`
}

type TwifId struct {
	TwifId aper.BitString //`bitstring:"sizeLB:0,sizeUB:150"`
}
