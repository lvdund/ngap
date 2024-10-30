package ies

import "github.com/lvdund/ngap/aper"

type ERABID struct {
	Value aper.Integer `True,0,15`
}

func (ie *ERABID) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 15}, true)
	return
}
func (ie *ERABID) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 15}, true); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
