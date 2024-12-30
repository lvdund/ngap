package ies

import "github.com/lvdund/ngap/aper"

const (
	TimeToWaitV1S  aper.Enumerated = 0
	TimeToWaitV2S  aper.Enumerated = 1
	TimeToWaitV5S  aper.Enumerated = 2
	TimeToWaitV10S aper.Enumerated = 3
	TimeToWaitV20S aper.Enumerated = 4
	TimeToWaitV60S aper.Enumerated = 5
)

type TimeToWait struct {
	Value aper.Enumerated
}

func (ie *TimeToWait) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, true)
	return
}
func (ie *TimeToWait) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, true)
	ie.Value = aper.Enumerated(v)
	return
}
