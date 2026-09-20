package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type SecurityIndication struct {
	IntegrityProtectionIndication       IntegrityProtectionIndication       `madatory`
	ConfidentialityProtectionIndication ConfidentialityProtectionIndication `madatory`
	MaximumIntegrityProtectedDataRateUL *MaximumIntegrityProtectedDataRate  `optional`
	// IEExtensions *SecurityIndicationExtIEs `optional`
}

func (ie *SecurityIndication) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.MaximumIntegrityProtectedDataRateUL != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.IntegrityProtectionIndication.Encode(w); err != nil {
		err = fmt.Errorf("Encode IntegrityProtectionIndication: %w", err)
		return
	}
	if err = ie.ConfidentialityProtectionIndication.Encode(w); err != nil {
		err = fmt.Errorf("Encode ConfidentialityProtectionIndication: %w", err)
		return
	}
	if ie.MaximumIntegrityProtectedDataRateUL != nil {
		if err = ie.MaximumIntegrityProtectedDataRateUL.Encode(w); err != nil {
			err = fmt.Errorf("Encode MaximumIntegrityProtectedDataRateUL: %w", err)
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
	if err = ie.IntegrityProtectionIndication.Decode(r); err != nil {
		err = fmt.Errorf("Read IntegrityProtectionIndication: %w", err)
		return
	}
	if err = ie.ConfidentialityProtectionIndication.Decode(r); err != nil {
		err = fmt.Errorf("Read ConfidentialityProtectionIndication: %w", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(MaximumIntegrityProtectedDataRate)
		if err = tmp.Decode(r); err != nil {
			err = fmt.Errorf("Read MaximumIntegrityProtectedDataRateUL: %w", err)
			return
		}
		ie.MaximumIntegrityProtectedDataRateUL = tmp
	}
	return
}
