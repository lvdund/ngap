package ies

import "github.com/lvdund/ngap/aper"

type TNLAssociationList struct {
	Value []*TNLAssociationItem `False,1,32`
}

func (ie *TNLAssociationList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*TNLAssociationItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return
	}
	return
}
func (ie *TNLAssociationList) Decode(r *aper.AperReader) (err error) {
	var newItems []*TNLAssociationItem
	newItems, err = aper.ReadSequenceOfEx[*TNLAssociationItem](func() *TNLAssociationItem { return new(TNLAssociationItem) }, r, &aper.Constraint{Lb: 1, Ub: 32}, false)
	ie.Value = []*TNLAssociationItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
