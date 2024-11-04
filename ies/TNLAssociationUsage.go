package ies

import "github.com/lvdund/ngap/aper"

const (
	TNLAssociationUsageUe    aper.Enumerated = 0
	TNLAssociationUsageNonue aper.Enumerated = 1
	TNLAssociationUsageBoth  aper.Enumerated = 2
)

type TNLAssociationUsage struct {
	Value aper.Enumerated `True,0,2`
}

func (ie *TNLAssociationUsage) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
	return
}
func (ie *TNLAssociationUsage) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	ie.Value = aper.Enumerated(v)
	return
}
