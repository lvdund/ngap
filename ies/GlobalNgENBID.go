package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type GlobalNgENBID struct {
	PLMNIdentity []byte  `lb:3,ub:3,madatory`
	NgENBID      NgENBID `madatory`
	// IEExtensions *GlobalNgENBIDExtIEs `optional`
}

func (ie *GlobalNgENBID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PLMNIdentity := NewOCTETSTRING(ie.PLMNIdentity, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_PLMNIdentity.Encode(w); err != nil {
		err = utils.WrapError("Encode PLMNIdentity", err)
		return
	}
	if err = ie.NgENBID.Encode(w); err != nil {
		err = utils.WrapError("Encode NgENBID", err)
		return
	}
	return
}
func (ie *GlobalNgENBID) Decode(r *aper.AperReader) (err error) {
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
		err = utils.WrapError("Read PLMNIdentity", err)
		return
	}
	ie.PLMNIdentity = tmp_PLMNIdentity.Value
	if err = ie.NgENBID.Decode(r); err != nil {
		err = utils.WrapError("Read NgENBID", err)
		return
	}
	return
}
