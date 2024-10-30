package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceModifyListModInd struct {
	Value []*PDUSessionResourceModifyItemModInd `False,1,256`
}

func (ie *PDUSessionResourceModifyListModInd) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceModifyItemModInd](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceModifyListModInd) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceModifyItemModInd
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceModifyItemModInd](func() *PDUSessionResourceModifyItemModInd { return new(PDUSessionResourceModifyItemModInd) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceModifyItemModInd{}
	ie.Value = append(ie.Value, newItems...)
	return
}
