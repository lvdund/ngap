package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	PWSFailedCellIDListPresentNothing uint64 = iota
	PWSFailedCellIDListPresentEutraCgiPwsfailedlist
	PWSFailedCellIDListPresentNrCgiPwsfailedlist
	PWSFailedCellIDListPresentChoiceExtensions
)

type PWSFailedCellIDList struct {
	Choice                uint64
	EUTRACGIPWSFailedList []EUTRACGI
	NRCGIPWSFailedList    []NRCGI
	// ChoiceExtensions *PWSFailedCellIDListExtIEs
}

func (ie *PWSFailedCellIDList) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case PWSFailedCellIDListPresentEutraCgiPwsfailedlist:
		tmp := Sequence[*EUTRACGI]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellsinngeNB},
			ext: false,
		}
		for _, i := range ie.EUTRACGIPWSFailedList {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case PWSFailedCellIDListPresentNrCgiPwsfailedlist:
		tmp := Sequence[*NRCGI]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellsingNB},
			ext: false,
		}
		for _, i := range ie.NRCGIPWSFailedList {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	}
	return
}
func (ie *PWSFailedCellIDList) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case PWSFailedCellIDListPresentEutraCgiPwsfailedlist:
		tmp := NewSequence[*EUTRACGI](nil, aper.Constraint{Lb: 1, Ub: maxnoofCellsinngeNB}, false)
		fn := func() *EUTRACGI {
			return new(EUTRACGI)
		}
		if err = tmp.Decode(r, fn); err != nil {
			err = utils.WrapError("Read EUTRACGIPWSFailedList", err)
			return
		}
		for _, i := range tmp.Value {
			ie.EUTRACGIPWSFailedList = append(ie.EUTRACGIPWSFailedList, *i)
		}
	case PWSFailedCellIDListPresentNrCgiPwsfailedlist:
		tmp := NewSequence[*NRCGI](nil, aper.Constraint{Lb: 1, Ub: maxnoofCellsingNB}, false)
		fn := func() *NRCGI {
			return new(NRCGI)
		}
		if err = tmp.Decode(r, fn); err != nil {
			err = utils.WrapError("Read NRCGIPWSFailedList", err)
			return
		}
		for _, i := range tmp.Value {
			ie.NRCGIPWSFailedList = append(ie.NRCGIPWSFailedList, *i)
		}
	}
	return
}
