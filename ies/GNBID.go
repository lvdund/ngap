package ies

import "github.com/lvdund/ngap/aper"

const (
	GNBIDPresentNothing uint64 = iota
	GNBIDPresentGnbId
	GNBIDPresentChoiceExtensions
)

type GNBID struct {
	Choice uint64
	GNBID  *[]byte
	// ChoiceExtensions *GNBIDExtIEs
}

func (ie *GNBID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case GNBIDPresentGnbId:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
			Value: aper.BitString{
				Bytes:   *ie.GNBID,
				NumBits: uint64(len(*ie.GNBID)),
			},
		}
		err = tmp.Encode(w)
	}
	return
}
func (ie *GNBID) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case GNBIDPresentGnbId:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GNBID = &tmp.Value.Bytes
	}
	return
}
