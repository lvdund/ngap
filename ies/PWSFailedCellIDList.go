package ies

import "github.com/lvdund/ngap/aper"

type PWSFailedCellIDList struct {
	Choice                uint64
	EUTRACGIPWSFailedList *EUTRACGIList `False,,,`
	NRCGIPWSFailedList    *NRCGIList    `False,,,`
	// ChoiceExtensions *PWSFailedCellIDListExtIEs `False,,,`
}

func (ie *PWSFailedCellIDList) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.EUTRACGIPWSFailedList.Encode(w)
	case 2:
		err = ie.NRCGIPWSFailedList.Encode(w)
	}
	return
}
func (ie *PWSFailedCellIDList) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp EUTRACGIList
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRACGIPWSFailedList = &tmp
	case 2:
		var tmp NRCGIList
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NRCGIPWSFailedList = &tmp
	}
	return
}
