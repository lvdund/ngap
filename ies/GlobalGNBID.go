package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type GlobalGNBID struct {
	PLMNIdentity []byte `lb:3,ub:3,madatory`
	GNBID        GNBID  `madatory`
	// IEExtensions *GlobalGNBIDExtIEs `optional`
}

func (ie *GlobalGNBID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PLMNIdentity := NewOCTETSTRING(ie.PLMNIdentity, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_PLMNIdentity.Encode(w); err != nil {
		err = fmt.Errorf("Encode PLMNIdentity: %w", err)
		return
	}
	if err = ie.GNBID.Encode(w); err != nil {
		err = fmt.Errorf("Encode GNBID: %w", err)
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
		err = fmt.Errorf("Read PLMNIdentity: %w", err)
		return
	}
	ie.PLMNIdentity = tmp_PLMNIdentity.Value
	if err = ie.GNBID.Decode(r); err != nil {
		err = fmt.Errorf("Read GNBID: %w", err)
		return
	}
	return
}
