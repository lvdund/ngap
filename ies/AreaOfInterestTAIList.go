package ies

import "github.com/lvdund/ngap/aper"

type AreaOfInterestTAIList struct {
	Value []*AreaOfInterestTAIItem `False,1,16`
}

func (ie *AreaOfInterestTAIList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*AreaOfInterestTAIItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return
	}
	return
}
func (ie *AreaOfInterestTAIList) Decode(r *aper.AperReader) (err error) {
	var newItems []*AreaOfInterestTAIItem
	newItems, err = aper.ReadSequenceOfEx[*AreaOfInterestTAIItem](func() *AreaOfInterestTAIItem { return new(AreaOfInterestTAIItem) }, r, &aper.Constraint{Lb: 1, Ub: 16}, false)
	ie.Value = []*AreaOfInterestTAIItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
