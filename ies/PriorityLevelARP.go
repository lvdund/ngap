package ies

import "github.com/lvdund/ngap/aper"

type PriorityLevelARP struct {
	Value aper.Integer `False,1,15`
}

func (ie *PriorityLevelARP) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 1, Ub: 15}, false)
	return
}
func (ie *PriorityLevelARP) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
