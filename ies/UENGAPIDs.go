package ies

import "github.com/lvdund/ngap/aper"

type UENGAPIDs struct {
	Choice       uint64
	UENGAPIDpair *UENGAPIDpair `True,,,`
	AMFUENGAPID  *AMFUENGAPID  `False,,,`
	// ChoiceExtensions *UENGAPIDsExtIEs `False,,,`
}

func (ie *UENGAPIDs) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.UENGAPIDpair.Encode(w)
	case 2:
		err = ie.AMFUENGAPID.Encode(w)
	}
	return
}
func (ie *UENGAPIDs) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp UENGAPIDpair
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.UENGAPIDpair = &tmp
	case 2:
		var tmp AMFUENGAPID
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.AMFUENGAPID = &tmp
	}
	return
}
