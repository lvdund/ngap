package ies

import "github.com/lvdund/ngap/aper"

type CauseRadioNetwork struct {
	Value aper.Enumerated `True,0,45`
}

func (ie *CauseRadioNetwork) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 45}, true)
	return
}
func (ie *CauseRadioNetwork) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 45}, true)
	ie.Value = aper.Enumerated(v)
	return
}
