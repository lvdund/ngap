package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
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
		err = utils.WrapError("Encode IntegrityProtectionResult", err)
		return
	}
	if err = ie.ConfidentialityProtectionResult.Encode(w); err != nil {
		err = utils.WrapError("Encode ConfidentialityProtectionResult", err)
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
		err = utils.WrapError("Read IntegrityProtectionResult", err)
		return
	}
	if err = ie.ConfidentialityProtectionResult.Decode(r); err != nil {
		err = utils.WrapError("Read ConfidentialityProtectionResult", err)
		return
	}
	return
}
