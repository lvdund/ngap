package ies

import "github.com/lvdund/ngap/aper"

const (
	OverloadResponsePresentNothing uint64 = iota
	OverloadResponsePresentOverloadaction
	OverloadResponsePresentChoiceExtensions
)

type OverloadResponse struct {
	Choice         uint64
	OverloadAction *OverloadAction
	// ChoiceExtensions *OverloadResponseExtIEs
}

func (ie *OverloadResponse) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case OverloadResponsePresentOverloadaction:
		err = ie.OverloadAction.Encode(w)
	}
	return
}
func (ie *OverloadResponse) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return
	}
	switch ie.Choice {
	case OverloadResponsePresentOverloadaction:
		var tmp OverloadAction
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.OverloadAction = &tmp
	}
	return
}
