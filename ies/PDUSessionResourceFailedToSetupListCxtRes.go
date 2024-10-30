package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceFailedToSetupListCxtRes struct {
	Value []*PDUSessionResourceFailedToSetupItemCxtRes `False,1,256`
}

func (ie *PDUSessionResourceFailedToSetupListCxtRes) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceFailedToSetupItemCxtRes](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceFailedToSetupListCxtRes) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceFailedToSetupItemCxtRes
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceFailedToSetupItemCxtRes](func() *PDUSessionResourceFailedToSetupItemCxtRes {
		return new(PDUSessionResourceFailedToSetupItemCxtRes)
	}, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceFailedToSetupItemCxtRes{}
	ie.Value = append(ie.Value, newItems...)
	return
}
