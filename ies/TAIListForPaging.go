package ies

import "github.com/lvdund/ngap/aper"

type TAIListForPaging struct {
	Value []*TAIListForPagingItem `False,1,16`
}

func (ie *TAIListForPaging) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*TAIListForPagingItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return
	}
	return
}
func (ie *TAIListForPaging) Decode(r *aper.AperReader) (err error) {
	var newItems []*TAIListForPagingItem
	newItems, err = aper.ReadSequenceOfEx[*TAIListForPagingItem](func() *TAIListForPagingItem { return new(TAIListForPagingItem) }, r, &aper.Constraint{Lb: 1, Ub: 16}, false)
	ie.Value = []*TAIListForPagingItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
