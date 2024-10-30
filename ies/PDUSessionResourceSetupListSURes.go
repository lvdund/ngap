package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceSetupListSURes struct {
	Value []*PDUSessionResourceSetupItemSURes `False,1,256`
}

func (ie *PDUSessionResourceSetupListSURes) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceSetupItemSURes](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceSetupListSURes) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceSetupItemSURes
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceSetupItemSURes](func() *PDUSessionResourceSetupItemSURes { return new(PDUSessionResourceSetupItemSURes) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceSetupItemSURes{}
	ie.Value = append(ie.Value, newItems...)
	return
}
