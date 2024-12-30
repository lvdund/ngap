package ies

import "github.com/lvdund/ngap/aper"

const (
	UENGAPIDsPresentNothing uint64 = iota
	UENGAPIDsPresentUeNgapIdPair
	UENGAPIDsPresentAmfUeNgapId
	UENGAPIDsPresentChoiceExtensions
)

type UENGAPIDs struct {
	Choice       uint64
	UENGAPIDpair *UENGAPIDpair
	AMFUENGAPID  *INTEGER
	// ChoiceExtensions *UENGAPIDsExtIEs
}

func (ie *UENGAPIDs) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case UENGAPIDsPresentUeNgapIdPair:
		err = ie.UENGAPIDpair.Encode(w)
	case UENGAPIDsPresentAmfUeNgapId:
		err = ie.AMFUENGAPID.Encode(w)
	}
	return
}
func (ie *UENGAPIDs) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case UENGAPIDsPresentUeNgapIdPair:
		var tmp UENGAPIDpair
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.UENGAPIDpair = &tmp
	case UENGAPIDsPresentAmfUeNgapId:
		var tmp INTEGER
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.AMFUENGAPID = &tmp
	}
	return
}
