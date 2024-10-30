package ies

import "github.com/lvdund/ngap/aper"

type CellIDListForRestart struct {
	Choice                 uint64
	EUTRACGIListforRestart *EUTRACGIList `False,,,`
	NRCGIListforRestart    *NRCGIList    `False,,,`
	// ChoiceExtensions *CellIDListForRestartExtIEs `False,,,`
}

func (ie *CellIDListForRestart) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.EUTRACGIListforRestart.Encode(w)
	case 2:
		err = ie.NRCGIListforRestart.Encode(w)
	}
	return
}
func (ie *CellIDListForRestart) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp EUTRACGIList
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRACGIListforRestart = &tmp
	case 2:
		var tmp NRCGIList
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NRCGIListforRestart = &tmp
	}
	return
}
