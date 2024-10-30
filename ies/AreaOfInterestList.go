package ies

import "github.com/lvdund/ngap/aper"

type AreaOfInterestList struct {
	Value []*AreaOfInterestItem `False,1,64`
}

func (ie *AreaOfInterestList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*AreaOfInterestItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return
	}
	return
}
func (ie *AreaOfInterestList) Decode(r *aper.AperReader) (err error) {
	var newItems []*AreaOfInterestItem
	newItems, err = aper.ReadSequenceOfEx[*AreaOfInterestItem](func() *AreaOfInterestItem { return new(AreaOfInterestItem) }, r, &aper.Constraint{Lb: 1, Ub: 64}, false)
	ie.Value = []*AreaOfInterestItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
