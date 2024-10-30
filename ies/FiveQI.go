package ies

import "github.com/lvdund/ngap/aper"

type FiveQI struct {
	Value aper.Integer `True,0,255`
}

func (ie *FiveQI) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 255}, true)
	return
}
func (ie *FiveQI) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
