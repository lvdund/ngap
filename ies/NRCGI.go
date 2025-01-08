package ies

import "github.com/lvdund/ngap/aper"

type NRCGI struct {
	PLMNIdentity   []byte
	NRCellIdentity []byte
	// IEExtensions *NRCGIExtIEs `optional`
}

func (ie *NRCGI) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PLMNIdentity := NewOCTETSTRING(ie.PLMNIdentity, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_PLMNIdentity.Encode(w); err != nil {
		return
	}
	tmp_NRCellIdentity := NewBITSTRING(ie.NRCellIdentity, aper.Constraint{Lb: 36, Ub: 36}, false)
	if err = tmp_NRCellIdentity.Encode(w); err != nil {
		return
	}
	return
}
func (ie *NRCGI) Decode(r *aper.AperReader) (err error) {
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
	tmp_NRCellIdentity := BITSTRING{
		c:   aper.Constraint{Lb: 36, Ub: 36},
		ext: false,
	}
	if err = tmp_NRCellIdentity.Decode(r); err != nil {
		return
	}
	ie.NRCellIdentity = tmp_NRCellIdentity.Value.Bytes
	return
}
