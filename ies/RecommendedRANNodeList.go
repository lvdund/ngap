package ies

import "github.com/lvdund/ngap/aper"

type RecommendedRANNodeList struct {
	Value []*RecommendedRANNodeItem `False,1,16`
}

func (ie *RecommendedRANNodeList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*RecommendedRANNodeItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return
	}
	return
}
func (ie *RecommendedRANNodeList) Decode(r *aper.AperReader) (err error) {
	var newItems []*RecommendedRANNodeItem
	newItems, err = aper.ReadSequenceOfEx[*RecommendedRANNodeItem](func() *RecommendedRANNodeItem { return new(RecommendedRANNodeItem) }, r, &aper.Constraint{Lb: 1, Ub: 16}, false)
	ie.Value = []*RecommendedRANNodeItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
