package ies

import "github.com/lvdund/ngap/aper"

type AMFTNLAssociationToUpdateList struct {
	Value []*AMFTNLAssociationToUpdateItem `False,1,32`
}

func (ie *AMFTNLAssociationToUpdateList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*AMFTNLAssociationToUpdateItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return
	}
	return
}
func (ie *AMFTNLAssociationToUpdateList) Decode(r *aper.AperReader) (err error) {
	var newItems []*AMFTNLAssociationToUpdateItem
	newItems, err = aper.ReadSequenceOfEx[*AMFTNLAssociationToUpdateItem](func() *AMFTNLAssociationToUpdateItem { return new(AMFTNLAssociationToUpdateItem) }, r, &aper.Constraint{Lb: 1, Ub: 32}, false)
	ie.Value = []*AMFTNLAssociationToUpdateItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
