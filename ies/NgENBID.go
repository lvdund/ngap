package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	NgENBIDPresentNothing uint64 = iota
	NgENBIDPresentMacrongenbId
	NgENBIDPresentShortmacrongenbId
	NgENBIDPresentLongmacrongenbId
	NgENBIDPresentChoiceExtensions
)

type NgENBID struct {
	Choice            uint64
	MacroNgENBID      *aper.BitString
	ShortMacroNgENBID *aper.BitString
	LongMacroNgENBID  *aper.BitString
	// ChoiceExtensions *NgENBIDExtIEs
}

func (ie *NgENBID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case NgENBIDPresentMacrongenbId:
		tmp := NewBITSTRING(*ie.MacroNgENBID, aper.Constraint{Lb: 20, Ub: 20}, false)
		err = tmp.Encode(w)
	case NgENBIDPresentShortmacrongenbId:
		tmp := NewBITSTRING(*ie.ShortMacroNgENBID, aper.Constraint{Lb: 18, Ub: 18}, false)
		err = tmp.Encode(w)
	case NgENBIDPresentLongmacrongenbId:
		tmp := NewBITSTRING(*ie.LongMacroNgENBID, aper.Constraint{Lb: 21, Ub: 21}, false)
		err = tmp.Encode(w)
	}
	return
}
func (ie *NgENBID) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case NgENBIDPresentMacrongenbId:
		tmp := BITSTRING{c: aper.Constraint{Lb: 20, Ub: 20}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read MacroNgENBID", err)
			return
		}
		ie.MacroNgENBID = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	case NgENBIDPresentShortmacrongenbId:
		tmp := BITSTRING{c: aper.Constraint{Lb: 18, Ub: 18}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read ShortMacroNgENBID", err)
			return
		}
		ie.ShortMacroNgENBID = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	case NgENBIDPresentLongmacrongenbId:
		tmp := BITSTRING{c: aper.Constraint{Lb: 21, Ub: 21}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read LongMacroNgENBID", err)
			return
		}
		ie.LongMacroNgENBID = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	}
	return
}
