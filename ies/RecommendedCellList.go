package ies

import "github.com/lvdund/ngap/aper"

type RecommendedCellList struct {
	Value []*RecommendedCellItem `False,1,16`
}

func (ie *RecommendedCellList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*RecommendedCellItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return
	}
	return
}
func (ie *RecommendedCellList) Decode(r *aper.AperReader) (err error) {
	var newItems []*RecommendedCellItem
	newItems, err = aper.ReadSequenceOfEx[*RecommendedCellItem](func() *RecommendedCellItem { return new(RecommendedCellItem) }, r, &aper.Constraint{Lb: 1, Ub: 16}, false)
	ie.Value = []*RecommendedCellItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
