package ies

import "github.com/lvdund/ngap/aper"

type GlobalN3IWFID struct {
	PLMNIdentity []byte
	N3IWFID      N3IWFID
	// IEExtensions  *GlobalN3IWFIDExtIEs
}

func (ie *GlobalN3IWFID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PLMNIdentity := NewOCTETSTRING(ie.PLMNIdentity, aper.Constraint{Lb: 3, Ub: 3}, true)
	if err = tmp_PLMNIdentity.Encode(w); err != nil {
		return
	}
	if err = ie.N3IWFID.Encode(w); err != nil {
		return
	}
	return
}
func (ie *GlobalN3IWFID) Decode(r *aper.AperReader) (err error) {
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
	if err = ie.N3IWFID.Decode(r); err != nil {
		return
	}
	return
}
