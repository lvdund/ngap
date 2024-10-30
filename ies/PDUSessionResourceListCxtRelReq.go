package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceListCxtRelReq struct {
	Value []*PDUSessionResourceItemCxtRelReq `False,1,256`
}

func (ie *PDUSessionResourceListCxtRelReq) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceItemCxtRelReq](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceListCxtRelReq) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceItemCxtRelReq
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceItemCxtRelReq](func() *PDUSessionResourceItemCxtRelReq { return new(PDUSessionResourceItemCxtRelReq) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceItemCxtRelReq{}
	ie.Value = append(ie.Value, newItems...)
	return
}
