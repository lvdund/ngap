package ies

import "github.com/lvdund/ngap/aper"

const (
	CellSizeVerysmall aper.Enumerated = 0
	CellSizeSmall     aper.Enumerated = 1
	CellSizeMedium    aper.Enumerated = 2
	CellSizeLarge     aper.Enumerated = 3
)

type CellSize struct {
	Value aper.Enumerated `True,0,3`
}

func (ie *CellSize) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}
func (ie *CellSize) Decode(r *aper.AperReader) (err error) {
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true); err != nil {
		return
	}
	ie.Value = aper.Enumerated(v)
	return
}
