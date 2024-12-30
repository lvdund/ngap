package ies

import "github.com/lvdund/ngap/aper"

const (
	UEIdentityIndexValuePresentNothing uint64 = iota
	UEIdentityIndexValuePresentIndexlength10
	UEIdentityIndexValuePresentChoiceExtensions
)

type UEIdentityIndexValue struct {
	Choice        uint64
	IndexLength10 *BITSTRING
	// ChoiceExtensions *UEIdentityIndexValueExtIEs
}

func (ie *UEIdentityIndexValue) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case UEIdentityIndexValuePresentIndexlength10:
		err = ie.IndexLength10.Encode(w)
	}
	return
}
func (ie *UEIdentityIndexValue) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case UEIdentityIndexValuePresentIndexlength10:
		var tmp BITSTRING
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.IndexLength10 = &tmp
	}
	return
}
