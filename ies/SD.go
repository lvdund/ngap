package ies

import "github.com/lvdund/ngap/aper"

type SD struct {
	Value aper.OctetString `False,3,3`
}

func (ie *SD) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteOctetString(ie.Value, &aper.Constraint{Lb: 3, Ub: 3}, false)
	return
}
func (ie *SD) Decode(r *aper.AperReader) (err error) {
	var v aper.OctetString
	if v, err = r.ReadOctetString(&aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return
	}
	ie.Value = v
	return
}
