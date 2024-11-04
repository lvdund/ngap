package ies

import "github.com/lvdund/ngap/aper"

const (
	PagingDRXV32  aper.Enumerated = 0
	PagingDRXV64  aper.Enumerated = 1
	PagingDRXV128 aper.Enumerated = 2
	PagingDRXV256 aper.Enumerated = 3
)

type PagingDRX struct {
	Value aper.Enumerated `True,0,3`
}

func (ie *PagingDRX) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}
func (ie *PagingDRX) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
