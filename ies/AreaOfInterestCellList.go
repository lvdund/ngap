package ies

import "github.com/lvdund/ngap/aper"

type AreaOfInterestCellList struct {
	Value []*AreaOfInterestCellItem `False,1,256`
}

func (ie *AreaOfInterestCellList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*AreaOfInterestCellItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *AreaOfInterestCellList) Decode(r *aper.AperReader) (err error) {
	var newItems []*AreaOfInterestCellItem
	newItems, err = aper.ReadSequenceOfEx[*AreaOfInterestCellItem](func() *AreaOfInterestCellItem { return new(AreaOfInterestCellItem) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*AreaOfInterestCellItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
