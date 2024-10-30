package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceModifyListModRes struct {
	Value []*PDUSessionResourceModifyItemModRes `False,1,256`
}

func (ie *PDUSessionResourceModifyListModRes) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceModifyItemModRes](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceModifyListModRes) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceModifyItemModRes
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceModifyItemModRes](func() *PDUSessionResourceModifyItemModRes { return new(PDUSessionResourceModifyItemModRes) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceModifyItemModRes{}
	ie.Value = append(ie.Value, newItems...)
	return
}
