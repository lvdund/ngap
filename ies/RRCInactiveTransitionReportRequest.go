package ies

import "github.com/lvdund/ngap/aper"

type RRCInactiveTransitionReportRequest struct {
	Value aper.Enumerated `True,0,3`
}

func (ie *RRCInactiveTransitionReportRequest) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}
func (ie *RRCInactiveTransitionReportRequest) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
