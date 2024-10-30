package ies

import "github.com/lvdund/ngap/aper"

type QoSFlowsUsageReportList struct {
	Value []*QoSFlowsUsageReportItem `False,1,64`
}

func (ie *QoSFlowsUsageReportList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*QoSFlowsUsageReportItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return
	}
	return
}
func (ie *QoSFlowsUsageReportList) Decode(r *aper.AperReader) (err error) {
	var newItems []*QoSFlowsUsageReportItem
	newItems, err = aper.ReadSequenceOfEx[*QoSFlowsUsageReportItem](func() *QoSFlowsUsageReportItem { return new(QoSFlowsUsageReportItem) }, r, &aper.Constraint{Lb: 1, Ub: 64}, false)
	ie.Value = []*QoSFlowsUsageReportItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
