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
	r.ReadBool() //exBit
	r.ReadBool() //optional bit (1 bit)
	return ie.SNssai.Decode(r)
}
