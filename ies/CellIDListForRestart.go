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
	EUTRACGIListforRestart *[]EUTRACGI
	NRCGIListforRestart    *[]NRCGI
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
		for _, i := range *ie.EUTRACGIListforRestart {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case CellIDListForRestartPresentNrCgilistforrestart:
		tmp := Sequence[*NRCGI]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellsingNB},
			ext: false,
		}
		for _, i := range *ie.NRCGIListforRestart {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	}
	return
}
func (ie *CellIDListForRestart) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case CellIDListForRestartPresentEutraCgilistforrestart:
		tmp := Sequence[*EUTRACGI]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellsinngeNB},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRACGIListforRestart = &[]EUTRACGI{}
		for _, i := range tmp.Value {
			*ie.EUTRACGIListforRestart = append(*ie.EUTRACGIListforRestart, *i)
		}
	case CellIDListForRestartPresentNrCgilistforrestart:
		tmp := Sequence[*NRCGI]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellsingNB},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NRCGIListforRestart = &[]NRCGI{}
		for _, i := range tmp.Value {
			*ie.NRCGIListforRestart = append(*ie.NRCGIListforRestart, *i)
		}
	}
	return
}
