package ies

import "github.com/lvdund/ngap/aper"

type NGRANCGI struct {
	Choice   uint64
	NRCGI    *NRCGI    `True,,,`
	EUTRACGI *EUTRACGI `True,,,`
	// ChoiceExtensions *NGRANCGIExtIEs `False,,,`
}

func (ie *NGRANCGI) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.NRCGI.Encode(w)
	case 2:
		err = ie.EUTRACGI.Encode(w)
	}
	return
}
func (ie *NGRANCGI) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp NRCGI
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NRCGI = &tmp
	case 2:
		var tmp EUTRACGI
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRACGI = &tmp
	}
	return
}
