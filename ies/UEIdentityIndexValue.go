package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	UEIdentityIndexValuePresentNothing uint64 = iota
	UEIdentityIndexValuePresentIndexlength10
	UEIdentityIndexValuePresentChoiceExtensions
)

type UEIdentityIndexValue struct {
	Choice        uint64
	IndexLength10 *aper.BitString `lb:10,ub:10`
	// ChoiceExtensions *UEIdentityIndexValueExtIEs
}

func (ie *UEIdentityIndexValue) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case UEIdentityIndexValuePresentIndexlength10:
		tmp := NewBITSTRING(*ie.IndexLength10, aper.Constraint{Lb: 10, Ub: 10}, false)
		err = tmp.Encode(w)
	}
	return
}
func (ie *UEIdentityIndexValue) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return
	}
	switch ie.Choice {
	case UEIdentityIndexValuePresentIndexlength10:
		tmp := BITSTRING{c: aper.Constraint{Lb: 10, Ub: 10}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read IndexLength10", err)
			return
		}
		ie.IndexLength10 = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	}
	return
}
