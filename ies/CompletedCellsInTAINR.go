package ies

import "github.com/lvdund/ngap/aper"

type CompletedCellsInTAINR struct {
	Value []*CompletedCellsInTAINRItem `False,1,65535`
}

func (ie *CompletedCellsInTAINR) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*CompletedCellsInTAINRItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *CompletedCellsInTAINR) Decode(r *aper.AperReader) (err error) {
	var newItems []*CompletedCellsInTAINRItem
	newItems, err = aper.ReadSequenceOfEx[*CompletedCellsInTAINRItem](func() *CompletedCellsInTAINRItem { return new(CompletedCellsInTAINRItem) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*CompletedCellsInTAINRItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
