package ies

import "github.com/lvdund/ngap/aper"

type OverloadResponse struct {
	Choice         uint64
	OverloadAction *OverloadAction `False,,,`
	// ChoiceExtensions *OverloadResponseExtIEs `False,,,`
}

func (ie *OverloadResponse) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.OverloadAction.Encode(w)
	}
	return
}
func (ie *OverloadResponse) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp OverloadAction
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.OverloadAction = &tmp
	}
	return
}
