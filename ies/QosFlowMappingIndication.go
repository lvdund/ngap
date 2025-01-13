package ies

import "github.com/lvdund/ngap/aper"

const (
	QosFlowMappingIndicationUl aper.Enumerated = 0
	QosFlowMappingIndicationDl aper.Enumerated = 1
)

type QosFlowMappingIndication struct {
	Value aper.Enumerated
}

func (ie *QosFlowMappingIndication) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}
func (ie *QosFlowMappingIndication) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
