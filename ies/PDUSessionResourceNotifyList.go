package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceNotifyList struct {
	Value []*PDUSessionResourceNotifyItem `False,1,256`
}

func (ie *PDUSessionResourceNotifyList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceNotifyItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceNotifyList) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceNotifyItem
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceNotifyItem](func() *PDUSessionResourceNotifyItem { return new(PDUSessionResourceNotifyItem) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceNotifyItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
