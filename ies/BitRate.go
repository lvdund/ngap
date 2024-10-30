package ies

import "github.com/lvdund/ngap/aper"

type BitRate struct {
	Value aper.Integer `True,0,4000000000000`
}

func (ie *BitRate) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 4000000000000}, true)
	return
}
func (ie *BitRate) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
