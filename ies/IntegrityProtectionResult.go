package ies

import "github.com/lvdund/ngap/aper"

const (
	IntegrityProtectionResultPerformed    aper.Enumerated = 0
	IntegrityProtectionResultNotperformed aper.Enumerated = 1
)

type IntegrityProtectionResult struct {
	Value aper.Enumerated
}

func (ie *IntegrityProtectionResult) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}
func (ie *IntegrityProtectionResult) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
