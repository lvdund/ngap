package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceFailedToModifyListModRes struct {
	Value []*PDUSessionResourceFailedToModifyItemModRes `False,1,256`
}

func (ie *PDUSessionResourceFailedToModifyListModRes) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceFailedToModifyItemModRes](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceFailedToModifyListModRes) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceFailedToModifyItemModRes
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceFailedToModifyItemModRes](func() *PDUSessionResourceFailedToModifyItemModRes {
		return new(PDUSessionResourceFailedToModifyItemModRes)
	}, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceFailedToModifyItemModRes{}
	ie.Value = append(ie.Value, newItems...)
	return
}
