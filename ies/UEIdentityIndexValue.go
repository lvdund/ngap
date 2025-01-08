package ies

import "github.com/lvdund/ngap/aper"

const (
	UEIdentityIndexValuePresentNothing uint64 = iota
	UEIdentityIndexValuePresentIndexlength10
	UEIdentityIndexValuePresentChoiceExtensions
)

type UEIdentityIndexValue struct {
	Choice        uint64
	IndexLength10 []byte
	// ChoiceExtensions *UEIdentityIndexValueExtIEs
}

func (ie *UEIdentityIndexValue) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case UEIdentityIndexValuePresentIndexlength10:
		tmp := NewBITSTRING(ie.IndexLength10, aper.Constraint{Lb: 10, Ub: 10}, false)
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
		tmp := NewBITSTRING(nil, aper.Constraint{Lb: 10, Ub: 10}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.IndexLength10 = tmp.Value.Bytes
	}
	return
}
