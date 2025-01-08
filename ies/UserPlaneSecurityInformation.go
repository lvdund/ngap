package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type UserPlaneSecurityInformation struct {
	SecurityResult     SecurityResult
	SecurityIndication SecurityIndication
	// IEExtensions *UserPlaneSecurityInformationExtIEs `optional`
}

func (ie *UserPlaneSecurityInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.SecurityResult.Encode(w); err != nil {
		err = utils.WrapError("Read SecurityResult", err)
		return
	}
	if err = ie.SecurityIndication.Encode(w); err != nil {
		err = utils.WrapError("Read SecurityIndication", err)
		return
	}
	return
}
func (ie *UserPlaneSecurityInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.SecurityResult.Decode(r); err != nil {
		err = utils.WrapError("Read SecurityResult", err)
		return
	}
	if err = ie.SecurityIndication.Decode(r); err != nil {
		err = utils.WrapError("Read SecurityIndication", err)
		return
	}
	return
}
