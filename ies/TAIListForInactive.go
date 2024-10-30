package ies

import "github.com/lvdund/ngap/aper"

type TAIListForInactive struct {
	Value []*TAIListForInactiveItem `False,1,16`
}

func (ie *TAIListForInactive) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*TAIListForInactiveItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return
	}
	return
}
func (ie *TAIListForInactive) Decode(r *aper.AperReader) (err error) {
	var newItems []*TAIListForInactiveItem
	newItems, err = aper.ReadSequenceOfEx[*TAIListForInactiveItem](func() *TAIListForInactiveItem { return new(TAIListForInactiveItem) }, r, &aper.Constraint{Lb: 1, Ub: 16}, false)
	ie.Value = []*TAIListForInactiveItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
