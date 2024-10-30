package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceSecondaryRATUsageList struct {
	Value []*PDUSessionResourceSecondaryRATUsageItem `False,1,256`
}

func (ie *PDUSessionResourceSecondaryRATUsageList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceSecondaryRATUsageItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceSecondaryRATUsageList) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceSecondaryRATUsageItem
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceSecondaryRATUsageItem](func() *PDUSessionResourceSecondaryRATUsageItem { return new(PDUSessionResourceSecondaryRATUsageItem) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceSecondaryRATUsageItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
