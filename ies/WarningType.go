package ies

import "github.com/lvdund/ngap/aper"

type WarningType struct {
	Value aper.OctetString `False,2,2`
}

func (ie *WarningType) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteOctetString(ie.Value, &aper.Constraint{Lb: 2, Ub: 2}, false)
	return
}
func (ie *WarningType) Decode(r *aper.AperReader) (err error) {
	var v aper.OctetString
	if v, err = r.ReadOctetString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return
	}
	ie.Value = v
	return
}
