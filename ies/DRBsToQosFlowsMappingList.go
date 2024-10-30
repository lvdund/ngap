package ies

import "github.com/lvdund/ngap/aper"

type DRBsToQosFlowsMappingList struct {
	Value []*DRBsToQosFlowsMappingItem `False,1,32`
}

func (ie *DRBsToQosFlowsMappingList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*DRBsToQosFlowsMappingItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return
	}
	return
}
func (ie *DRBsToQosFlowsMappingList) Decode(r *aper.AperReader) (err error) {
	var newItems []*DRBsToQosFlowsMappingItem
	newItems, err = aper.ReadSequenceOfEx[*DRBsToQosFlowsMappingItem](func() *DRBsToQosFlowsMappingItem { return new(DRBsToQosFlowsMappingItem) }, r, &aper.Constraint{Lb: 1, Ub: 32}, false)
	ie.Value = []*DRBsToQosFlowsMappingItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
