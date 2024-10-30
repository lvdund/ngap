package ies

import "github.com/lvdund/ngap/aper"

type QosFlowListWithCause struct {
	Value []*QosFlowWithCauseItem `False,1,64`
}

func (ie *QosFlowListWithCause) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*QosFlowWithCauseItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return
	}
	return
}
func (ie *QosFlowListWithCause) Decode(r *aper.AperReader) (err error) {
	var newItems []*QosFlowWithCauseItem
	newItems, err = aper.ReadSequenceOfEx[*QosFlowWithCauseItem](func() *QosFlowWithCauseItem { return new(QosFlowWithCauseItem) }, r, &aper.Constraint{Lb: 1, Ub: 64}, false)
	ie.Value = []*QosFlowWithCauseItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
