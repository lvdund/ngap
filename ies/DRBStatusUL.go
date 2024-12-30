package ies

import "github.com/lvdund/ngap/aper"

const (
	DRBStatusULPresentNothing uint64 = iota
	DRBStatusULPresentDrbstatusul12
	DRBStatusULPresentDrbstatusul18
	DRBStatusULPresentChoiceExtensions
)

type DRBStatusUL struct {
	Choice        uint64
	DRBStatusUL12 *DRBStatusUL12
	DRBStatusUL18 *DRBStatusUL18
	// ChoiceExtensions *DRBStatusULExtIEs
}

func (ie *DRBStatusUL) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case DRBStatusULPresentDrbstatusul12:
		err = ie.DRBStatusUL12.Encode(w)
	case DRBStatusULPresentDrbstatusul18:
		err = ie.DRBStatusUL18.Encode(w)
	}
	return
}
func (ie *DRBStatusUL) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case DRBStatusULPresentDrbstatusul12:
		var tmp DRBStatusUL12
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.DRBStatusUL12 = &tmp
	case DRBStatusULPresentDrbstatusul18:
		var tmp DRBStatusUL18
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.DRBStatusUL18 = &tmp
	}
	return
}
