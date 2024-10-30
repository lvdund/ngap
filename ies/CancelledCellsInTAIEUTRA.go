package ies

import "github.com/lvdund/ngap/aper"

type CancelledCellsInTAIEUTRA struct {
	Value []*CancelledCellsInTAIEUTRAItem `False,1,65535`
}

func (ie *CancelledCellsInTAIEUTRA) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*CancelledCellsInTAIEUTRAItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *CancelledCellsInTAIEUTRA) Decode(r *aper.AperReader) (err error) {
	var newItems []*CancelledCellsInTAIEUTRAItem
	newItems, err = aper.ReadSequenceOfEx[*CancelledCellsInTAIEUTRAItem](func() *CancelledCellsInTAIEUTRAItem { return new(CancelledCellsInTAIEUTRAItem) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*CancelledCellsInTAIEUTRAItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
