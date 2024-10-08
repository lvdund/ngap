package ie

import "ngap/aper"

type Nid struct {
	Nid aper.BitString `bitstring:"sizeLB:44,sizeUB:44"`
}
