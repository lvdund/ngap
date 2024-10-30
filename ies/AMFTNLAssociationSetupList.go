package ies

import "github.com/lvdund/ngap/aper"

type AMFTNLAssociationSetupList struct {
	Value []*AMFTNLAssociationSetupItem `False,1,32`
}

func (ie *AMFTNLAssociationSetupList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*AMFTNLAssociationSetupItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return
	}
	return
}
func (ie *AMFTNLAssociationSetupList) Decode(r *aper.AperReader) (err error) {
	var newItems []*AMFTNLAssociationSetupItem
	newItems, err = aper.ReadSequenceOfEx[*AMFTNLAssociationSetupItem](func() *AMFTNLAssociationSetupItem { return new(AMFTNLAssociationSetupItem) }, r, &aper.Constraint{Lb: 1, Ub: 32}, false)
	ie.Value = []*AMFTNLAssociationSetupItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
