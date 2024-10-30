package ies

import "github.com/lvdund/ngap/aper"

type QosFlowNotifyList struct {
	Value []*QosFlowNotifyItem `False,1,64`
}

func (ie *QosFlowNotifyList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*QosFlowNotifyItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return
	}
	return
}
func (ie *QosFlowNotifyList) Decode(r *aper.AperReader) (err error) {
	var newItems []*QosFlowNotifyItem
	newItems, err = aper.ReadSequenceOfEx[*QosFlowNotifyItem](func() *QosFlowNotifyItem { return new(QosFlowNotifyItem) }, r, &aper.Constraint{Lb: 1, Ub: 64}, false)
	ie.Value = []*QosFlowNotifyItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
