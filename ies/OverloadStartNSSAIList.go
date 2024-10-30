package ies

import "github.com/lvdund/ngap/aper"

type OverloadStartNSSAIList struct {
	Value []*OverloadStartNSSAIItem `False,1,1024`
}

func (ie *OverloadStartNSSAIList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*OverloadStartNSSAIItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 1024}, false); err != nil {
		return
	}
	return
}
func (ie *OverloadStartNSSAIList) Decode(r *aper.AperReader) (err error) {
	var newItems []*OverloadStartNSSAIItem
	newItems, err = aper.ReadSequenceOfEx[*OverloadStartNSSAIItem](func() *OverloadStartNSSAIItem { return new(OverloadStartNSSAIItem) }, r, &aper.Constraint{Lb: 1, Ub: 1024}, false)
	ie.Value = []*OverloadStartNSSAIItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
