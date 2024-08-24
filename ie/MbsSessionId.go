package ie

import "ngap/aper"

type MbsSessionId struct {
Tmgi	aper.OctetString	//`octetstring:"sizeLB:6,sizeUB:6"`
Nid	Nid	//`bitstring:"sizeLB:0,sizeUB:150"`
}
