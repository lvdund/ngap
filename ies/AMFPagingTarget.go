package ies

import "github.com/lvdund/ngap/aper"

const (
	AMFPagingTargetPresentNothing uint64 = iota /* No components present */
	AMFPagingTargetPresentGlobalRANNodeID
	AMFPagingTargetPresentTAI
	AMFPagingTargetPresentChoiceExtensions
)

type AMFPagingTarget struct {
	Choice          uint64
	GlobalRANNodeID *GlobalRANNodeID `False,,,`
	TAI             *TAI             `True,,,`
	// ChoiceExtensions *AMFPagingTargetExtIEs `False,,,`
}

func (ie *AMFPagingTarget) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case AMFPagingTargetPresentGlobalRANNodeID:
		err = ie.GlobalRANNodeID.Encode(w)
	case AMFPagingTargetPresentTAI:
		err = ie.TAI.Encode(w)
	}
	return
}
func (ie *AMFPagingTarget) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case AMFPagingTargetPresentGlobalRANNodeID:
		var tmp GlobalRANNodeID
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GlobalRANNodeID = &tmp
	case AMFPagingTargetPresentTAI:
		var tmp TAI
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TAI = &tmp
	}
	return
}
