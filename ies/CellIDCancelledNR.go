package ies

import "github.com/lvdund/ngap/aper"

type CellIDCancelledNR struct {
	Value []*CellIDCancelledNRItem `False,1,65535`
}

func (ie *CellIDCancelledNR) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*CellIDCancelledNRItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *CellIDCancelledNR) Decode(r *aper.AperReader) (err error) {
	var newItems []*CellIDCancelledNRItem
	newItems, err = aper.ReadSequenceOfEx[*CellIDCancelledNRItem](func() *CellIDCancelledNRItem { return new(CellIDCancelledNRItem) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*CellIDCancelledNRItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
