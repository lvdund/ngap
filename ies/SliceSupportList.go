package ies

import "github.com/lvdund/ngap/aper"

type SliceSupportList struct {
	Value []*SliceSupportItem `False,1,1024`
}

func (ie *SliceSupportList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*SliceSupportItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 1024}, false); err != nil {
		return
	}
	return
}
func (ie *SliceSupportList) Decode(r *aper.AperReader) (err error) {
	var newItems []*SliceSupportItem
	newItems, err = aper.ReadSequenceOfEx[*SliceSupportItem](func() *SliceSupportItem { return new(SliceSupportItem) }, r, &aper.Constraint{Lb: 1, Ub: 1024}, false)
	ie.Value = []*SliceSupportItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
