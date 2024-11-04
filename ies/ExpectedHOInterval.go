package ies

import "github.com/lvdund/ngap/aper"

const (
	ExpectedHOIntervalSec15  aper.Enumerated = 0
	ExpectedHOIntervalSec30  aper.Enumerated = 1
	ExpectedHOIntervalSec60  aper.Enumerated = 2
	ExpectedHOIntervalSec90  aper.Enumerated = 3
	ExpectedHOIntervalSec120 aper.Enumerated = 4
	ExpectedHOIntervalSec180 aper.Enumerated = 5
	ExpectedHOIntervalLongTime aper.Enumerated = 6
)

type ExpectedHOInterval struct {
	Value aper.Enumerated `True,0,6`
}

func (ie *ExpectedHOInterval) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, true)
	return
}
func (ie *ExpectedHOInterval) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, true)
	ie.Value = aper.Enumerated(v)
	return
}
