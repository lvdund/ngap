package ies

import "github.com/lvdund/ngap/aper"

type EmergencyAreaIDCancelledEUTRA struct {
	Value []*EmergencyAreaIDCancelledEUTRAItem `False,1,65535`
}

func (ie *EmergencyAreaIDCancelledEUTRA) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*EmergencyAreaIDCancelledEUTRAItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *EmergencyAreaIDCancelledEUTRA) Decode(r *aper.AperReader) (err error) {
	var newItems []*EmergencyAreaIDCancelledEUTRAItem
	newItems, err = aper.ReadSequenceOfEx[*EmergencyAreaIDCancelledEUTRAItem](func() *EmergencyAreaIDCancelledEUTRAItem { return new(EmergencyAreaIDCancelledEUTRAItem) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*EmergencyAreaIDCancelledEUTRAItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
