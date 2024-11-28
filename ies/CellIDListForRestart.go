package ies

import "github.com/lvdund/ngap/aper"

const (
	CellIDListForRestartPresentNothing uint64 = iota /* No components present */
	CellIDListForRestartPresentEUTRACGIListforRestart
	CellIDListForRestartPresentNRCGIListforRestart
	CellIDListForRestartPresentChoiceExtensions
)

type CellIDListForRestart struct {
	Choice                 uint64
	EUTRACGIListforRestart *EUTRACGIList `False,,,`
	NRCGIListforRestart    *NRCGIList    `False,,,`
	// ChoiceExtensions *CellIDListForRestartExtIEs `False,,,`
}

func (ie *CellIDListForRestart) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case CellIDListForRestartPresentEUTRACGIListforRestart:
		err = ie.EUTRACGIListforRestart.Encode(w)
	case CellIDListForRestartPresentNRCGIListforRestart:
		err = ie.NRCGIListforRestart.Encode(w)
	}
	return
}
func (ie *CellIDListForRestart) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case CellIDListForRestartPresentEUTRACGIListforRestart:
		var tmp EUTRACGIList
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRACGIListforRestart = &tmp
	case CellIDListForRestartPresentNRCGIListforRestart:
		var tmp NRCGIList
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NRCGIListforRestart = &tmp
	}
	return
}
