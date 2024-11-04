package ies

import "github.com/lvdund/ngap/aper"

const (
	RRCStateInactive  aper.Enumerated = 0
	RRCStateConnected aper.Enumerated = 1
)

type RRCState struct {
	Value aper.Enumerated `True,0,1`
}

func (ie *RRCState) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}
func (ie *RRCState) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
