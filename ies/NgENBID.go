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
	MacroNgENBID      *[]byte
	ShortMacroNgENBID *[]byte
	LongMacroNgENBID  *[]byte
	// ChoiceExtensions *NgENBIDExtIEs
}

func (ie *NgENBID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case NgENBIDPresentMacrongenbId:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
			Value: aper.BitString{
				Bytes:   *ie.MacroNgENBID,
				NumBits: uint64(len(*ie.MacroNgENBID)),
			},
		}
		err = tmp.Encode(w)
	case NgENBIDPresentShortmacrongenbId:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
			Value: aper.BitString{
				Bytes:   *ie.ShortMacroNgENBID,
				NumBits: uint64(len(*ie.ShortMacroNgENBID)),
			},
		}
		err = tmp.Encode(w)
	case NgENBIDPresentLongmacrongenbId:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
			Value: aper.BitString{
				Bytes:   *ie.LongMacroNgENBID,
				NumBits: uint64(len(*ie.LongMacroNgENBID)),
			},
		}
		err = tmp.Encode(w)
	}
	return
}
func (ie *NgENBID) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return
	}
	switch ie.Choice {
	case NgENBIDPresentMacrongenbId:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.MacroNgENBID = &tmp.Value.Bytes
	case NgENBIDPresentShortmacrongenbId:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.ShortMacroNgENBID = &tmp.Value.Bytes
	case NgENBIDPresentLongmacrongenbId:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.LongMacroNgENBID = &tmp.Value.Bytes
	}
	return
}
