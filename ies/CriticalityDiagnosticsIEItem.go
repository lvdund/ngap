package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type CriticalityDiagnosticsIEItem struct {
	IECriticality Criticality
	IEID          ProtocolIEID
	TypeOfError   TypeOfError
	// IEExtensions *CriticalityDiagnosticsIEItemExtIEs `optional`
}

func (ie *CriticalityDiagnosticsIEItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.IECriticality.Encode(w); err != nil {
		err = utils.WrapError("Read IECriticality", err)
		return
	}
	if err = ie.IEID.Encode(w); err != nil {
		err = utils.WrapError("Read IEID", err)
		return
	}
	if err = ie.TypeOfError.Encode(w); err != nil {
		err = utils.WrapError("Read TypeOfError", err)
		return
	}
	return
}
func (ie *CriticalityDiagnosticsIEItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.IECriticality.Decode(r); err != nil {
		err = utils.WrapError("Read IECriticality", err)
		return
	}
	if err = ie.IEID.Decode(r); err != nil {
		err = utils.WrapError("Read IEID", err)
		return
	}
	if err = ie.TypeOfError.Decode(r); err != nil {
		err = utils.WrapError("Read TypeOfError", err)
		return
	}
	return
}
