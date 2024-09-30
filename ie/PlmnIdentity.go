package ie

import (
	"ngap/aper"
	"fmt"
)

type PlmnIdentity struct {
	PlmnIdentity aper.OctetString `octetstring:"sizeLB:3,sizeUB:3"`
}

func (e *PlmnIdentity) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteOctetString([]byte(e.PlmnIdentity), &aper.Constraint{
		Lb: 3,
		Ub: 3,
	}, false)
	return
}

func (e *PlmnIdentity) Decode(r *aper.AperReader) (err error) {
	var octets []byte
	if octets, err = r.ReadOctetString(&aper.Constraint{
		Lb: 3,
		Ub: 3,
	}, false); err != nil {
		return
	}

	e.PlmnIdentity = octets

	fmt.Printf("Decoded PlmnId: %x[len=%d]\n", octets, len(octets))
	return
}
