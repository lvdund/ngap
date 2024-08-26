package ie

import "ngap/aper"

type PlmnIdentity struct {
	PlmnIdentity aper.OctetString `octetstring:"sizeLB:3,sizeUB:3"`
}
