package ies

import "github.com/lvdund/ngap/aper"

const (
	TimeToWaitV1s  aper.Enumerated = 0
	TimeToWaitV2s  aper.Enumerated = 1
	TimeToWaitV5s  aper.Enumerated = 2
	TimeToWaitV10s aper.Enumerated = 3
	TimeToWaitV20s aper.Enumerated = 4
	TimeToWaitV60s aper.Enumerated = 5
)

type TimeToWait struct {
	Value aper.Enumerated `True,0,5`
}

func (ie *TimeToWait) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, true)
	return
}
func (ie *TimeToWait) Decode(r *aper.AperReader) (err error) {
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, true); err != nil {
		return
	}
	ie.Value = aper.Enumerated(v)
	return
}
