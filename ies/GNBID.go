package ies

import "github.com/lvdund/ngap/aper"

type GNBID struct {
	Choice uint64
	GNBID  *aper.BitString `False,,22,32`
	// ChoiceExtensions *GNBIDExtIEs `False,,,`
}

func (ie *GNBID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = w.WriteBitString(ie.GNBID.Bytes, uint(ie.GNBID.NumBits), &aper.Constraint{Lb: 22, Ub: 32}, false)
	}
	return
}
func (ie *GNBID) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var b []byte
		var n uint
		if b, n, err = r.ReadBitString(&aper.Constraint{Lb: 22, Ub: 32}, false); err != nil {
			return
		}
		ie.GNBID = &aper.BitString{Bytes: b, NumBits: uint64(n)}
	}
	return
}
