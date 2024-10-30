package ies

import "github.com/lvdund/ngap/aper"

type EmergencyAreaIDListForRestart struct {
	Value []*EmergencyAreaID `False,1,256`
}

func (ie *EmergencyAreaIDListForRestart) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*EmergencyAreaID](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *EmergencyAreaIDListForRestart) Decode(r *aper.AperReader) (err error) {
	var newItems []*EmergencyAreaID
	newItems, err = aper.ReadSequenceOfEx[*EmergencyAreaID](func() *EmergencyAreaID { return new(EmergencyAreaID) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*EmergencyAreaID{}
	ie.Value = append(ie.Value, newItems...)
	return
}
