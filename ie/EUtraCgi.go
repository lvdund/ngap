package ie

import "ngap/aper"

type EUtraCgi struct {
PlmnIdentity	PlmnIdentity	//`bitstring:"sizeLB:0,sizeUB:150"`
EUtraCellIdentity	aper.BitString	//`bitstring:"sizeLB:28,sizeUB:28"`
}
