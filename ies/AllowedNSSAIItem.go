package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type AllowedNSSAIItem struct {
	SNSSAI SNSSAI `madatory`
	// IEExtensions *AllowedNSSAIItemExtIEs `optional`
}

func (ie *AllowedNSSAIItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.SNSSAI.Encode(w); err != nil {
		err = fmt.Errorf("Encode SNSSAI: %w", err)
		return
	}
	return
}
func (ie *AllowedNSSAIItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.SNSSAI.Decode(r); err != nil {
		err = fmt.Errorf("Read SNSSAI: %w", err)
		return
	}
	return
}
