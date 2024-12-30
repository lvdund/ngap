package ies

import "github.com/lvdund/ngap/aper"

type TAI struct {
	PLMNIdentity []byte
	TAC          []byte
	// IEExtensions  *TAIExtIEs
}

func (ie *TAI) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PLMNIdentity := NewOCTETSTRING(ie.PLMNIdentity, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_PLMNIdentity.Encode(w); err != nil {
		return
	}
	tmp_TAC := NewOCTETSTRING(ie.TAC, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_TAC.Encode(w); err != nil {
		return
	}
	return
}
func (ie *TAI) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_PLMNIdentity := OCTETSTRING{
		c:   aper.Constraint{Lb: 3, Ub: 3},
		ext: false,
	}
	if err = tmp_PLMNIdentity.Decode(r); err != nil {
		return
	}
	ie.PLMNIdentity = tmp_PLMNIdentity.Value
	tmp_TAC := OCTETSTRING{
		c:   aper.Constraint{Lb: 3, Ub: 3},
		ext: false,
	}
	if err = tmp_TAC.Decode(r); err != nil {
		return
	}
	ie.TAC = tmp_TAC.Value
	return
}
