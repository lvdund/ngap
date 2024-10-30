package ies

import "github.com/lvdund/ngap/aper"

type CancelledCellsInEAIEUTRA struct {
	Value []*CancelledCellsInEAIEUTRAItem `False,1,65535`
}

func (ie *CancelledCellsInEAIEUTRA) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*CancelledCellsInEAIEUTRAItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *CancelledCellsInEAIEUTRA) Decode(r *aper.AperReader) (err error) {
	var newItems []*CancelledCellsInEAIEUTRAItem
	newItems, err = aper.ReadSequenceOfEx[*CancelledCellsInEAIEUTRAItem](func() *CancelledCellsInEAIEUTRAItem { return new(CancelledCellsInEAIEUTRAItem) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*CancelledCellsInEAIEUTRAItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
