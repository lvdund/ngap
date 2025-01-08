package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	CellIDListForRestartPresentNothing uint64 = iota
	CellIDListForRestartPresentEutraCgilistforrestart
	CellIDListForRestartPresentNrCgilistforrestart
	CellIDListForRestartPresentChoiceExtensions
)

type CellIDListForRestart struct {
	Choice                 uint64
	EUTRACGIListforRestart []EUTRACGI
	NRCGIListforRestart    []NRCGI
	// ChoiceExtensions *CellIDListForRestartExtIEs
}

func (ie *CellIDListForRestart) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case CellIDListForRestartPresentEutraCgilistforrestart:
		tmp := Sequence[*EUTRACGI]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellsinngeNB},
			ext: false,
		}
		for _, i := range ie.EUTRACGIListforRestart {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case CellIDListForRestartPresentNrCgilistforrestart:
		tmp := Sequence[*NRCGI]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellsingNB},
			ext: false,
		}
		for _, i := range ie.NRCGIListforRestart {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	}
	return
}
func (ie *CellIDListForRestart) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case CellIDListForRestartPresentEutraCgilistforrestart:
		tmp := NewSequence[*EUTRACGI](nil, aper.Constraint{Lb: 1, Ub: maxnoofCellsinngeNB}, false)
		fn := func() *EUTRACGI {
			return new(EUTRACGI)
		}
		if err = tmp.Decode(r, fn); err != nil {
			err = utils.WrapError("Read EUTRACGIListforRestart", err)
			return
		}
		for _, i := range tmp.Value {
			ie.EUTRACGIListforRestart = append(ie.EUTRACGIListforRestart, *i)
		}
	case CellIDListForRestartPresentNrCgilistforrestart:
		tmp := NewSequence[*NRCGI](nil, aper.Constraint{Lb: 1, Ub: maxnoofCellsingNB}, false)
		fn := func() *NRCGI {
			return new(NRCGI)
		}
		if err = tmp.Decode(r, fn); err != nil {
			err = utils.WrapError("Read NRCGIListforRestart", err)
			return
		}
		for _, i := range tmp.Value {
			ie.NRCGIListforRestart = append(ie.NRCGIListforRestart, *i)
		}
	}
	return
}
