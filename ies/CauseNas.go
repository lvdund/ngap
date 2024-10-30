package ies

import "github.com/lvdund/ngap/aper"

type CauseNas struct {
	Value aper.Enumerated `True,0,4`
}

func (ie *CauseNas) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, true)
	return
}
func (ie *CauseNas) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, true)
	ie.Value = aper.Enumerated(v)
	return
}
