package ies

import "github.com/lvdund/ngap/aper"

type AreaOfInterestRANNodeList struct {
	Value []*AreaOfInterestRANNodeItem `False,1,64`
}

func (ie *AreaOfInterestRANNodeList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*AreaOfInterestRANNodeItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return
	}
	return
}
func (ie *AreaOfInterestRANNodeList) Decode(r *aper.AperReader) (err error) {
	var newItems []*AreaOfInterestRANNodeItem
	newItems, err = aper.ReadSequenceOfEx[*AreaOfInterestRANNodeItem](func() *AreaOfInterestRANNodeItem { return new(AreaOfInterestRANNodeItem) }, r, &aper.Constraint{Lb: 1, Ub: 64}, false)
	ie.Value = []*AreaOfInterestRANNodeItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
