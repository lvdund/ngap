package ies

import "github.com/lvdund/ngap/aper"

type UESecurityCapabilities struct {
	NRencryptionAlgorithms             *NRencryptionAlgorithms             `False,`
	NRintegrityProtectionAlgorithms    *NRintegrityProtectionAlgorithms    `False,`
	EUTRAencryptionAlgorithms          *EUTRAencryptionAlgorithms          `False,`
	EUTRAintegrityProtectionAlgorithms *EUTRAintegrityProtectionAlgorithms `False,`
	// IEExtensions UESecurityCapabilitiesExtIEs `False,OPTIONAL`
}

func (ie *UESecurityCapabilities) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.NRencryptionAlgorithms != nil {
		if err = ie.NRencryptionAlgorithms.Encode(w); err != nil {
			return
		}
	}
	if ie.NRintegrityProtectionAlgorithms != nil {
		if err = ie.NRintegrityProtectionAlgorithms.Encode(w); err != nil {
			return
		}
	}
	if ie.EUTRAencryptionAlgorithms != nil {
		if err = ie.EUTRAencryptionAlgorithms.Encode(w); err != nil {
			return
		}
	}
	if ie.EUTRAintegrityProtectionAlgorithms != nil {
		if err = ie.EUTRAintegrityProtectionAlgorithms.Encode(w); err != nil {
			return
		}
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
	ie.NRencryptionAlgorithms = new(NRencryptionAlgorithms)
	ie.NRintegrityProtectionAlgorithms = new(NRintegrityProtectionAlgorithms)
	ie.EUTRAencryptionAlgorithms = new(EUTRAencryptionAlgorithms)
	ie.EUTRAintegrityProtectionAlgorithms = new(EUTRAintegrityProtectionAlgorithms)
	if err = ie.NRencryptionAlgorithms.Decode(r); err != nil {
		return
	}
	if err = ie.NRintegrityProtectionAlgorithms.Decode(r); err != nil {
		return
	}
	if err = ie.EUTRAencryptionAlgorithms.Decode(r); err != nil {
		return
	}
	if err = ie.EUTRAintegrityProtectionAlgorithms.Decode(r); err != nil {
		return
	}
	return
}
