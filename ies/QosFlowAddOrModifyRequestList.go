package ies

import "github.com/lvdund/ngap/aper"

type QosFlowAddOrModifyRequestList struct {
	Value []*QosFlowAddOrModifyRequestItem `False,1,64`
}

func (ie *QosFlowAddOrModifyRequestList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*QosFlowAddOrModifyRequestItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return
	}
	return
}
func (ie *QosFlowAddOrModifyRequestList) Decode(r *aper.AperReader) (err error) {
	var newItems []*QosFlowAddOrModifyRequestItem
	newItems, err = aper.ReadSequenceOfEx[*QosFlowAddOrModifyRequestItem](func() *QosFlowAddOrModifyRequestItem { return new(QosFlowAddOrModifyRequestItem) }, r, &aper.Constraint{Lb: 1, Ub: 64}, false)
	ie.Value = []*QosFlowAddOrModifyRequestItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
