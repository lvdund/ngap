package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceSwitchedList struct {
	Value []*PDUSessionResourceSwitchedItem `False,1,256`
}

func (ie *PDUSessionResourceSwitchedList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceSwitchedItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceSwitchedList) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceSwitchedItem
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceSwitchedItem](func() *PDUSessionResourceSwitchedItem { return new(PDUSessionResourceSwitchedItem) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceSwitchedItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
