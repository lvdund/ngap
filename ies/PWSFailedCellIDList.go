package ies

import "github.com/lvdund/ngap/aper"

const (
	PWSFailedCellIDListPresentNothing uint64 = iota
	PWSFailedCellIDListPresentEutraCgiPwsfailedlist
	PWSFailedCellIDListPresentNrCgiPwsfailedlist
	PWSFailedCellIDListPresentChoiceExtensions
)

type PWSFailedCellIDList struct {
	Choice                uint64
	EUTRACGIPWSFailedList *EUTRACGI
	NRCGIPWSFailedList    *NRCGI
	// ChoiceExtensions *PWSFailedCellIDListExtIEs
}

func (ie *PWSFailedCellIDList) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case PWSFailedCellIDListPresentEutraCgiPwsfailedlist:
		err = ie.EUTRACGIPWSFailedList.Encode(w)
	case PWSFailedCellIDListPresentNrCgiPwsfailedlist:
		err = ie.NRCGIPWSFailedList.Encode(w)
	}
	return
}
func (ie *PWSFailedCellIDList) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case PWSFailedCellIDListPresentEutraCgiPwsfailedlist:
		var tmp EUTRACGI
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRACGIPWSFailedList = &tmp
	case PWSFailedCellIDListPresentNrCgiPwsfailedlist:
		var tmp NRCGI
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NRCGIPWSFailedList = &tmp
	}
	return
}
