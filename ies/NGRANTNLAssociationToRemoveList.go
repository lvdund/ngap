package ies

import "github.com/lvdund/ngap/aper"

type NGRANTNLAssociationToRemoveList struct {
	Value []*NGRANTNLAssociationToRemoveItem `False,1,32`
}

func (ie *NGRANTNLAssociationToRemoveList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*NGRANTNLAssociationToRemoveItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return
	}
	return
}
func (ie *NGRANTNLAssociationToRemoveList) Decode(r *aper.AperReader) (err error) {
	var newItems []*NGRANTNLAssociationToRemoveItem
	newItems, err = aper.ReadSequenceOfEx[*NGRANTNLAssociationToRemoveItem](func() *NGRANTNLAssociationToRemoveItem { return new(NGRANTNLAssociationToRemoveItem) }, r, &aper.Constraint{Lb: 1, Ub: 32}, false)
	ie.Value = []*NGRANTNLAssociationToRemoveItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
