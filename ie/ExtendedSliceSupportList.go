package ie

import (
	"ngap/aper"
)

type ExtendedSliceSupportList struct {
	SliceSupportItem SliceSupportItem `bitstring:"sizeLB:0,sizeUB:150"`
}

type SliceSupportItem struct {
	SNssai SNssai
}

func (ie *SliceSupportItem) Decode(r *aper.AperReader) (err error) {
	return ie.SNssai.Decode(r)
}
