package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type RATRestrictionsItem struct {
	PLMNIdentity              []byte         `lb:3,ub:3,madatory`
	RATRestrictionInformation aper.BitString `lb:8,ub:8,madatory,valExt`
	// IEExtensions *RATRestrictionsItemExtIEs `optional`
}

func (ie *RATRestrictionsItem) Encode(w *aper.AperWriter) (err error) {
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
	tmp_RATRestrictionInformation := NewBITSTRING(ie.RATRestrictionInformation, aper.Constraint{Lb: 8, Ub: 8}, true)
	if err = tmp_RATRestrictionInformation.Encode(w); err != nil {
		err = utils.WrapError("Encode RATRestrictionInformation", err)
		return
	}
	return
}
func (ie *RATRestrictionsItem) Decode(r *aper.AperReader) (err error) {
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
	tmp_RATRestrictionInformation := BITSTRING{
		c:   aper.Constraint{Lb: 8, Ub: 8},
		ext: true,
	}
	if err = tmp_RATRestrictionInformation.Decode(r); err != nil {
		err = utils.WrapError("Read RATRestrictionInformation", err)
		return
	}
	ie.RATRestrictionInformation = aper.BitString{Bytes: tmp_RATRestrictionInformation.Value.Bytes, NumBits: tmp_RATRestrictionInformation.Value.NumBits}
	return
}
