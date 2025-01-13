package ies

import "github.com/lvdund/ngap/aper"

const (
	RATTypeNr    aper.Enumerated = 0
	RATTypeEutra aper.Enumerated = 1
)

type RATType struct {
	Value aper.Enumerated
}

func (ie *RATType) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}
func (ie *RATType) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
