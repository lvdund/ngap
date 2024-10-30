package ies

import "github.com/lvdund/ngap/aper"

type TimeUEStayedInCellEnhancedGranularity struct {
	Value aper.Integer `False,0,40950`
}

func (ie *TimeUEStayedInCellEnhancedGranularity) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 40950}, false)
	return
}
func (ie *TimeUEStayedInCellEnhancedGranularity) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 40950}, false); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
