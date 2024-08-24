package ie

import "ngap/aper"

type GlobalWAgfId struct {
PlmnIdentity	PlmnIdentity	//`bitstring:"sizeLB:0,sizeUB:150"`
ChoiceWAgfId	ChoiceWAgfId	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceWAgfId struct {
WAgfId	WAgfId	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type WAgfId struct {
WAgfId	aper.BitString	//`bitstring:"sizeLB:0,sizeUB:150"`
}
