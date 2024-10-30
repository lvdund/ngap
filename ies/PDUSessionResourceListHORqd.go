package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceListHORqd struct {
	Value []*PDUSessionResourceItemHORqd `False,1,256`
}

func (ie *PDUSessionResourceListHORqd) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceItemHORqd](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceListHORqd) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceItemHORqd
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceItemHORqd](func() *PDUSessionResourceItemHORqd { return new(PDUSessionResourceItemHORqd) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceItemHORqd{}
	ie.Value = append(ie.Value, newItems...)
	return
}
