package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceReleasedListNot struct {
	Value []*PDUSessionResourceReleasedItemNot `False,1,256`
}

func (ie *PDUSessionResourceReleasedListNot) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceReleasedItemNot](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceReleasedListNot) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceReleasedItemNot
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceReleasedItemNot](func() *PDUSessionResourceReleasedItemNot { return new(PDUSessionResourceReleasedItemNot) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceReleasedItemNot{}
	ie.Value = append(ie.Value, newItems...)
	return
}
