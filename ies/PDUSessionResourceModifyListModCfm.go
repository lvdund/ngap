package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceModifyListModCfm struct {
	Value []*PDUSessionResourceModifyItemModCfm `False,1,256`
}

func (ie *PDUSessionResourceModifyListModCfm) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceModifyItemModCfm](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceModifyListModCfm) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceModifyItemModCfm
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceModifyItemModCfm](func() *PDUSessionResourceModifyItemModCfm { return new(PDUSessionResourceModifyItemModCfm) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceModifyItemModCfm{}
	ie.Value = append(ie.Value, newItems...)
	return
}
