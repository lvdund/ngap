package ies

import "github.com/lvdund/ngap/aper"

const (
	NGRANCGIPresentNothing uint64 = iota
	NGRANCGIPresentNrCgi
	NGRANCGIPresentEutraCgi
	NGRANCGIPresentChoiceExtensions
)

type NGRANCGI struct {
	Choice   uint64
	NRCGI    *NRCGI
	EUTRACGI *EUTRACGI
	// ChoiceExtensions *NGRANCGIExtIEs
}

func (ie *NGRANCGI) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case NGRANCGIPresentNrCgi:
		err = ie.NRCGI.Encode(w)
	case NGRANCGIPresentEutraCgi:
		err = ie.EUTRACGI.Encode(w)
	}
	return
}
func (ie *NGRANCGI) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case NGRANCGIPresentNrCgi:
		var tmp NRCGI
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NRCGI = &tmp
	case NGRANCGIPresentEutraCgi:
		var tmp EUTRACGI
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRACGI = &tmp
	}
	return
}
