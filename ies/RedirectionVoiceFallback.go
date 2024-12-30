package ies

import "github.com/lvdund/ngap/aper"

const (
	RedirectionVoiceFallbackPossible    aper.Enumerated = 0
	RedirectionVoiceFallbackNotpossible aper.Enumerated = 1
)

type RedirectionVoiceFallback struct {
	Value aper.Enumerated
}

func (ie *RedirectionVoiceFallback) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}
func (ie *RedirectionVoiceFallback) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
