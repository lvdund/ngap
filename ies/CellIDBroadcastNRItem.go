package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type CellIDBroadcastNRItem struct {
	NRCGI NRCGI `madatory`
	// IEExtensions *CellIDBroadcastNRItemExtIEs `optional`
}

func (ie *CellIDBroadcastNRItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.NRCGI.Encode(w); err != nil {
		err = fmt.Errorf("Encode NRCGI: %w", err)
		return
	}
	return
}
func (ie *CellIDBroadcastNRItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.NRCGI.Decode(r); err != nil {
		err = fmt.Errorf("Read NRCGI: %w", err)
		return
	}
	return
}
