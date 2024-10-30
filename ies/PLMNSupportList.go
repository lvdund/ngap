package ies

import "github.com/lvdund/ngap/aper"

type PLMNSupportList struct {
	Value []*PLMNSupportItem `False,1,12`
}

func (ie *PLMNSupportList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PLMNSupportItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 12}, false); err != nil {
		return
	}
	return
}
func (ie *PLMNSupportList) Decode(r *aper.AperReader) (err error) {
	var newItems []*PLMNSupportItem
	newItems, err = aper.ReadSequenceOfEx[*PLMNSupportItem](func() *PLMNSupportItem { return new(PLMNSupportItem) }, r, &aper.Constraint{Lb: 1, Ub: 12}, false)
	ie.Value = []*PLMNSupportItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
