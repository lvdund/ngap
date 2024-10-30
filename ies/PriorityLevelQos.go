package ies

import "github.com/lvdund/ngap/aper"

type PriorityLevelQos struct {
	Value aper.Integer `True,1,127`
}

func (ie *PriorityLevelQos) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 1, Ub: 127}, true)
	return
}
func (ie *PriorityLevelQos) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 127}, true); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
