package ies

import "github.com/lvdund/ngap/aper"

type TargetID struct {
	Choice          uint64
	TargetRANNodeID *TargetRANNodeID `True,,,`
	TargeteNBID     *TargeteNBID     `True,,,`
	// ChoiceExtensions *TargetIDExtIEs `False,,,`
}

func (ie *TargetID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.TargetRANNodeID.Encode(w)
	case 2:
		err = ie.TargeteNBID.Encode(w)
	}
	return
}
func (ie *TargetID) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp TargetRANNodeID
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TargetRANNodeID = &tmp
	case 2:
		var tmp TargeteNBID
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TargeteNBID = &tmp
	}
	return
}
