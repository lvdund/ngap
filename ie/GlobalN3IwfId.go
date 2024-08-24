package ie

import "ngap/aper"

type GlobalN3IwfId struct {
PlmnIdentity	PlmnIdentity	//`bitstring:"sizeLB:0,sizeUB:150"`
ChoiceN3IwfId	ChoiceN3IwfId	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceN3IwfId struct {
N3IwfId	N3IwfId	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type N3IwfId struct {
N3IwfId	aper.BitString	//`bitstring:"sizeLB:16,sizeUB:16"`
}
