package ies

import "github.com/lvdund/ngap/aper"

type SecurityResult struct {
	IntegrityProtectionResult       IntegrityProtectionResult
	ConfidentialityProtectionResult ConfidentialityProtectionResult
	// IEExtensions  *SecurityResultExtIEs
}

func (ie *SecurityResult) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.IntegrityProtectionResult.Encode(w); err != nil {
		return
	}
	if err = ie.ConfidentialityProtectionResult.Encode(w); err != nil {
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
		return
	}
	if err = ie.ConfidentialityProtectionResult.Decode(r); err != nil {
		return
	}
	return
}
