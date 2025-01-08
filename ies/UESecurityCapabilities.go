package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type UESecurityCapabilities struct {
	NRencryptionAlgorithms             []byte
	NRintegrityProtectionAlgorithms    []byte
	EUTRAencryptionAlgorithms          []byte
	EUTRAintegrityProtectionAlgorithms []byte
	// IEExtensions *UESecurityCapabilitiesExtIEs `optional`
}

func (ie *UESecurityCapabilities) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_NRencryptionAlgorithms := NewBITSTRING(ie.NRencryptionAlgorithms, aper.Constraint{Lb: 16, Ub: 16}, false)
	if err = tmp_NRencryptionAlgorithms.Encode(w); err != nil {
		err = utils.WrapError("Read NRencryptionAlgorithms", err)
		return
	}
	tmp_NRintegrityProtectionAlgorithms := NewBITSTRING(ie.NRintegrityProtectionAlgorithms, aper.Constraint{Lb: 16, Ub: 16}, false)
	if err = tmp_NRintegrityProtectionAlgorithms.Encode(w); err != nil {
		err = utils.WrapError("Read NRintegrityProtectionAlgorithms", err)
		return
	}
	tmp_EUTRAencryptionAlgorithms := NewBITSTRING(ie.EUTRAencryptionAlgorithms, aper.Constraint{Lb: 16, Ub: 16}, false)
	if err = tmp_EUTRAencryptionAlgorithms.Encode(w); err != nil {
		err = utils.WrapError("Read EUTRAencryptionAlgorithms", err)
		return
	}
	tmp_EUTRAintegrityProtectionAlgorithms := NewBITSTRING(ie.EUTRAintegrityProtectionAlgorithms, aper.Constraint{Lb: 16, Ub: 16}, false)
	if err = tmp_EUTRAintegrityProtectionAlgorithms.Encode(w); err != nil {
		err = utils.WrapError("Read EUTRAintegrityProtectionAlgorithms", err)
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
		ext: false,
	}
	if err = tmp_NRencryptionAlgorithms.Decode(r); err != nil {
		err = utils.WrapError("Read NRencryptionAlgorithms", err)
		return
	}
	ie.NRencryptionAlgorithms = tmp_NRencryptionAlgorithms.Value.Bytes
	tmp_NRintegrityProtectionAlgorithms := BITSTRING{
		c:   aper.Constraint{Lb: 16, Ub: 16},
		ext: false,
	}
	if err = tmp_NRintegrityProtectionAlgorithms.Decode(r); err != nil {
		err = utils.WrapError("Read NRintegrityProtectionAlgorithms", err)
		return
	}
	ie.NRintegrityProtectionAlgorithms = tmp_NRintegrityProtectionAlgorithms.Value.Bytes
	tmp_EUTRAencryptionAlgorithms := BITSTRING{
		c:   aper.Constraint{Lb: 16, Ub: 16},
		ext: false,
	}
	if err = tmp_EUTRAencryptionAlgorithms.Decode(r); err != nil {
		err = utils.WrapError("Read EUTRAencryptionAlgorithms", err)
		return
	}
	ie.EUTRAencryptionAlgorithms = tmp_EUTRAencryptionAlgorithms.Value.Bytes
	tmp_EUTRAintegrityProtectionAlgorithms := BITSTRING{
		c:   aper.Constraint{Lb: 16, Ub: 16},
		ext: false,
	}
	if err = tmp_EUTRAintegrityProtectionAlgorithms.Decode(r); err != nil {
		err = utils.WrapError("Read EUTRAintegrityProtectionAlgorithms", err)
		return
	}
	ie.EUTRAintegrityProtectionAlgorithms = tmp_EUTRAintegrityProtectionAlgorithms.Value.Bytes
	return
}
