package ies

import "github.com/lvdund/ngap/aper"

const (
	PagingPriorityPriolevel1 aper.Enumerated = 0
	PagingPriorityPriolevel2 aper.Enumerated = 1
	PagingPriorityPriolevel3 aper.Enumerated = 2
	PagingPriorityPriolevel4 aper.Enumerated = 3
	PagingPriorityPriolevel5 aper.Enumerated = 4
	PagingPriorityPriolevel6 aper.Enumerated = 5
	PagingPriorityPriolevel7 aper.Enumerated = 6
	PagingPriorityPriolevel8 aper.Enumerated = 7
)

type PagingPriority struct {
	Value aper.Enumerated `True,0,7`
}

func (ie *PagingPriority) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, true)
	return
}
func (ie *PagingPriority) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, true)
	ie.Value = aper.Enumerated(v)
	return
}
