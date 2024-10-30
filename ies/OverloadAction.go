package ies

import "github.com/lvdund/ngap/aper"

type OverloadAction struct {
	Value aper.Enumerated `True,0,4`
}

func (ie *OverloadAction) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, true)
	return
}
func (ie *OverloadAction) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, true)
	ie.Value = aper.Enumerated(v)
	return
}
