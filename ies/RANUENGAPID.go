package ies

import "github.com/lvdund/ngap/aper"

type RANUENGAPID struct {
	Value aper.Integer `False,0,4294967295`
}

func (ie *RANUENGAPID) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 4294967295}, false)
	return
}
func (ie *RANUENGAPID) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
