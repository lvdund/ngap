package ies

import "github.com/lvdund/ngap/aper"

const (
	DRBStatusULPresentNothing uint64 = iota /* No components present */
	DRBStatusULPresentDRBStatusUL12
	DRBStatusULPresentDRBStatusUL18
	DRBStatusULPresentChoiceExtensions
)

type DRBStatusUL struct {
	Choice        uint64
	DRBStatusUL12 *DRBStatusUL12 `False,,,`
	DRBStatusUL18 *DRBStatusUL18 `False,,,`
	// ChoiceExtensions *DRBStatusULExtIEs `False,,,`
}

func (ie *DRBStatusUL) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case DRBStatusULPresentDRBStatusUL12:
		err = ie.DRBStatusUL12.Encode(w)
	case DRBStatusULPresentDRBStatusUL18:
		err = ie.DRBStatusUL18.Encode(w)
	}
	return
}
func (ie *DRBStatusUL) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case DRBStatusULPresentDRBStatusUL12:
		var tmp DRBStatusUL12
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.DRBStatusUL12 = &tmp
	case DRBStatusULPresentDRBStatusUL18:
		var tmp DRBStatusUL18
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.DRBStatusUL18 = &tmp
	}
	return
}
