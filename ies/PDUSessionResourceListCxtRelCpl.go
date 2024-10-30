package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceListCxtRelCpl struct {
	Value []*PDUSessionResourceItemCxtRelCpl `False,1,256`
}

func (ie *PDUSessionResourceListCxtRelCpl) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceItemCxtRelCpl](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceListCxtRelCpl) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceItemCxtRelCpl
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceItemCxtRelCpl](func() *PDUSessionResourceItemCxtRelCpl { return new(PDUSessionResourceItemCxtRelCpl) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceItemCxtRelCpl{}
	ie.Value = append(ie.Value, newItems...)
	return
}
