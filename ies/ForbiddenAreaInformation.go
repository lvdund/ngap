package ies

import "github.com/lvdund/ngap/aper"

type ForbiddenAreaInformation struct {
	Value []*ForbiddenAreaInformationItem `False,1,16`
}

func (ie *ForbiddenAreaInformation) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*ForbiddenAreaInformationItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return
	}
	return
}
func (ie *ForbiddenAreaInformation) Decode(r *aper.AperReader) (err error) {
	var newItems []*ForbiddenAreaInformationItem
	newItems, err = aper.ReadSequenceOfEx[*ForbiddenAreaInformationItem](func() *ForbiddenAreaInformationItem { return new(ForbiddenAreaInformationItem) }, r, &aper.Constraint{Lb: 1, Ub: 16}, false)
	ie.Value = []*ForbiddenAreaInformationItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
