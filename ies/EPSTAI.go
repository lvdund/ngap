package ies

import "github.com/lvdund/ngap/aper"

type EPSTAI struct {
	PLMNIdentity []byte
	EPSTAC       []byte
	// IEExtensions  *EPSTAIExtIEs
}

func (ie *EPSTAI) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PLMNIdentity := NewOCTETSTRING(ie.PLMNIdentity, aper.Constraint{Lb: 3, Ub: 3}, true)
	if err = tmp_PLMNIdentity.Encode(w); err != nil {
		return
	}
	tmp_EPSTAC := NewOCTETSTRING(ie.EPSTAC, aper.Constraint{Lb: 2, Ub: 2}, true)
	if err = tmp_EPSTAC.Encode(w); err != nil {
		return
	}
	return
}
func (ie *EPSTAI) Decode(r *aper.AperReader) (err error) {
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
	tmp_EPSTAC := OCTETSTRING{
		c:   aper.Constraint{Lb: 2, Ub: 2},
		ext: false,
	}
	if err = tmp_EPSTAC.Decode(r); err != nil {
		return
	}
	ie.EPSTAC = tmp_EPSTAC.Value
	return
}
