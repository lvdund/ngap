package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type LastVisitedCellItem struct {
	LastVisitedCellInformation LastVisitedCellInformation `madatory`
	// IEExtensions *LastVisitedCellItemExtIEs `optional`
}

func (ie *LastVisitedCellItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.LastVisitedCellInformation.Encode(w); err != nil {
		err = fmt.Errorf("Encode LastVisitedCellInformation: %w", err)
		return
	}
	return
}
func (ie *LastVisitedCellItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.LastVisitedCellInformation.Decode(r); err != nil {
		err = fmt.Errorf("Read LastVisitedCellInformation: %w", err)
		return
	}
	return
}
