package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type SecurityResult struct {
	IntegrityProtectionResult       IntegrityProtectionResult       `madatory`
	ConfidentialityProtectionResult ConfidentialityProtectionResult `madatory`
	// IEExtensions *SecurityResultExtIEs `optional`
}

func (ie *SecurityResult) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.IntegrityProtectionResult.Encode(w); err != nil {
		err = fmt.Errorf("Encode IntegrityProtectionResult: %w", err)
		return
	}
	if err = ie.ConfidentialityProtectionResult.Encode(w); err != nil {
		err = fmt.Errorf("Encode ConfidentialityProtectionResult: %w", err)
		return
	}
	return
}
func (ie *SecurityResult) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.IntegrityProtectionResult.Decode(r); err != nil {
		err = fmt.Errorf("Read IntegrityProtectionResult: %w", err)
		return
	}
	if err = ie.ConfidentialityProtectionResult.Decode(r); err != nil {
		err = fmt.Errorf("Read ConfidentialityProtectionResult: %w", err)
		return
	}
	return
}
