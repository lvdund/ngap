package ies

import "github.com/lvdund/ngap/aper"

type CellIDCancelledEUTRA struct {
	Value []*CellIDCancelledEUTRAItem `False,1,65535`
}

func (ie *CellIDCancelledEUTRA) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*CellIDCancelledEUTRAItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *CellIDCancelledEUTRA) Decode(r *aper.AperReader) (err error) {
	var newItems []*CellIDCancelledEUTRAItem
	newItems, err = aper.ReadSequenceOfEx[*CellIDCancelledEUTRAItem](func() *CellIDCancelledEUTRAItem { return new(CellIDCancelledEUTRAItem) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*CellIDCancelledEUTRAItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
