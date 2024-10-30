package ies

import "github.com/lvdund/ngap/aper"

type SecurityResult struct {
	IntegrityProtectionResult       *IntegrityProtectionResult       `False,`
	ConfidentialityProtectionResult *ConfidentialityProtectionResult `False,`
	// IEExtensions SecurityResultExtIEs `False,OPTIONAL`
}

func (ie *SecurityResult) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.IntegrityProtectionResult != nil {
		if err = ie.IntegrityProtectionResult.Encode(w); err != nil {
			return
		}
	}
	if ie.ConfidentialityProtectionResult != nil {
		if err = ie.ConfidentialityProtectionResult.Encode(w); err != nil {
			return
		}
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
	ie.IntegrityProtectionResult = new(IntegrityProtectionResult)
	ie.ConfidentialityProtectionResult = new(ConfidentialityProtectionResult)
	if err = ie.IntegrityProtectionResult.Decode(r); err != nil {
		return
	}
	if err = ie.ConfidentialityProtectionResult.Decode(r); err != nil {
		return
	}
	return
}
