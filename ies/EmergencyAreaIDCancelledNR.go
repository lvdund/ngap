package ies

import "github.com/lvdund/ngap/aper"

type EmergencyAreaIDCancelledNR struct {
	Value []*EmergencyAreaIDCancelledNRItem `False,1,65535`
}

func (ie *EmergencyAreaIDCancelledNR) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*EmergencyAreaIDCancelledNRItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *EmergencyAreaIDCancelledNR) Decode(r *aper.AperReader) (err error) {
	var newItems []*EmergencyAreaIDCancelledNRItem
	newItems, err = aper.ReadSequenceOfEx[*EmergencyAreaIDCancelledNRItem](func() *EmergencyAreaIDCancelledNRItem { return new(EmergencyAreaIDCancelledNRItem) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*EmergencyAreaIDCancelledNRItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
