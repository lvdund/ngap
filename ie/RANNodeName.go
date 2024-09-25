package ie

import (
	"ngap/aper"
)

type RANNodeName struct {
	Value string `aper:"sizeExt,sizeLB:1,sizeUB:150"`
}

func (ie *RANNodeName) Encode(w aper.AperWriter) (err error) {
	err = w.WriteOctetString([]byte(ie.Value), &aper.Constraint{
		Lb: 1,
		Ub: 150,
	}, true)
	return
}

func (ie *RANNodeName) Decode(r aper.AperReader) (err error) {
	var octets []byte
	if octets, err = r.ReadOctetString(&aper.Constraint{
		Lb: 1,
		Ub: 150,
	}, true); err != nil {
		return
	}
	ie = &RANNodeName{
		Value: string(octets),
	}
	return
}
