package ies

import "github.com/lvdund/ngap/aper"

type QosFlowSetupRequestList struct {
	Value []*QosFlowSetupRequestItem `False,1,64`
}

func (ie *QosFlowSetupRequestList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*QosFlowSetupRequestItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return
	}
	return
}
func (ie *QosFlowSetupRequestList) Decode(r *aper.AperReader) (err error) {
	var newItems []*QosFlowSetupRequestItem
	newItems, err = aper.ReadSequenceOfEx[*QosFlowSetupRequestItem](func() *QosFlowSetupRequestItem { return new(QosFlowSetupRequestItem) }, r, &aper.Constraint{Lb: 1, Ub: 64}, false)
	ie.Value = []*QosFlowSetupRequestItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
