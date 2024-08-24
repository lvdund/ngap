package ie

import "ngap/aper"

type SNssai struct {
	Sst aper.OctetString //`octetstring:"sizeLB:1,sizeUB:1"`
	Sd  aper.OctetString //`octetstring:"sizeLB:3,sizeUB:3"`
}
