package ies

import "github.com/lvdund/ngap/aper"

type NgENBID struct {
	Choice            uint64
	MacroNgENBID      *aper.BitString `False,,20,20`
	ShortMacroNgENBID *aper.BitString `False,,18,18`
	LongMacroNgENBID  *aper.BitString `False,,21,21`
	// ChoiceExtensions *NgENBIDExtIEs `False,,,`
}

func (ie *NgENBID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = w.WriteBitString(ie.MacroNgENBID.Bytes, uint(ie.MacroNgENBID.NumBits), &aper.Constraint{Lb: 20, Ub: 20}, false)
	case 2:
		err = w.WriteBitString(ie.ShortMacroNgENBID.Bytes, uint(ie.ShortMacroNgENBID.NumBits), &aper.Constraint{Lb: 18, Ub: 18}, false)
	case 3:
		err = w.WriteBitString(ie.LongMacroNgENBID.Bytes, uint(ie.LongMacroNgENBID.NumBits), &aper.Constraint{Lb: 21, Ub: 21}, false)
	}
	return
}
func (ie *NgENBID) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var b []byte
		var n uint
		if b, n, err = r.ReadBitString(&aper.Constraint{Lb: 20, Ub: 20}, false); err != nil {
			return
		}
		ie.MacroNgENBID = &aper.BitString{Bytes: b, NumBits: uint64(n)}
	case 2:
		var b []byte
		var n uint
		if b, n, err = r.ReadBitString(&aper.Constraint{Lb: 18, Ub: 18}, false); err != nil {
			return
		}
		ie.ShortMacroNgENBID = &aper.BitString{Bytes: b, NumBits: uint64(n)}
	case 3:
		var b []byte
		var n uint
		if b, n, err = r.ReadBitString(&aper.Constraint{Lb: 21, Ub: 21}, false); err != nil {
			return
		}
		ie.LongMacroNgENBID = &aper.BitString{Bytes: b, NumBits: uint64(n)}
	}
	return
}
