package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceFailedToSetupListCxtFail struct {
	Value []*PDUSessionResourceFailedToSetupItemCxtFail `False,1,256`
}

func (ie *PDUSessionResourceFailedToSetupListCxtFail) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceFailedToSetupItemCxtFail](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceFailedToSetupListCxtFail) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceFailedToSetupItemCxtFail
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceFailedToSetupItemCxtFail](func() *PDUSessionResourceFailedToSetupItemCxtFail {
		return new(PDUSessionResourceFailedToSetupItemCxtFail)
	}, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceFailedToSetupItemCxtFail{}
	ie.Value = append(ie.Value, newItems...)
	return
}
