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
	AMFUENGAPID  *int64
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
		tmp := INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 1099511627775},
			ext:   false,
			Value: aper.Integer(*ie.AMFUENGAPID),
		}
		err = tmp.Encode(w)
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
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 1099511627775},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		*ie.AMFUENGAPID = int64(tmp.Value)
	}
	return
}
