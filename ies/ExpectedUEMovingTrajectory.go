package ies

import "github.com/lvdund/ngap/aper"

type ExpectedUEMovingTrajectory struct {
	Value []*ExpectedUEMovingTrajectoryItem `False,1,16`
}

func (ie *ExpectedUEMovingTrajectory) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*ExpectedUEMovingTrajectoryItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return
	}
	return
}
func (ie *ExpectedUEMovingTrajectory) Decode(r *aper.AperReader) (err error) {
	var newItems []*ExpectedUEMovingTrajectoryItem
	newItems, err = aper.ReadSequenceOfEx[*ExpectedUEMovingTrajectoryItem](func() *ExpectedUEMovingTrajectoryItem { return new(ExpectedUEMovingTrajectoryItem) }, r, &aper.Constraint{Lb: 1, Ub: 16}, false)
	ie.Value = []*ExpectedUEMovingTrajectoryItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
