package ie

import "ngap/aper"

type NrCgi struct {
PlmnIdentity	PlmnIdentity	//`bitstring:"sizeLB:0,sizeUB:150"`
NrCellIdentity	aper.BitString	//`bitstring:"sizeLB:36,sizeUB:36"`
}
