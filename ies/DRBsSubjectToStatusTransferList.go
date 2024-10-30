package ies

import "github.com/lvdund/ngap/aper"

type DRBsSubjectToStatusTransferList struct {
	Value []*DRBsSubjectToStatusTransferItem `False,1,32`
}

func (ie *DRBsSubjectToStatusTransferList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*DRBsSubjectToStatusTransferItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return
	}
	return
}
func (ie *DRBsSubjectToStatusTransferList) Decode(r *aper.AperReader) (err error) {
	var newItems []*DRBsSubjectToStatusTransferItem
	newItems, err = aper.ReadSequenceOfEx[*DRBsSubjectToStatusTransferItem](func() *DRBsSubjectToStatusTransferItem { return new(DRBsSubjectToStatusTransferItem) }, r, &aper.Constraint{Lb: 1, Ub: 32}, false)
	ie.Value = []*DRBsSubjectToStatusTransferItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
