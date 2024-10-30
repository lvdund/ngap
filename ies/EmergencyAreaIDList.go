package ies

import "github.com/lvdund/ngap/aper"

type EmergencyAreaIDList struct {
	Value []*EmergencyAreaID `False,1,65535`
}

func (ie *EmergencyAreaIDList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*EmergencyAreaID](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *EmergencyAreaIDList) Decode(r *aper.AperReader) (err error) {
	var newItems []*EmergencyAreaID
	newItems, err = aper.ReadSequenceOfEx[*EmergencyAreaID](func() *EmergencyAreaID { return new(EmergencyAreaID) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*EmergencyAreaID{}
	ie.Value = append(ie.Value, newItems...)
	return
}
