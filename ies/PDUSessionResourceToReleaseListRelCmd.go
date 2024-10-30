package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceToReleaseListRelCmd struct {
	Value []*PDUSessionResourceToReleaseItemRelCmd `False,1,256`
}

func (ie *PDUSessionResourceToReleaseListRelCmd) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceToReleaseItemRelCmd](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceToReleaseListRelCmd) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceToReleaseItemRelCmd
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceToReleaseItemRelCmd](func() *PDUSessionResourceToReleaseItemRelCmd { return new(PDUSessionResourceToReleaseItemRelCmd) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceToReleaseItemRelCmd{}
	ie.Value = append(ie.Value, newItems...)
	return
}
