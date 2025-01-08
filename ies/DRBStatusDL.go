package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	DRBStatusDLPresentNothing uint64 = iota
	DRBStatusDLPresentDrbstatusdl12
	DRBStatusDLPresentDrbstatusdl18
	DRBStatusDLPresentChoiceExtensions
)

type DRBStatusDL struct {
	Choice        uint64
	DRBStatusDL12 *DRBStatusDL12
	DRBStatusDL18 *DRBStatusDL18
	// ChoiceExtensions *DRBStatusDLExtIEs
}

func (ie *DRBStatusDL) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case DRBStatusDLPresentDrbstatusdl12:
		err = ie.DRBStatusDL12.Encode(w)
	case DRBStatusDLPresentDrbstatusdl18:
		err = ie.DRBStatusDL18.Encode(w)
	}
	return
}
func (ie *DRBStatusDL) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case DRBStatusDLPresentDrbstatusdl12:
		var tmp DRBStatusDL12
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read DRBStatusDL12", err)
			return
		}
		ie.DRBStatusDL12 = &tmp
	case DRBStatusDLPresentDrbstatusdl18:
		var tmp DRBStatusDL18
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read DRBStatusDL18", err)
			return
		}
		ie.DRBStatusDL18 = &tmp
	}
	return
}
