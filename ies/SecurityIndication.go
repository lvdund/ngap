package ies

import "github.com/lvdund/ngap/aper"

type SecurityIndication struct {
	IntegrityProtectionIndication       *IntegrityProtectionIndication       `False,`
	ConfidentialityProtectionIndication *ConfidentialityProtectionIndication `False,`
	MaximumIntegrityProtectedDataRateUL *MaximumIntegrityProtectedDataRate   `False,OPTIONAL`
	// IEExtensions SecurityIndicationExtIEs `False,OPTIONAL`
}

func (ie *SecurityIndication) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.MaximumIntegrityProtectedDataRateUL != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.IntegrityProtectionIndication != nil {
		if err = ie.IntegrityProtectionIndication.Encode(w); err != nil {
			return
		}
	}
	if ie.ConfidentialityProtectionIndication != nil {
		if err = ie.ConfidentialityProtectionIndication.Encode(w); err != nil {
			return
		}
	}
	if ie.MaximumIntegrityProtectedDataRateUL != nil {
		if err = ie.MaximumIntegrityProtectedDataRateUL.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *SecurityIndication) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	ie.IntegrityProtectionIndication = new(IntegrityProtectionIndication)
	ie.ConfidentialityProtectionIndication = new(ConfidentialityProtectionIndication)
	ie.MaximumIntegrityProtectedDataRateUL = new(MaximumIntegrityProtectedDataRate)
	if err = ie.IntegrityProtectionIndication.Decode(r); err != nil {
		return
	}
	if err = ie.ConfidentialityProtectionIndication.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.MaximumIntegrityProtectedDataRateUL.Decode(r); err != nil {
			return
		}
	}
	return
}
