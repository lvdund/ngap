package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	AMFPagingTargetPresentNothing uint64 = iota
	AMFPagingTargetPresentGlobalrannodeid
	AMFPagingTargetPresentTai
	AMFPagingTargetPresentChoiceExtensions
)

type AMFPagingTarget struct {
	Choice          uint64
	GlobalRANNodeID *GlobalRANNodeID
	TAI             *TAI
	// ChoiceExtensions *AMFPagingTargetExtIEs
}

func (ie *AMFPagingTarget) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case AMFPagingTargetPresentGlobalrannodeid:
		err = ie.GlobalRANNodeID.Encode(w)
	case AMFPagingTargetPresentTai:
		err = ie.TAI.Encode(w)
	}
	return
}
func (ie *AMFPagingTarget) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case AMFPagingTargetPresentGlobalrannodeid:
		var tmp GlobalRANNodeID
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read GlobalRANNodeID", err)
			return
		}
		ie.GlobalRANNodeID = &tmp
	case AMFPagingTargetPresentTai:
		var tmp TAI
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read TAI", err)
			return
		}
		ie.TAI = &tmp
	}
	return
}
