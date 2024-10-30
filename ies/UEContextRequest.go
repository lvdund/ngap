package ies

import "github.com/lvdund/ngap/aper"

type UEContextRequest struct {
	Value aper.Enumerated `True,0,1`
}

func (ie *UEContextRequest) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}
func (ie *UEContextRequest) Decode(r *aper.AperReader) (err error) {
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
		return
	}
	ie.Value = aper.Enumerated(v)
	return
}
