package ies

import "github.com/lvdund/ngap/aper"

type CompletedCellsInEAINR struct {
	Value []*CompletedCellsInEAINRItem `False,1,65535`
}

func (ie *CompletedCellsInEAINR) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*CompletedCellsInEAINRItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *CompletedCellsInEAINR) Decode(r *aper.AperReader) (err error) {
	var newItems []*CompletedCellsInEAINRItem
	newItems, err = aper.ReadSequenceOfEx[*CompletedCellsInEAINRItem](func() *CompletedCellsInEAINRItem { return new(CompletedCellsInEAINRItem) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*CompletedCellsInEAINRItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
