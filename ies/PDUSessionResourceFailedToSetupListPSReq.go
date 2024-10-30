package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceFailedToSetupListPSReq struct {
	Value []*PDUSessionResourceFailedToSetupItemPSReq `False,1,256`
}

func (ie *PDUSessionResourceFailedToSetupListPSReq) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceFailedToSetupItemPSReq](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceFailedToSetupListPSReq) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceFailedToSetupItemPSReq
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceFailedToSetupItemPSReq](func() *PDUSessionResourceFailedToSetupItemPSReq { return new(PDUSessionResourceFailedToSetupItemPSReq) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceFailedToSetupItemPSReq{}
	ie.Value = append(ie.Value, newItems...)
	return
}
