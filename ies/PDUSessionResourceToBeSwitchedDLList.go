package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceToBeSwitchedDLList struct {
	Value []*PDUSessionResourceToBeSwitchedDLItem `False,1,256`
}

func (ie *PDUSessionResourceToBeSwitchedDLList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceToBeSwitchedDLItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceToBeSwitchedDLList) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceToBeSwitchedDLItem
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceToBeSwitchedDLItem](func() *PDUSessionResourceToBeSwitchedDLItem { return new(PDUSessionResourceToBeSwitchedDLItem) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceToBeSwitchedDLItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
