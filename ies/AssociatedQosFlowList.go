package ies

import "github.com/lvdund/ngap/aper"

type AssociatedQosFlowList struct {
	Value []*AssociatedQosFlowItem `False,1,64`
}

func (ie *AssociatedQosFlowList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*AssociatedQosFlowItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return
	}
	return
}
func (ie *AssociatedQosFlowList) Decode(r *aper.AperReader) (err error) {
	var newItems []*AssociatedQosFlowItem
	newItems, err = aper.ReadSequenceOfEx[*AssociatedQosFlowItem](func() *AssociatedQosFlowItem { return new(AssociatedQosFlowItem) }, r, &aper.Constraint{Lb: 1, Ub: 64}, false)
	ie.Value = []*AssociatedQosFlowItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
