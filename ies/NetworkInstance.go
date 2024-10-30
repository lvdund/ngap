package ies

import "github.com/lvdund/ngap/aper"

type NetworkInstance struct {
	Value aper.Integer `True,1,256`
}

func (ie *NetworkInstance) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 1, Ub: 256}, true)
	return
}
func (ie *NetworkInstance) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 256}, true); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
