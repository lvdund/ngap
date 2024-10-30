package ies

import "github.com/lvdund/ngap/aper"

type NextHopChainingCount struct {
	Value aper.Integer `False,0,7`
}

func (ie *NextHopChainingCount) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 7}, false)
	return
}
func (ie *NextHopChainingCount) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
