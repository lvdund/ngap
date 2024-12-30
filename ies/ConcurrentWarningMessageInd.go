package ies

import "github.com/lvdund/ngap/aper"

const (
	ConcurrentWarningMessageIndTrue aper.Enumerated = 0
)

type ConcurrentWarningMessageInd struct {
	Value aper.Enumerated
}

func (ie *ConcurrentWarningMessageInd) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, true)
	return
}
func (ie *ConcurrentWarningMessageInd) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true)
	ie.Value = aper.Enumerated(v)
	return
}
