package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SecurityIndication struct {
	IntegrityProtectionIndication       IntegrityProtectionIndication
	ConfidentialityProtectionIndication ConfidentialityProtectionIndication
	MaximumIntegrityProtectedDataRateUL *MaximumIntegrityProtectedDataRate `optional`
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
		err = utils.WrapError("Read IntegrityProtectionIndication", err)
		return
	}
	if err = ie.ConfidentialityProtectionIndication.Encode(w); err != nil {
		err = utils.WrapError("Read ConfidentialityProtectionIndication", err)
		return
	}
	if ie.MaximumIntegrityProtectedDataRateUL != nil {
		if err = ie.MaximumIntegrityProtectedDataRateUL.Encode(w); err != nil {
			err = utils.WrapError("Read MaximumIntegrityProtectedDataRateUL", err)
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
		err = utils.WrapError("Read IntegrityProtectionIndication", err)
		return
	}
	if err = ie.ConfidentialityProtectionIndication.Decode(r); err != nil {
		err = utils.WrapError("Read ConfidentialityProtectionIndication", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.MaximumIntegrityProtectedDataRateUL.Decode(r); err != nil {
			err = utils.WrapError("Read MaximumIntegrityProtectedDataRateUL", err)
			return
		}
	}
	return
}
