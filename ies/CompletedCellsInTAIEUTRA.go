package ies

import "github.com/lvdund/ngap/aper"

type CompletedCellsInTAIEUTRA struct {
	Value []*CompletedCellsInTAIEUTRAItem `False,1,65535`
}

func (ie *CompletedCellsInTAIEUTRA) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*CompletedCellsInTAIEUTRAItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *CompletedCellsInTAIEUTRA) Decode(r *aper.AperReader) (err error) {
	var newItems []*CompletedCellsInTAIEUTRAItem
	newItems, err = aper.ReadSequenceOfEx[*CompletedCellsInTAIEUTRAItem](func() *CompletedCellsInTAIEUTRAItem { return new(CompletedCellsInTAIEUTRAItem) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*CompletedCellsInTAIEUTRAItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
