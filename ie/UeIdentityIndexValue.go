package ie

import "ngap/aper"

type UeIdentityIndexValue struct {
ChoiceUeIdentityIndexValue	ChoiceUeIdentityIndexValue	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceUeIdentityIndexValue struct {
IndexLength10	IndexLength10	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type IndexLength10 struct {
IndexLength10	aper.BitString	//`bitstring:"sizeLB:10,sizeUB:10"`
}
