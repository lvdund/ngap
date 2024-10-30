package ies

import "github.com/lvdund/ngap/aper"

type SourceOfUEActivityBehaviourInformation struct {
	Value aper.Enumerated `True,0,2`
}

func (ie *SourceOfUEActivityBehaviourInformation) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
	return
}
func (ie *SourceOfUEActivityBehaviourInformation) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	ie.Value = aper.Enumerated(v)
	return
}
