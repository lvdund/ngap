package ies

import "github.com/lvdund/ngap/aper"

type RATRestrictions struct {
	Value []*RATRestrictionsItem `False,1,16`
}

func (ie *RATRestrictions) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*RATRestrictionsItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return
	}
	return
}
func (ie *RATRestrictions) Decode(r *aper.AperReader) (err error) {
	var newItems []*RATRestrictionsItem
	newItems, err = aper.ReadSequenceOfEx[*RATRestrictionsItem](func() *RATRestrictionsItem { return new(RATRestrictionsItem) }, r, &aper.Constraint{Lb: 1, Ub: 16}, false)
	ie.Value = []*RATRestrictionsItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
