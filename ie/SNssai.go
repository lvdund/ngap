package ie

import (
	"fmt"
	"ngap/aper"
)

type SNssai struct {
	Sst aper.OctetString `octetstring:"sizeLB:1,sizeUB:1"`
	Sd  aper.OctetString `octetstring:"sizeLB:3,sizeUB:3"`
}

func (ie *SNssai) Decode(r *aper.AperReader) (err error) {
	var exBit bool
	if exBit, err = r.ReadBool(); err != nil {
		return
	}
	fmt.Printf("exBit=%v\n", exBit)
	if ie.Sst, err = r.ReadOctetString(&aper.Constraint{
		Lb: 1,
		Ub: 1,
	}, false); err != nil {
		return
	}
	if ie.Sd, err = r.ReadOctetString(&aper.Constraint{
		Lb: 3,
		Ub: 3,
	}, false); err != nil {
		return
	}
	fmt.Printf("SNssai: %x-%x\n", ie.Sst, ie.Sd)
	return
}
