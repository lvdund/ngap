package ies

import "github.com/lvdund/ngap/aper"

type AveragingWindow struct {
	Value aper.Integer `True,0,4095`
}

func (ie *AveragingWindow) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 4095}, true)
	return
}
func (ie *AveragingWindow) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4095}, true); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
