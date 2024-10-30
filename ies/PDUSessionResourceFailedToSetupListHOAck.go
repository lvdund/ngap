package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceFailedToSetupListHOAck struct {
	Value []*PDUSessionResourceFailedToSetupItemHOAck `False,1,256`
}

func (ie *PDUSessionResourceFailedToSetupListHOAck) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceFailedToSetupItemHOAck](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceFailedToSetupListHOAck) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceFailedToSetupItemHOAck
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceFailedToSetupItemHOAck](func() *PDUSessionResourceFailedToSetupItemHOAck { return new(PDUSessionResourceFailedToSetupItemHOAck) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceFailedToSetupItemHOAck{}
	ie.Value = append(ie.Value, newItems...)
	return
}
