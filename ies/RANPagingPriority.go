package ies

import "github.com/lvdund/ngap/aper"

type RANPagingPriority struct {
	Value aper.Integer `False,1,256`
}

func (ie *RANPagingPriority) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 1, Ub: 256}, false)
	return
}
func (ie *RANPagingPriority) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
