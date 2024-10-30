package ies

import "github.com/lvdund/ngap/aper"

type UEPresence struct {
	Value aper.Enumerated `True,0,3`
}

func (ie *UEPresence) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}
func (ie *UEPresence) Decode(r *aper.AperReader) (err error) {
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true); err != nil {
		return
	}
	ie.Value = aper.Enumerated(v)
	return
}
