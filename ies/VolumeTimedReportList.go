package ies

import "github.com/lvdund/ngap/aper"

type VolumeTimedReportList struct {
	Value []*VolumeTimedReportItem `False,1,2`
}

func (ie *VolumeTimedReportList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*VolumeTimedReportItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return
	}
	return
}
func (ie *VolumeTimedReportList) Decode(r *aper.AperReader) (err error) {
	var newItems []*VolumeTimedReportItem
	newItems, err = aper.ReadSequenceOfEx[*VolumeTimedReportItem](func() *VolumeTimedReportItem { return new(VolumeTimedReportItem) }, r, &aper.Constraint{Lb: 1, Ub: 2}, false)
	ie.Value = []*VolumeTimedReportItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
