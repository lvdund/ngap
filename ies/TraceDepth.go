package ies

import "github.com/lvdund/ngap/aper"

const (
	TraceDepthMinimum                               aper.Enumerated = 0
	TraceDepthMedium                                aper.Enumerated = 1
	TraceDepthMaximum                               aper.Enumerated = 2
	TraceDepthMinimumwithoutvendorspecificextension aper.Enumerated = 3
	TraceDepthMediumwithoutvendorspecificextension  aper.Enumerated = 4
	TraceDepthMaximumwithoutvendorspecificextension aper.Enumerated = 5
)

type TraceDepth struct {
	Value aper.Enumerated
}

func (ie *TraceDepth) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, true)
	return
}
func (ie *TraceDepth) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, true)
	ie.Value = aper.Enumerated(v)
	return
}
