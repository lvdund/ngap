package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceSetupListHOReq struct {
	Value []*PDUSessionResourceSetupItemHOReq `False,1,256`
}

func (ie *PDUSessionResourceSetupListHOReq) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceSetupItemHOReq](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceSetupListHOReq) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceSetupItemHOReq
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceSetupItemHOReq](func() *PDUSessionResourceSetupItemHOReq { return new(PDUSessionResourceSetupItemHOReq) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceSetupItemHOReq{}
	ie.Value = append(ie.Value, newItems...)
	return
}
