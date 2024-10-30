package ies

import "github.com/lvdund/ngap/aper"

type DRBID struct {
	Value aper.Integer `True,1,32`
}

func (ie *DRBID) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 1, Ub: 32}, true)
	return
}
func (ie *DRBID) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 32}, true); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
