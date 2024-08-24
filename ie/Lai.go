package ie

import "ngap/aper"

type Lai struct {
	PlmnIdentity PlmnIdentity     //`bitstring:"sizeLB:0,sizeUB:150"`
	Lac          aper.OctetString //`octetstring:"sizeLB:2,sizeUB:2"`
}
