package ie

import (
	"ngap/aper"
	"fmt"
)

type Tac struct {
	Tac aper.OctetString `octetstring:"sizeLB:3,sizeUB:3"`
}

func (t *Tac) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteOctetString([]byte(t.Tac), &aper.Constraint{
		Lb: 3,
		Ub: 3,
	}, false)
	return
}

func (t *Tac) Decode(r *aper.AperReader) (err error) {
	var octets []byte
	if octets, err = r.ReadOctetString(&aper.Constraint{
		Lb: 3,
		Ub: 3,
	}, false); err != nil {
		return
	}
	fmt.Printf("TAC=%x\n", octets)
	t.Tac = octets
	return 
}
