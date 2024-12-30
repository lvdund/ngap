package ies

import "github.com/lvdund/ngap/aper"

const (
	CellIDListForRestartPresentNothing uint64 = iota
	CellIDListForRestartPresentEutraCgilistforrestart
	CellIDListForRestartPresentNrCgilistforrestart
	CellIDListForRestartPresentChoiceExtensions
)

type CellIDListForRestart struct {
	Choice                 uint64
	EUTRACGIListforRestart *EUTRACGI
	NRCGIListforRestart    *NRCGI
	// ChoiceExtensions *CellIDListForRestartExtIEs
}

func (ie *CellIDListForRestart) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case CellIDListForRestartPresentEutraCgilistforrestart:
		err = ie.EUTRACGIListforRestart.Encode(w)
	case CellIDListForRestartPresentNrCgilistforrestart:
		err = ie.NRCGIListforRestart.Encode(w)
	}
	return
}
func (ie *CellIDListForRestart) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case CellIDListForRestartPresentEutraCgilistforrestart:
		var tmp EUTRACGI
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRACGIListforRestart = &tmp
	case CellIDListForRestartPresentNrCgilistforrestart:
		var tmp NRCGI
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NRCGIListforRestart = &tmp
	}
	return
}
