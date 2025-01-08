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
	MacroNgENBID      []byte
	ShortMacroNgENBID []byte
	LongMacroNgENBID  []byte
	// ChoiceExtensions *NgENBIDExtIEs
}

func (ie *NgENBID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case NgENBIDPresentMacrongenbId:
		tmp := NewBITSTRING(ie.MacroNgENBID, aper.Constraint{Lb: 20, Ub: 20}, false)
		err = tmp.Encode(w)
	case NgENBIDPresentShortmacrongenbId:
		tmp := NewBITSTRING(ie.ShortMacroNgENBID, aper.Constraint{Lb: 18, Ub: 18}, false)
		err = tmp.Encode(w)
	case NgENBIDPresentLongmacrongenbId:
		tmp := NewBITSTRING(ie.LongMacroNgENBID, aper.Constraint{Lb: 21, Ub: 21}, false)
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
		tmp := NewBITSTRING(nil, aper.Constraint{Lb: 20, Ub: 20}, false)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read MacroNgENBID", err)
			return
		}
		ie.MacroNgENBID = tmp.Value.Bytes
	case NgENBIDPresentShortmacrongenbId:
		tmp := NewBITSTRING(nil, aper.Constraint{Lb: 18, Ub: 18}, false)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read ShortMacroNgENBID", err)
			return
		}
		ie.ShortMacroNgENBID = tmp.Value.Bytes
	case NgENBIDPresentLongmacrongenbId:
		tmp := NewBITSTRING(nil, aper.Constraint{Lb: 21, Ub: 21}, false)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read LongMacroNgENBID", err)
			return
		}
		ie.LongMacroNgENBID = tmp.Value.Bytes
	}
	return
}
