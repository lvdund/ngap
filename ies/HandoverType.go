package ies

import "github.com/lvdund/ngap/aper"

const (
	HandoverTypeIntra5Gs    aper.Enumerated = 0
	HandoverTypeFivegstoeps aper.Enumerated = 1
	HandoverTypeEpsto5Gs    aper.Enumerated = 2
)

type HandoverType struct {
	Value aper.Enumerated
}

func (ie *HandoverType) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
	return
}
func (ie *HandoverType) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	ie.Value = aper.Enumerated(v)
	return
}
