package ies

import "github.com/lvdund/ngap/aper"

type N3IWFID struct {
	Choice  uint64
	N3IWFID *aper.BitString `False,,16,16`
	// ChoiceExtensions *N3IWFIDExtIEs `False,,,`
}

func (ie *N3IWFID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = w.WriteBitString(ie.N3IWFID.Bytes, uint(ie.N3IWFID.NumBits), &aper.Constraint{Lb: 16, Ub: 16}, false)
	}
	return
}
func (ie *N3IWFID) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var b []byte
		var n uint
		if b, n, err = r.ReadBitString(&aper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return
		}
		ie.N3IWFID = &aper.BitString{Bytes: b, NumBits: uint64(n)}
	}
	return
}
