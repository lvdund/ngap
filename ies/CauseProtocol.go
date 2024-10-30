package ies

import "github.com/lvdund/ngap/aper"

type CauseProtocol struct {
	Value aper.Enumerated `True,0,7`
}

func (ie *CauseProtocol) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, true)
	return
}
func (ie *CauseProtocol) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, true)
	ie.Value = aper.Enumerated(v)
	return
}
