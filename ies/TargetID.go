package ies

import "github.com/lvdund/ngap/aper"

const (
	TargetIDPresentNothing uint64 = iota /* No components present */
	TargetIDPresentTargetRANNodeID
	TargetIDPresentTargeteNBID
	TargetIDPresentChoiceExtensions
)

type TargetID struct {
	Choice          uint64
	TargetRANNodeID *TargetRANNodeID `True,,,`
	TargeteNBID     *TargeteNBID     `True,,,`
	// ChoiceExtensions *TargetIDExtIEs `False,,,`
}

func (ie *TargetID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case TargetIDPresentTargetRANNodeID:
		err = ie.TargetRANNodeID.Encode(w)
	case TargetIDPresentTargeteNBID:
		err = ie.TargeteNBID.Encode(w)
	}
	return
}
func (ie *TargetID) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case TargetIDPresentTargetRANNodeID:
		var tmp TargetRANNodeID
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TargetRANNodeID = &tmp
	case TargetIDPresentTargeteNBID:
		var tmp TargeteNBID
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TargeteNBID = &tmp
	}
	return
}
