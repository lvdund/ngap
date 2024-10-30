package ies

import "github.com/lvdund/ngap/aper"

type CancelledCellsInTAINR struct {
	Value []*CancelledCellsInTAINRItem `False,1,65535`
}

func (ie *CancelledCellsInTAINR) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*CancelledCellsInTAINRItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *CancelledCellsInTAINR) Decode(r *aper.AperReader) (err error) {
	var newItems []*CancelledCellsInTAINRItem
	newItems, err = aper.ReadSequenceOfEx[*CancelledCellsInTAINRItem](func() *CancelledCellsInTAINRItem { return new(CancelledCellsInTAINRItem) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*CancelledCellsInTAINRItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
