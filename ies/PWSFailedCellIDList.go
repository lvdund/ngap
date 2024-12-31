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
	EUTRACGIPWSFailedList *[]EUTRACGI
	NRCGIPWSFailedList    *[]NRCGI
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
		for _, i := range *ie.EUTRACGIPWSFailedList {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case PWSFailedCellIDListPresentNrCgiPwsfailedlist:
		tmp := Sequence[*NRCGI]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellsingNB},
			ext: false,
		}
		for _, i := range *ie.NRCGIPWSFailedList {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	}
	return
}
func (ie *PWSFailedCellIDList) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case PWSFailedCellIDListPresentEutraCgiPwsfailedlist:
		tmp := Sequence[*EUTRACGI]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellsinngeNB},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRACGIPWSFailedList = &[]EUTRACGI{}
		for _, i := range tmp.Value {
			*ie.EUTRACGIPWSFailedList = append(*ie.EUTRACGIPWSFailedList, *i)
		}
	case PWSFailedCellIDListPresentNrCgiPwsfailedlist:
		tmp := Sequence[*NRCGI]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellsingNB},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NRCGIPWSFailedList = &[]NRCGI{}
		for _, i := range tmp.Value {
			*ie.NRCGIPWSFailedList = append(*ie.NRCGIPWSFailedList, *i)
		}
	}
	return
}
