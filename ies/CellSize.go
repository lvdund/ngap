package ies

import "github.com/lvdund/ngap/aper"

type CellSize struct {
	Value aper.Enumerated `True,0,4`
}

func (ie *CellSize) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, true)
	return
}
func (ie *CellSize) Decode(r *aper.AperReader) (err error) {
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, true); err != nil {
		return
	}
	ie.Value = aper.Enumerated(v)
	return
}
