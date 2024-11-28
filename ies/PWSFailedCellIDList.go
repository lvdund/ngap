package ies

import "github.com/lvdund/ngap/aper"

const (
	PWSFailedCellIDListPresentNothing uint64 = iota /* No components present */
	PWSFailedCellIDListPresentEUTRACGIPWSFailedList
	PWSFailedCellIDListPresentNRCGIPWSFailedList
	PWSFailedCellIDListPresentChoiceExtensions
)

type PWSFailedCellIDList struct {
	Choice                uint64
	EUTRACGIPWSFailedList *EUTRACGIList `False,,,`
	NRCGIPWSFailedList    *NRCGIList    `False,,,`
	// ChoiceExtensions *PWSFailedCellIDListExtIEs `False,,,`
}

func (ie *PWSFailedCellIDList) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case PWSFailedCellIDListPresentEUTRACGIPWSFailedList:
		err = ie.EUTRACGIPWSFailedList.Encode(w)
	case PWSFailedCellIDListPresentNRCGIPWSFailedList:
		err = ie.NRCGIPWSFailedList.Encode(w)
	}
	return
}
func (ie *PWSFailedCellIDList) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case PWSFailedCellIDListPresentEUTRACGIPWSFailedList:
		var tmp EUTRACGIList
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRACGIPWSFailedList = &tmp
	case PWSFailedCellIDListPresentNRCGIPWSFailedList:
		var tmp NRCGIList
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NRCGIPWSFailedList = &tmp
	}
	return
}
