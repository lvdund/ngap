package ies

import "github.com/lvdund/ngap/aper"

type RRCEstablishmentCause struct {
	Value aper.Enumerated `True,0,10`
}

func (ie *RRCEstablishmentCause) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 10}, true)
	return
}
func (ie *RRCEstablishmentCause) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 10}, true)
	ie.Value = aper.Enumerated(v)
	return
}
