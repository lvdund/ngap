package ies

import "github.com/lvdund/ngap/aper"

type AMFTNLAssociationToAddList struct {
	Value []*AMFTNLAssociationToAddItem `False,1,32`
}

func (ie *AMFTNLAssociationToAddList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*AMFTNLAssociationToAddItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return
	}
	return
}
func (ie *AMFTNLAssociationToAddList) Decode(r *aper.AperReader) (err error) {
	var newItems []*AMFTNLAssociationToAddItem
	newItems, err = aper.ReadSequenceOfEx[*AMFTNLAssociationToAddItem](func() *AMFTNLAssociationToAddItem { return new(AMFTNLAssociationToAddItem) }, r, &aper.Constraint{Lb: 1, Ub: 32}, false)
	ie.Value = []*AMFTNLAssociationToAddItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
