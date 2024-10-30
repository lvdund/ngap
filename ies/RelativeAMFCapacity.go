package ies

import "github.com/lvdund/ngap/aper"

type RelativeAMFCapacity struct {
	Value aper.Integer `False,0,255`
}

func (ie *RelativeAMFCapacity) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 255}, false)
	return
}
func (ie *RelativeAMFCapacity) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
