package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AreaOfInterestCellItem struct {
	NGRANCGI NGRANCGI `madatory`
	// IEExtensions *AreaOfInterestCellItemExtIEs `optional`
}

func (ie *AreaOfInterestCellItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.NGRANCGI.Encode(w); err != nil {
		err = utils.WrapError("Encode NGRANCGI", err)
		return
	}
	return
}
func (ie *AreaOfInterestCellItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.NGRANCGI.Decode(r); err != nil {
		err = utils.WrapError("Read NGRANCGI", err)
		return
	}
	return
}
