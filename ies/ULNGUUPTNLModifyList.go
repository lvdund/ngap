package ies

import "github.com/lvdund/ngap/aper"

type ULNGUUPTNLModifyList struct {
	Value []*ULNGUUPTNLModifyItem `False,1,4`
}

func (ie *ULNGUUPTNLModifyList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*ULNGUUPTNLModifyItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return
	}
	return
}
func (ie *ULNGUUPTNLModifyList) Decode(r *aper.AperReader) (err error) {
	var newItems []*ULNGUUPTNLModifyItem
	newItems, err = aper.ReadSequenceOfEx[*ULNGUUPTNLModifyItem](func() *ULNGUUPTNLModifyItem { return new(ULNGUUPTNLModifyItem) }, r, &aper.Constraint{Lb: 1, Ub: 4}, false)
	ie.Value = []*ULNGUUPTNLModifyItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
