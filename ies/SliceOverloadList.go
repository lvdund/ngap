package ies

import "github.com/lvdund/ngap/aper"

type SliceOverloadList struct {
	Value []*SliceOverloadItem `False,1,1024`
}

func (ie *SliceOverloadList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*SliceOverloadItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 1024}, false); err != nil {
		return
	}
	return
}
func (ie *SliceOverloadList) Decode(r *aper.AperReader) (err error) {
	var newItems []*SliceOverloadItem
	newItems, err = aper.ReadSequenceOfEx[*SliceOverloadItem](func() *SliceOverloadItem { return new(SliceOverloadItem) }, r, &aper.Constraint{Lb: 1, Ub: 1024}, false)
	ie.Value = []*SliceOverloadItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
