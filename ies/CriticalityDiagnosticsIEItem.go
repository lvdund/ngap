package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type CriticalityDiagnosticsIEItem struct {
	IECriticality Criticality  `madatory`
	IEID          ProtocolIEID `madatory`
	TypeOfError   TypeOfError  `madatory`
	// IEExtensions *CriticalityDiagnosticsIEItemExtIEs `optional`
}

func (ie *CriticalityDiagnosticsIEItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.IECriticality.Encode(w); err != nil {
		err = fmt.Errorf("Encode IECriticality: %w", err)
		return
	}
	if err = ie.IEID.Encode(w); err != nil {
		err = fmt.Errorf("Encode IEID: %w", err)
		return
	}
	if err = ie.TypeOfError.Encode(w); err != nil {
		err = fmt.Errorf("Encode TypeOfError: %w", err)
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
		err = fmt.Errorf("Read IECriticality: %w", err)
		return
	}
	if err = ie.IEID.Decode(r); err != nil {
		err = fmt.Errorf("Read IEID: %w", err)
		return
	}
	if err = ie.TypeOfError.Decode(r); err != nil {
		err = fmt.Errorf("Read TypeOfError: %w", err)
		return
	}
	return
}
