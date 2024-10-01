package ie

import (
	"ngap/aper"
)

type RANNodeName struct {
	Value string `aper:"sizeExt,sizeLB:1,sizeUB:150"`
}

func (e *RANNodeName) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteOctetString([]byte(e.Value), &aper.Constraint{
		Lb: 1,
		Ub: 150,
	}, true)
	return
}

func (e *RANNodeName) Decode(r *aper.AperReader) (err error) {
	var octets []byte
	if octets, err = r.ReadOctetString(&aper.Constraint{
		Lb: 1,
		Ub: 150,
	}, true); err != nil {
		return
	}

	e.Value = string(octets)

	// fmt.Printf("Decoded RANNodeName:%s[%d]\n", e.Value, len(octets))
	return
}
