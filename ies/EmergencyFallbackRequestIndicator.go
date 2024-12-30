package ies

import "github.com/lvdund/ngap/aper"

const (
	EmergencyFallbackRequestIndicatorEmergencyfallbackrequested aper.Enumerated = 0
)

type EmergencyFallbackRequestIndicator struct {
	Value aper.Enumerated
}

func (ie *EmergencyFallbackRequestIndicator) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, true)
	return
}
func (ie *EmergencyFallbackRequestIndicator) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true)
	ie.Value = aper.Enumerated(v)
	return
}
