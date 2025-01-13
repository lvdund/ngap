package ies

import "github.com/lvdund/ngap/aper"

const (
	CnTypeEpcForbidden aper.Enumerated = 0
	CnType5GCForbidden aper.Enumerated = 1
)

type CnType struct {
	Value aper.Enumerated
}

func (ie *CnType) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}
func (ie *CnType) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
