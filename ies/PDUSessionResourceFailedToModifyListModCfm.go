package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceFailedToModifyListModCfm struct {
	Value []*PDUSessionResourceFailedToModifyItemModCfm `False,1,256`
}

func (ie *PDUSessionResourceFailedToModifyListModCfm) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceFailedToModifyItemModCfm](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceFailedToModifyListModCfm) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceFailedToModifyItemModCfm
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceFailedToModifyItemModCfm](func() *PDUSessionResourceFailedToModifyItemModCfm {
		return new(PDUSessionResourceFailedToModifyItemModCfm)
	}, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceFailedToModifyItemModCfm{}
	ie.Value = append(ie.Value, newItems...)
	return
}
