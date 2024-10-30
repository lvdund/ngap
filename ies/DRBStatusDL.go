package ies

import "github.com/lvdund/ngap/aper"

type DRBStatusDL struct {
	Choice        uint64
	DRBStatusDL12 *DRBStatusDL12 `False,,,`
	DRBStatusDL18 *DRBStatusDL18 `False,,,`
	// ChoiceExtensions *DRBStatusDLExtIEs `False,,,`
}

func (ie *DRBStatusDL) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.DRBStatusDL12.Encode(w)
	case 2:
		err = ie.DRBStatusDL18.Encode(w)
	}
	return
}
func (ie *DRBStatusDL) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp DRBStatusDL12
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.DRBStatusDL12 = &tmp
	case 2:
		var tmp DRBStatusDL18
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.DRBStatusDL18 = &tmp
	}
	return
}
