package ies

import "github.com/lvdund/ngap/aper"

type TrafficLoadReductionIndication struct {
	Value aper.Integer `False,1,99`
}

func (ie *TrafficLoadReductionIndication) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 1, Ub: 99}, false)
	return
}
func (ie *TrafficLoadReductionIndication) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 99}, false); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
