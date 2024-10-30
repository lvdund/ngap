package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceSetupListCxtReq struct {
	Value []*PDUSessionResourceSetupItemCxtReq `False,1,256`
}

func (ie *PDUSessionResourceSetupListCxtReq) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceSetupItemCxtReq](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceSetupListCxtReq) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceSetupItemCxtReq
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceSetupItemCxtReq](func() *PDUSessionResourceSetupItemCxtReq { return new(PDUSessionResourceSetupItemCxtReq) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceSetupItemCxtReq{}
	ie.Value = append(ie.Value, newItems...)
	return
}
