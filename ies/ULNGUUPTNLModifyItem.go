package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ULNGUUPTNLModifyItem struct {
	ULNGUUPTNLInformation UPTransportLayerInformation `madatory`
	DLNGUUPTNLInformation UPTransportLayerInformation `madatory`
	// IEExtensions *ULNGUUPTNLModifyItemExtIEs `optional`
}

func (ie *ULNGUUPTNLModifyItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.ULNGUUPTNLInformation.Encode(w); err != nil {
		err = utils.WrapError("Encode ULNGUUPTNLInformation", err)
		return
	}
	if err = ie.DLNGUUPTNLInformation.Encode(w); err != nil {
		err = utils.WrapError("Encode DLNGUUPTNLInformation", err)
		return
	}
	return
}
func (ie *ULNGUUPTNLModifyItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.ULNGUUPTNLInformation.Decode(r); err != nil {
		err = utils.WrapError("Read ULNGUUPTNLInformation", err)
		return
	}
	if err = ie.DLNGUUPTNLInformation.Decode(r); err != nil {
		err = utils.WrapError("Read DLNGUUPTNLInformation", err)
		return
	}
	return
}
