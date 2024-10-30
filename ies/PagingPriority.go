package ies

import "github.com/lvdund/ngap/aper"

type PagingPriority struct {
	Value aper.Enumerated `True,0,8`
}

func (ie *PagingPriority) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 8}, true)
	return
}
func (ie *PagingPriority) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 8}, true)
	ie.Value = aper.Enumerated(v)
	return
}
