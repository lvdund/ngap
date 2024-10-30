package ies

import "github.com/lvdund/ngap/aper"

type RepetitionPeriod struct {
	Value aper.Integer `False,0,131071`
}

func (ie *RepetitionPeriod) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 131071}, false)
	return
}
func (ie *RepetitionPeriod) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 131071}, false); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
