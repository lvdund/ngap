package ies

import "github.com/lvdund/ngap/aper"

type UEPresenceInAreaOfInterestList struct {
	Value []*UEPresenceInAreaOfInterestItem `False,1,64`
}

func (ie *UEPresenceInAreaOfInterestList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*UEPresenceInAreaOfInterestItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return
	}
	return
}
func (ie *UEPresenceInAreaOfInterestList) Decode(r *aper.AperReader) (err error) {
	var newItems []*UEPresenceInAreaOfInterestItem
	newItems, err = aper.ReadSequenceOfEx[*UEPresenceInAreaOfInterestItem](func() *UEPresenceInAreaOfInterestItem { return new(UEPresenceInAreaOfInterestItem) }, r, &aper.Constraint{Lb: 1, Ub: 64}, false)
	ie.Value = []*UEPresenceInAreaOfInterestItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
