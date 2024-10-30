package ies

import "github.com/lvdund/ngap/aper"

type QosFlowSetupResponseListSURes struct {
	Value []*QosFlowSetupResponseItemSURes `False,1,64`
}

func (ie *QosFlowSetupResponseListSURes) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*QosFlowSetupResponseItemSURes](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return
	}
	return
}
func (ie *QosFlowSetupResponseListSURes) Decode(r *aper.AperReader) (err error) {
	var newItems []*QosFlowSetupResponseItemSURes
	newItems, err = aper.ReadSequenceOfEx[*QosFlowSetupResponseItemSURes](func() *QosFlowSetupResponseItemSURes { return new(QosFlowSetupResponseItemSURes) }, r, &aper.Constraint{Lb: 1, Ub: 64}, false)
	ie.Value = []*QosFlowSetupResponseItemSURes{}
	ie.Value = append(ie.Value, newItems...)
	return
}
