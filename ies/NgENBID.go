package ies

import "github.com/lvdund/ngap/aper"

const (
	NgENBIDPresentNothing uint64 = iota
	NgENBIDPresentMacrongenbId
	NgENBIDPresentShortmacrongenbId
	NgENBIDPresentLongmacrongenbId
	NgENBIDPresentChoiceExtensions
)

type NgENBID struct {
	Choice            uint64
	MacroNgENBID      *BITSTRING
	ShortMacroNgENBID *BITSTRING
	LongMacroNgENBID  *BITSTRING
	// ChoiceExtensions *NgENBIDExtIEs
}

func (ie *NgENBID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case NgENBIDPresentMacrongenbId:
		err = ie.MacroNgENBID.Encode(w)
	case NgENBIDPresentShortmacrongenbId:
		err = ie.ShortMacroNgENBID.Encode(w)
	case NgENBIDPresentLongmacrongenbId:
		err = ie.LongMacroNgENBID.Encode(w)
	}
	return
}
func (ie *NgENBID) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return
	}
	switch ie.Choice {
	case NgENBIDPresentMacrongenbId:
		var tmp BITSTRING
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.MacroNgENBID = &tmp
	case NgENBIDPresentShortmacrongenbId:
		var tmp BITSTRING
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.ShortMacroNgENBID = &tmp
	case NgENBIDPresentLongmacrongenbId:
		var tmp BITSTRING
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.LongMacroNgENBID = &tmp
	}
	return
}
