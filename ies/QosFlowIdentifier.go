package ies

import "github.com/lvdund/ngap/aper"

type QosFlowIdentifier struct {
	Value aper.Integer `True,0,63`
}

func (ie *QosFlowIdentifier) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 63}, true)
	return
}
func (ie *QosFlowIdentifier) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 63}, true); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
