package ies

import "github.com/lvdund/ngap/aper"

const (
	UENGAPIDsPresentNothing uint64 = iota /* No components present */
	UENGAPIDsPresentUENGAPIDPair
	UENGAPIDsPresentAMFUENGAPID
	UENGAPIDsPresentChoiceExtensions
)

type UENGAPIDs struct {
	Choice       uint64
	UENGAPIDpair *UENGAPIDpair `True,,,`
	AMFUENGAPID  *AMFUENGAPID  `False,,,`
	// ChoiceExtensions *UENGAPIDsExtIEs `False,,,`
}

func (ie *UENGAPIDs) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case UENGAPIDsPresentUENGAPIDPair:
		err = ie.UENGAPIDpair.Encode(w)
	case UENGAPIDsPresentAMFUENGAPID:
		err = ie.AMFUENGAPID.Encode(w)
	}
	return
}
func (ie *UENGAPIDs) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case UENGAPIDsPresentUENGAPIDPair:
		var tmp UENGAPIDpair
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.UENGAPIDpair = &tmp
	case UENGAPIDsPresentAMFUENGAPID:
		var tmp AMFUENGAPID
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.AMFUENGAPID = &tmp
	}
	return
}
