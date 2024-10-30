package ies

import "github.com/lvdund/ngap/aper"

type QosFlowPerTNLInformationList struct {
	Value []*QosFlowPerTNLInformationItem `False,1,3`
}

func (ie *QosFlowPerTNLInformationList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*QosFlowPerTNLInformationItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 3}, false); err != nil {
		return
	}
	return
}
func (ie *QosFlowPerTNLInformationList) Decode(r *aper.AperReader) (err error) {
	var newItems []*QosFlowPerTNLInformationItem
	newItems, err = aper.ReadSequenceOfEx[*QosFlowPerTNLInformationItem](func() *QosFlowPerTNLInformationItem { return new(QosFlowPerTNLInformationItem) }, r, &aper.Constraint{Lb: 1, Ub: 3}, false)
	ie.Value = []*QosFlowPerTNLInformationItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
