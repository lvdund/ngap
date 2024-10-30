package ies

import "github.com/lvdund/ngap/aper"

type DRBStatusUL struct {
	Choice        uint64
	DRBStatusUL12 *DRBStatusUL12 `False,,,`
	DRBStatusUL18 *DRBStatusUL18 `False,,,`
	// ChoiceExtensions *DRBStatusULExtIEs `False,,,`
}

func (ie *DRBStatusUL) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.DRBStatusUL12.Encode(w)
	case 2:
		err = ie.DRBStatusUL18.Encode(w)
	}
	return
}
func (ie *DRBStatusUL) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp DRBStatusUL12
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.DRBStatusUL12 = &tmp
	case 2:
		var tmp DRBStatusUL18
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.DRBStatusUL18 = &tmp
	}
	return
}
