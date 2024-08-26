package ie

import "ngap/aper"

type MessageIdentifier struct {
	MessageIdentifier aper.BitString `bitstring:"sizeLB:16,sizeUB:16"`
}
