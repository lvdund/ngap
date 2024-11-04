package ies

import "github.com/lvdund/ngap/aper"

const (
	UEIdentityIndexValuePresentNothing uint64 = iota /* No components present */
	UEIdentityIndexValuePresentIndexLength10
	UEIdentityIndexValuePresentChoiceExtensions
)


type UEIdentityIndexValue struct {
	Choice        uint64
	IndexLength10 *aper.BitString `False,,10,10`
	// ChoiceExtensions *UEIdentityIndexValueExtIEs `False,,,`
}

func (ie *UEIdentityIndexValue) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case UEIdentityIndexValuePresentIndexLength10:
		err = w.WriteBitString(ie.IndexLength10.Bytes, uint(ie.IndexLength10.NumBits), &aper.Constraint{Lb: 10, Ub: 10}, false)
	}
	return
}
func (ie *UEIdentityIndexValue) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case UEIdentityIndexValuePresentIndexLength10:
		var b []byte
		var n uint
		if b, n, err = r.ReadBitString(&aper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
			return
		}
		ie.IndexLength10 = &aper.BitString{Bytes: b, NumBits: uint64(n)}
	}
	return
}
