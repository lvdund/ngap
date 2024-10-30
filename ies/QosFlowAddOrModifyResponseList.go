package ies

import "github.com/lvdund/ngap/aper"

type QosFlowAddOrModifyResponseList struct {
	Value []*QosFlowAddOrModifyResponseItem `False,1,64`
}

func (ie *QosFlowAddOrModifyResponseList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*QosFlowAddOrModifyResponseItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return
	}
	return
}
func (ie *QosFlowAddOrModifyResponseList) Decode(r *aper.AperReader) (err error) {
	var newItems []*QosFlowAddOrModifyResponseItem
	newItems, err = aper.ReadSequenceOfEx[*QosFlowAddOrModifyResponseItem](func() *QosFlowAddOrModifyResponseItem { return new(QosFlowAddOrModifyResponseItem) }, r, &aper.Constraint{Lb: 1, Ub: 64}, false)
	ie.Value = []*QosFlowAddOrModifyResponseItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
