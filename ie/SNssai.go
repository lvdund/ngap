package ie

import (
	"fmt"
	"ngap/aper"
)

type SNssai struct {
	Sst aper.OctetString `octetstring:"sizeLB:1,sizeUB:1"`
	Sd  aper.OctetString `octetstring:"sizeLB:3,sizeUB:3"`
	//extention
}

func (ie *SNssai) Decode(r *aper.AperReader) (err error) {
	var exBit bool
	if exBit, err = r.ReadBool(); err != nil {
		return
	}
	fmt.Printf("exBit<SNssai>=%v\n", exBit)
	optionals, _ := r.ReadBits(2) //optional bits
	fmt.Printf("optionals<SNssai>=%.8b\n", optionals)
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
	fmt.Printf("SNssai: %.8b-%.8b[%d]\n", ie.Sst, ie.Sd, len(ie.Sst))
	return
}
