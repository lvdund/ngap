package ies

import "github.com/lvdund/ngap/aper"

const (
	PreemptionCapabilityShallnottriggerpreemption aper.Enumerated = 0
	PreemptionCapabilityMaytriggerpreemption      aper.Enumerated = 1
)

type PreemptionCapability struct {
	Value aper.Enumerated
}

func (ie *PreemptionCapability) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}
func (ie *PreemptionCapability) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
