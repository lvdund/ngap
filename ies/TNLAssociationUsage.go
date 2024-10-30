package ies

import "github.com/lvdund/ngap/aper"

type TNLAssociationUsage struct {
	Value aper.Enumerated `True,0,3`
}

func (ie *TNLAssociationUsage) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}
func (ie *TNLAssociationUsage) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
