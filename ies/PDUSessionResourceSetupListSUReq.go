package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceSetupListSUReq struct {
	Value []*PDUSessionResourceSetupItemSUReq `False,1,256`
}

func (ie *PDUSessionResourceSetupListSUReq) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceSetupItemSUReq](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceSetupListSUReq) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceSetupItemSUReq
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceSetupItemSUReq](func() *PDUSessionResourceSetupItemSUReq { return new(PDUSessionResourceSetupItemSUReq) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceSetupItemSUReq{}
	ie.Value = append(ie.Value, newItems...)
	return
}
