package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type UESecurityCapabilities struct {
	NRencryptionAlgorithms             aper.BitString `lb:16,ub:16,madatory,valExt`
	NRintegrityProtectionAlgorithms    aper.BitString `lb:16,ub:16,madatory,valExt`
	EUTRAencryptionAlgorithms          aper.BitString `lb:16,ub:16,madatory,valExt`
	EUTRAintegrityProtectionAlgorithms aper.BitString `lb:16,ub:16,madatory,valExt`
	// IEExtensions *UESecurityCapabilitiesExtIEs `optional`
}

func (ie *UESecurityCapabilities) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_NRencryptionAlgorithms := NewBITSTRING(ie.NRencryptionAlgorithms, aper.Constraint{Lb: 16, Ub: 16}, true)
	if err = tmp_NRencryptionAlgorithms.Encode(w); err != nil {
		err = utils.WrapError("Encode NRencryptionAlgorithms", err)
		return
	}
	tmp_NRintegrityProtectionAlgorithms := NewBITSTRING(ie.NRintegrityProtectionAlgorithms, aper.Constraint{Lb: 16, Ub: 16}, true)
	if err = tmp_NRintegrityProtectionAlgorithms.Encode(w); err != nil {
		err = utils.WrapError("Encode NRintegrityProtectionAlgorithms", err)
		return
	}
	tmp_EUTRAencryptionAlgorithms := NewBITSTRING(ie.EUTRAencryptionAlgorithms, aper.Constraint{Lb: 16, Ub: 16}, true)
	if err = tmp_EUTRAencryptionAlgorithms.Encode(w); err != nil {
		err = utils.WrapError("Encode EUTRAencryptionAlgorithms", err)
		return
	}
	tmp_EUTRAintegrityProtectionAlgorithms := NewBITSTRING(ie.EUTRAintegrityProtectionAlgorithms, aper.Constraint{Lb: 16, Ub: 16}, true)
	if err = tmp_EUTRAintegrityProtectionAlgorithms.Encode(w); err != nil {
		err = utils.WrapError("Encode EUTRAintegrityProtectionAlgorithms", err)
		return
	}
	return
}
func (ie *UESecurityCapabilities) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_NRencryptionAlgorithms := BITSTRING{
		c:   aper.Constraint{Lb: 16, Ub: 16},
		ext: true,
	}
	if err = tmp_NRencryptionAlgorithms.Decode(r); err != nil {
		err = utils.WrapError("Read NRencryptionAlgorithms", err)
		return
	}
	ie.NRencryptionAlgorithms = aper.BitString{Bytes: tmp_NRencryptionAlgorithms.Value.Bytes, NumBits: tmp_NRencryptionAlgorithms.Value.NumBits}
	tmp_NRintegrityProtectionAlgorithms := BITSTRING{
		c:   aper.Constraint{Lb: 16, Ub: 16},
		ext: true,
	}
	if err = tmp_NRintegrityProtectionAlgorithms.Decode(r); err != nil {
		err = utils.WrapError("Read NRintegrityProtectionAlgorithms", err)
		return
	}
	ie.NRintegrityProtectionAlgorithms = aper.BitString{Bytes: tmp_NRintegrityProtectionAlgorithms.Value.Bytes, NumBits: tmp_NRintegrityProtectionAlgorithms.Value.NumBits}
	tmp_EUTRAencryptionAlgorithms := BITSTRING{
		c:   aper.Constraint{Lb: 16, Ub: 16},
		ext: true,
	}
	if err = tmp_EUTRAencryptionAlgorithms.Decode(r); err != nil {
		err = utils.WrapError("Read EUTRAencryptionAlgorithms", err)
		return
	}
	ie.EUTRAencryptionAlgorithms = aper.BitString{Bytes: tmp_EUTRAencryptionAlgorithms.Value.Bytes, NumBits: tmp_EUTRAencryptionAlgorithms.Value.NumBits}
	tmp_EUTRAintegrityProtectionAlgorithms := BITSTRING{
		c:   aper.Constraint{Lb: 16, Ub: 16},
		ext: true,
	}
	if err = tmp_EUTRAintegrityProtectionAlgorithms.Decode(r); err != nil {
		err = utils.WrapError("Read EUTRAintegrityProtectionAlgorithms", err)
		return
	}
	ie.EUTRAintegrityProtectionAlgorithms = aper.BitString{Bytes: tmp_EUTRAintegrityProtectionAlgorithms.Value.Bytes, NumBits: tmp_EUTRAintegrityProtectionAlgorithms.Value.NumBits}
	return
}
