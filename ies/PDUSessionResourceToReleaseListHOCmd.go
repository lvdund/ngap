package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceToReleaseListHOCmd struct {
	Value []*PDUSessionResourceToReleaseItemHOCmd `False,1,256`
}

func (ie *PDUSessionResourceToReleaseListHOCmd) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PDUSessionResourceToReleaseItemHOCmd](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceToReleaseListHOCmd) Decode(r *aper.AperReader) (err error) {
	var newItems []*PDUSessionResourceToReleaseItemHOCmd
	newItems, err = aper.ReadSequenceOfEx[*PDUSessionResourceToReleaseItemHOCmd](func() *PDUSessionResourceToReleaseItemHOCmd { return new(PDUSessionResourceToReleaseItemHOCmd) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*PDUSessionResourceToReleaseItemHOCmd{}
	ie.Value = append(ie.Value, newItems...)
	return
}
