package ies

import "github.com/lvdund/ngap/aper"

const (
	ReflectiveQosAttributeSubjectto aper.Enumerated = 0
)

type ReflectiveQosAttribute struct {
	Value aper.Enumerated `True,0,0`
}

func (ie *ReflectiveQosAttribute) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, true)
	return
}
func (ie *ReflectiveQosAttribute) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true)
	ie.Value = aper.Enumerated(v)
	return
}
