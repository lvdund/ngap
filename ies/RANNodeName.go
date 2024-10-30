package ies

import "github.com/lvdund/ngap/aper"

type RANNodeName struct {
	Value aper.OctetString `True,1,150`
}

func (ie *RANNodeName) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteOctetString(ie.Value, &aper.Constraint{Lb: 1, Ub: 150}, true)
	return
}
func (ie *RANNodeName) Decode(r *aper.AperReader) (err error) {
	var v aper.OctetString
	if v, err = r.ReadOctetString(&aper.Constraint{Lb: 1, Ub: 150}, true); err != nil {
		return
	}
	ie.Value = v
	return
}
