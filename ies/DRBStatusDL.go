package ies

import "github.com/lvdund/ngap/aper"

const (
	DRBStatusDLPresentNothing uint64 = iota /* No components present */
	DRBStatusDLPresentDRBStatusDL12
	DRBStatusDLPresentDRBStatusDL18
	DRBStatusDLPresentChoiceExtensions
)

type DRBStatusDL struct {
	Choice        uint64
	DRBStatusDL12 *DRBStatusDL12 `False,,,`
	DRBStatusDL18 *DRBStatusDL18 `False,,,`
	// ChoiceExtensions *DRBStatusDLExtIEs `False,,,`
}

func (ie *DRBStatusDL) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case DRBStatusDLPresentDRBStatusDL12:
		err = ie.DRBStatusDL12.Encode(w)
	case DRBStatusDLPresentDRBStatusDL18:
		err = ie.DRBStatusDL18.Encode(w)
	}
	return
}
func (ie *DRBStatusDL) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case DRBStatusDLPresentDRBStatusDL12:
		var tmp DRBStatusDL12
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.DRBStatusDL12 = &tmp
	case DRBStatusDLPresentDRBStatusDL18:
		var tmp DRBStatusDL18
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.DRBStatusDL18 = &tmp
	}
	return
}
