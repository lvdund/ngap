package ies

import "github.com/lvdund/ngap/aper"

const (
	NGRANCGIPresentNothing uint64 = iota /* No components present */
	NGRANCGIPresentNRCGI
	NGRANCGIPresentEUTRACGI
	NGRANCGIPresentChoiceExtensions
)

type NGRANCGI struct {
	Choice   uint64
	NRCGI    *NRCGI    `True,,,`
	EUTRACGI *EUTRACGI `True,,,`
	// ChoiceExtensions *NGRANCGIExtIEs `False,,,`
}

func (ie *NGRANCGI) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case NGRANCGIPresentNRCGI:
		err = ie.NRCGI.Encode(w)
	case NGRANCGIPresentEUTRACGI:
		err = ie.EUTRACGI.Encode(w)
	}
	return
}
func (ie *NGRANCGI) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case NGRANCGIPresentNRCGI:
		var tmp NRCGI
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NRCGI = &tmp
	case NGRANCGIPresentEUTRACGI:
		var tmp EUTRACGI
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRACGI = &tmp
	}
	return
}
