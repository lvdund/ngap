package ies

import "github.com/lvdund/ngap/aper"

type TimeUEStayedInCell struct {
	Value aper.Integer `False,0,4095`
}

func (ie *TimeUEStayedInCell) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 4095}, false)
	return
}
func (ie *TimeUEStayedInCell) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
