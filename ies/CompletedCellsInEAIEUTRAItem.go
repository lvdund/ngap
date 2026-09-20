package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type CompletedCellsInEAIEUTRAItem struct {
	EUTRACGI EUTRACGI `madatory`
	// IEExtensions *CompletedCellsInEAIEUTRAItemExtIEs `optional`
}

func (ie *CompletedCellsInEAIEUTRAItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.EUTRACGI.Encode(w); err != nil {
		err = fmt.Errorf("Encode EUTRACGI: %w", err)
		return
	}
	return
}
func (ie *CompletedCellsInEAIEUTRAItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.EUTRACGI.Decode(r); err != nil {
		err = fmt.Errorf("Read EUTRACGI: %w", err)
		return
	}
	return
}
