package ies

import "github.com/lvdund/ngap/aper"

type CancelledCellsInEAINR struct {
	Value []*CancelledCellsInEAINRItem `False,1,65535`
}

func (ie *CancelledCellsInEAINR) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*CancelledCellsInEAINRItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *CancelledCellsInEAINR) Decode(r *aper.AperReader) (err error) {
	var newItems []*CancelledCellsInEAINRItem
	newItems, err = aper.ReadSequenceOfEx[*CancelledCellsInEAINRItem](func() *CancelledCellsInEAINRItem { return new(CancelledCellsInEAINRItem) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*CancelledCellsInEAINRItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
