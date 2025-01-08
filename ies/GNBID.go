package ies

import "github.com/lvdund/ngap/aper"

const (
	GNBIDPresentNothing uint64 = iota
	GNBIDPresentGnbId
	GNBIDPresentChoiceExtensions
)

type GNBID struct {
	Choice uint64
	GNBID  []byte
	// ChoiceExtensions *GNBIDExtIEs
}

func (ie *GNBID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case GNBIDPresentGnbId:
		tmp := NewBITSTRING(ie.GNBID, aper.Constraint{Lb: 22, Ub: 32}, false)
		err = tmp.Encode(w)
	}
	return
}
func (ie *GNBID) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return
	}
	switch ie.Choice {
	case GNBIDPresentGnbId:
		tmp := NewBITSTRING(nil, aper.Constraint{Lb: 22, Ub: 32}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GNBID = tmp.Value.Bytes
	}
	return
}
