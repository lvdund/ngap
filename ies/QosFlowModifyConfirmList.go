package ies

import "github.com/lvdund/ngap/aper"

type QosFlowModifyConfirmList struct {
	Value []*QosFlowModifyConfirmItem `False,1,64`
}

func (ie *QosFlowModifyConfirmList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*QosFlowModifyConfirmItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return
	}
	return
}
func (ie *QosFlowModifyConfirmList) Decode(r *aper.AperReader) (err error) {
	var newItems []*QosFlowModifyConfirmItem
	newItems, err = aper.ReadSequenceOfEx[*QosFlowModifyConfirmItem](func() *QosFlowModifyConfirmItem { return new(QosFlowModifyConfirmItem) }, r, &aper.Constraint{Lb: 1, Ub: 64}, false)
	ie.Value = []*QosFlowModifyConfirmItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
