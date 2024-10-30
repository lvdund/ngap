package ies

import "github.com/lvdund/ngap/aper"

type TimerApproachForGUAMIRemoval struct {
	Value aper.Enumerated `True,0,1`
}

func (ie *TimerApproachForGUAMIRemoval) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}
func (ie *TimerApproachForGUAMIRemoval) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
