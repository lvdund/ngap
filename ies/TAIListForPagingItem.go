package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TAIListForPagingItem struct {
	TAI TAI `madatory`
	// IEExtensions *TAIListForPagingItemExtIEs `optional`
}

func (ie *TAIListForPagingItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.TAI.Encode(w); err != nil {
		err = utils.WrapError("Encode TAI", err)
		return
	}
	return
}
func (ie *TAIListForPagingItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.TAI.Decode(r); err != nil {
		err = utils.WrapError("Read TAI", err)
		return
	}
	return
}
