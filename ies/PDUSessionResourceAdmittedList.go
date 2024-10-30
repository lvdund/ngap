package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceAdmittedList struct {
	Value []*PDUSessionResourceAdmittedItem `False,1,256`
}

func (ie *PDUSessionResourceAdmittedList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceAdmittedItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceAdmittedList) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceAdmittedItem
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceAdmittedItem](func() *PDUSessionResourceAdmittedItem { return new(PDUSessionResourceAdmittedItem) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceAdmittedItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
