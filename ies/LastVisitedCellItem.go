package ies

import "github.com/lvdund/ngap/aper"

type LastVisitedCellItem struct {
	LastVisitedCellInformation LastVisitedCellInformation
	// IEExtensions  *LastVisitedCellItemExtIEs
}

func (ie *LastVisitedCellItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.LastVisitedCellInformation.Encode(w); err != nil {
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
		return
	}
	return
}
