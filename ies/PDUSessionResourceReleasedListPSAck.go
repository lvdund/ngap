package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceReleasedListPSAck struct {
	Value []*PDUSessionResourceReleasedItemPSAck `False,1,256`
}

func (ie *PDUSessionResourceReleasedListPSAck) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceReleasedItemPSAck](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceReleasedListPSAck) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceReleasedItemPSAck
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceReleasedItemPSAck](func() *PDUSessionResourceReleasedItemPSAck { return new(PDUSessionResourceReleasedItemPSAck) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceReleasedItemPSAck{}
	ie.Value = append(ie.Value, newItems...)
	return
}
