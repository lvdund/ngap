package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceModifyListModReq struct {
	Value []*PDUSessionResourceModifyItemModReq `False,1,256`
}

func (ie *PDUSessionResourceModifyListModReq) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceModifyItemModReq](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceModifyListModReq) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceModifyItemModReq
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceModifyItemModReq](func() *PDUSessionResourceModifyItemModReq { return new(PDUSessionResourceModifyItemModReq) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceModifyItemModReq{}
	ie.Value = append(ie.Value, newItems...)
	return
}
