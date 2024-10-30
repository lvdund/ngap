package ies

import "github.com/lvdund/ngap/aper"

type CompletedCellsInEAIEUTRA struct {
	Value []*CompletedCellsInEAIEUTRAItem `False,1,65535`
}

func (ie *CompletedCellsInEAIEUTRA) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*CompletedCellsInEAIEUTRAItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *CompletedCellsInEAIEUTRA) Decode(r *aper.AperReader) (err error) {
	var newItems []*CompletedCellsInEAIEUTRAItem
	newItems, err = aper.ReadSequenceOfEx[*CompletedCellsInEAIEUTRAItem](func() *CompletedCellsInEAIEUTRAItem { return new(CompletedCellsInEAIEUTRAItem) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*CompletedCellsInEAIEUTRAItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
