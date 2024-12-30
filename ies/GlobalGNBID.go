package ies

import "github.com/lvdund/ngap/aper"

type GlobalGNBID struct {
	PLMNIdentity []byte
	GNBID        GNBID
	// IEExtensions  *GlobalGNBIDExtIEs
}

func (ie *GlobalGNBID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PLMNIdentity := NewOCTETSTRING(ie.PLMNIdentity, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_PLMNIdentity.Encode(w); err != nil {
		return
	}
	if err = ie.GNBID.Encode(w); err != nil {
		return
	}
	return
}
func (ie *GlobalGNBID) Decode(r *aper.AperReader) (err error) {
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
	if err = ie.GNBID.Decode(r); err != nil {
		return
	}
	return
}
