package ies

import "github.com/lvdund/ngap/aper"

const (
	IntegrityProtectionIndicationRequired  aper.Enumerated = 0
	IntegrityProtectionIndicationPreferred aper.Enumerated = 1
	IntegrityProtectionIndicationNotneeded aper.Enumerated = 2
)

type IntegrityProtectionIndication struct {
	Value aper.Enumerated
}

func (ie *IntegrityProtectionIndication) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
	return
}
func (ie *IntegrityProtectionIndication) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	ie.Value = aper.Enumerated(v)
	return
}
