package ies

import "github.com/lvdund/ngap/aper"

type UEPagingIdentity struct {
	Choice     uint64
	FiveGSTMSI *FiveGSTMSI `True,,,`
	// ChoiceExtensions *UEPagingIdentityExtIEs `False,,,`
}

func (ie *UEPagingIdentity) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.FiveGSTMSI.Encode(w)
	}
	return
}
func (ie *UEPagingIdentity) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp FiveGSTMSI
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.FiveGSTMSI = &tmp
	}
	return
}
