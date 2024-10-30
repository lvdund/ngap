package ies

import "github.com/lvdund/ngap/aper"

type ExpectedActivityPeriod struct {
	Value aper.Integer `True,1,30`
}

func (ie *ExpectedActivityPeriod) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 1, Ub: 30}, true)
	return
}
func (ie *ExpectedActivityPeriod) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 30}, true); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
