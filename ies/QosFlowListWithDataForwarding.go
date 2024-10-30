package ies

import "github.com/lvdund/ngap/aper"

type QosFlowListWithDataForwarding struct {
	Value []*QosFlowItemWithDataForwarding `False,1,64`
}

func (ie *QosFlowListWithDataForwarding) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*QosFlowItemWithDataForwarding](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return
	}
	return
}
func (ie *QosFlowListWithDataForwarding) Decode(r *aper.AperReader) (err error) {
	var newItems []*QosFlowItemWithDataForwarding
	newItems, err = aper.ReadSequenceOfEx[*QosFlowItemWithDataForwarding](func() *QosFlowItemWithDataForwarding { return new(QosFlowItemWithDataForwarding) }, r, &aper.Constraint{Lb: 1, Ub: 64}, false)
	ie.Value = []*QosFlowItemWithDataForwarding{}
	ie.Value = append(ie.Value, newItems...)
	return
}
