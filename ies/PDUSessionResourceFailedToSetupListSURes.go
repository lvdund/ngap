package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceFailedToSetupListSURes struct {
	Value []*PDUSessionResourceFailedToSetupItemSURes `False,1,256`
}

func (ie *PDUSessionResourceFailedToSetupListSURes) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceFailedToSetupItemSURes](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceFailedToSetupListSURes) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceFailedToSetupItemSURes
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceFailedToSetupItemSURes](func() *PDUSessionResourceFailedToSetupItemSURes { return new(PDUSessionResourceFailedToSetupItemSURes) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceFailedToSetupItemSURes{}
	ie.Value = append(ie.Value, newItems...)
	return
}
