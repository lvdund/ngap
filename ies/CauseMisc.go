package ies

import "github.com/lvdund/ngap/aper"

type CauseMisc struct {
	Value aper.Enumerated `True,0,6`
}

func (ie *CauseMisc) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, true)
	return
}
func (ie *CauseMisc) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, true)
	ie.Value = aper.Enumerated(v)
	return
}
