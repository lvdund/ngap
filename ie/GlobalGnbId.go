package ie

import (
	"fmt"
	"ngap/aper"
)

type GlobalGnbId struct {
	PlmnIdentity PlmnIdentity
	ChoiceGnbId  ChoiceGnbId
}

func (ie *GlobalGnbId) Decode(r *aper.AperReader) (err error) {
	var exBit bool
	if exBit, err = r.ReadBool(); err != nil {
		return
	}
	_ = exBit
	if ie.PlmnIdentity.Decode(r); err != nil {
		return
	}
	var choice uint64
	if choice, err = r.ReadChoice(1, false); err != nil {
		return
	}
	fmt.Printf("Gnb ID choice = %d\n", choice)
	var gnbId []byte
	var nbits uint
	if gnbId, nbits, err = r.ReadBitString(&aper.Constraint{
		Lb: 22,
		Ub: 32,
	}, false); err != nil {
		return
	}
	fmt.Printf("gnb id= %.8b[%d]\n", gnbId, nbits)
	return
}

type ChoiceGnbId struct {
	Choice uint64
	GnbId  GnbId
}

type GnbId struct {
	GnbId aper.BitString `bitstring:"sizeLB:22,sizeUB:32"`
}
