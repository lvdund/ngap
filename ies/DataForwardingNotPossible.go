package ies

import "github.com/lvdund/ngap/aper"

const (
	DataForwardingNotPossibleDataforwardingnotpossible aper.Enumerated = 0
)

type DataForwardingNotPossible struct {
	Value aper.Enumerated `True,0,0`
}

func (ie *DataForwardingNotPossible) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, true)
	return
}
func (ie *DataForwardingNotPossible) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true)
	ie.Value = aper.Enumerated(v)
	return
}
