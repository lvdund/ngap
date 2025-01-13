package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SliceOverloadItem struct {
	SNSSAI SNSSAI `madatory`
	// IEExtensions *SliceOverloadItemExtIEs `optional`
}

func (ie *SliceOverloadItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.SNSSAI.Encode(w); err != nil {
		err = utils.WrapError("Encode SNSSAI", err)
		return
	}
	return
}
func (ie *SliceOverloadItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.SNSSAI.Decode(r); err != nil {
		err = utils.WrapError("Read SNSSAI", err)
		return
	}
	return
}
