package ies

import "github.com/lvdund/ngap/aper"

type UEassociatedLogicalNGconnectionList struct {
	Value []*UEassociatedLogicalNGconnectionItem `False,1,65536`
}

func (ie *UEassociatedLogicalNGconnectionList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*UEassociatedLogicalNGconnectionItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65536}, false); err != nil {
		return
	}
	return
}
func (ie *UEassociatedLogicalNGconnectionList) Decode(r *aper.AperReader) (err error) {
	var newItems []*UEassociatedLogicalNGconnectionItem
	newItems, err = aper.ReadSequenceOfEx[*UEassociatedLogicalNGconnectionItem](func() *UEassociatedLogicalNGconnectionItem { return new(UEassociatedLogicalNGconnectionItem) }, r, &aper.Constraint{Lb: 1, Ub: 65536}, false)
	ie.Value = []*UEassociatedLogicalNGconnectionItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
