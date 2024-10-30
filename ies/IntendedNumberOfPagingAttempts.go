package ies

import "github.com/lvdund/ngap/aper"

type IntendedNumberOfPagingAttempts struct {
	Value aper.Integer `True,1,16`
}

func (ie *IntendedNumberOfPagingAttempts) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 1, Ub: 16}, true)
	return
}
func (ie *IntendedNumberOfPagingAttempts) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 16}, true); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
