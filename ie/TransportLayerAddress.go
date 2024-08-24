package ie

import "ngap/aper"

type TransportLayerAddress struct {
	TransportLayerAddress aper.BitString //`bitstring:"sizeLB:1,sizeUB:160"`
}
