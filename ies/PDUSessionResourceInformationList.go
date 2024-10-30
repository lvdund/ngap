package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceInformationList struct {
	Value []*PDUSessionResourceInformationItem `False,1,256`
}

func (ie *PDUSessionResourceInformationList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceInformationItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceInformationList) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceInformationItem
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceInformationItem](func() *PDUSessionResourceInformationItem { return new(PDUSessionResourceInformationItem) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceInformationItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
