package ies

import "github.com/lvdund/ngap/aper"

type AMFTNLAssociationToRemoveList struct {
	Value []*AMFTNLAssociationToRemoveItem `False,1,32`
}

func (ie *AMFTNLAssociationToRemoveList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*AMFTNLAssociationToRemoveItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return
	}
	return
}
func (ie *AMFTNLAssociationToRemoveList) Decode(r *aper.AperReader) (err error) {
	var newItems []*AMFTNLAssociationToRemoveItem
	newItems, err = aper.ReadSequenceOfEx[*AMFTNLAssociationToRemoveItem](func() *AMFTNLAssociationToRemoveItem { return new(AMFTNLAssociationToRemoveItem) }, r, &aper.Constraint{Lb: 1, Ub: 32}, false)
	ie.Value = []*AMFTNLAssociationToRemoveItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
