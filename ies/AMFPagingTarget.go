package ies

import "github.com/lvdund/ngap/aper"

type AMFPagingTarget struct {
	Choice          uint64
	GlobalRANNodeID *GlobalRANNodeID `False,,,`
	TAI             *TAI             `True,,,`
	// ChoiceExtensions *AMFPagingTargetExtIEs `False,,,`
}

func (ie *AMFPagingTarget) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.GlobalRANNodeID.Encode(w)
	case 2:
		err = ie.TAI.Encode(w)
	}
	return
}
func (ie *AMFPagingTarget) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp GlobalRANNodeID
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GlobalRANNodeID = &tmp
	case 2:
		var tmp TAI
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TAI = &tmp
	}
	return
}
