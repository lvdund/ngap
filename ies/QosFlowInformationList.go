package ies

import "github.com/lvdund/ngap/aper"

type QosFlowInformationList struct {
	Value []*QosFlowInformationItem `False,1,64`
}

func (ie *QosFlowInformationList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*QosFlowInformationItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return
	}
	return
}
func (ie *QosFlowInformationList) Decode(r *aper.AperReader) (err error) {
	var newItems []*QosFlowInformationItem
	newItems, err = aper.ReadSequenceOfEx[*QosFlowInformationItem](func() *QosFlowInformationItem { return new(QosFlowInformationItem) }, r, &aper.Constraint{Lb: 1, Ub: 64}, false)
	ie.Value = []*QosFlowInformationItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
