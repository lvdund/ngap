package ies

import "github.com/lvdund/ngap/aper"

type PagingAttemptCount struct {
	Value aper.Integer `True,1,16`
}

func (ie *PagingAttemptCount) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 1, Ub: 16}, true)
	return
}
func (ie *PagingAttemptCount) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 16}, true); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
