package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceHandoverList struct {
	Value []*PDUSessionResourceHandoverItem `False,1,256`
}

func (ie *PDUSessionResourceHandoverList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceHandoverItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceHandoverList) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceHandoverItem
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceHandoverItem](func() *PDUSessionResourceHandoverItem { return new(PDUSessionResourceHandoverItem) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceHandoverItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
