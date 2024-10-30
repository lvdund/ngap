package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceReleasedListRelRes struct {
	Value []*PDUSessionResourceReleasedItemRelRes `False,1,256`
}

func (ie *PDUSessionResourceReleasedListRelRes) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceReleasedItemRelRes](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceReleasedListRelRes) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceReleasedItemRelRes
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceReleasedItemRelRes](func() *PDUSessionResourceReleasedItemRelRes { return new(PDUSessionResourceReleasedItemRelRes) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceReleasedItemRelRes{}
	ie.Value = append(ie.Value, newItems...)
	return
}
