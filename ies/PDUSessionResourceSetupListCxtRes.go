package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceSetupListCxtRes struct {
	Value []*PDUSessionResourceSetupItemCxtRes `False,1,256`
}

func (ie *PDUSessionResourceSetupListCxtRes) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceSetupItemCxtRes](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceSetupListCxtRes) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceSetupItemCxtRes
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceSetupItemCxtRes](func() *PDUSessionResourceSetupItemCxtRes { return new(PDUSessionResourceSetupItemCxtRes) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceSetupItemCxtRes{}
	ie.Value = append(ie.Value, newItems...)
	return
}
