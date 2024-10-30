package ies

import "github.com/lvdund/ngap/aper"

type QosFlowToBeForwardedList struct {
	Value []*QosFlowToBeForwardedItem `False,1,64`
}

func (ie *QosFlowToBeForwardedList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*QosFlowToBeForwardedItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return
	}
	return
}
func (ie *QosFlowToBeForwardedList) Decode(r *aper.AperReader) (err error) {
	var newItems []*QosFlowToBeForwardedItem
	newItems, err = aper.ReadSequenceOfEx[*QosFlowToBeForwardedItem](func() *QosFlowToBeForwardedItem { return new(QosFlowToBeForwardedItem) }, r, &aper.Constraint{Lb: 1, Ub: 64}, false)
	ie.Value = []*QosFlowToBeForwardedItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
