package ies

import "github.com/lvdund/ngap/aper"

type QosFlowAcceptedList struct {
	Value []*QosFlowAcceptedItem `False,1,64`
}

func (ie *QosFlowAcceptedList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*QosFlowAcceptedItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return
	}
	return
}
func (ie *QosFlowAcceptedList) Decode(r *aper.AperReader) (err error) {
	var newItems []*QosFlowAcceptedItem
	newItems, err = aper.ReadSequenceOfEx[*QosFlowAcceptedItem](func() *QosFlowAcceptedItem { return new(QosFlowAcceptedItem) }, r, &aper.Constraint{Lb: 1, Ub: 64}, false)
	ie.Value = []*QosFlowAcceptedItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
