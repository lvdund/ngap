package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceReleasedListPSFail struct {
	Value []*PDUSessionResourceReleasedItemPSFail `False,1,256`
}

func (ie *PDUSessionResourceReleasedListPSFail) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceReleasedItemPSFail](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceReleasedListPSFail) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceReleasedItemPSFail
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceReleasedItemPSFail](func() *PDUSessionResourceReleasedItemPSFail { return new(PDUSessionResourceReleasedItemPSFail) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceReleasedItemPSFail{}
	ie.Value = append(ie.Value, newItems...)
	return
}
