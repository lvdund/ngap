package ies

import "github.com/lvdund/ngap/aper"

type CNTypeRestrictionsForEquivalent struct {
	Value []*CNTypeRestrictionsForEquivalentItem `False,1,15`
}

func (ie *CNTypeRestrictionsForEquivalent) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*CNTypeRestrictionsForEquivalentItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return
	}
	return
}
func (ie *CNTypeRestrictionsForEquivalent) Decode(r *aper.AperReader) (err error) {
	var newItems []*CNTypeRestrictionsForEquivalentItem
	newItems, err = aper.ReadSequenceOfEx[*CNTypeRestrictionsForEquivalentItem](func() *CNTypeRestrictionsForEquivalentItem { return new(CNTypeRestrictionsForEquivalentItem) }, r, &aper.Constraint{Lb: 1, Ub: 15}, false)
	ie.Value = []*CNTypeRestrictionsForEquivalentItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
