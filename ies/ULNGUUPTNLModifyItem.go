package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
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
		err = fmt.Errorf("Encode ULNGUUPTNLInformation: %w", err)
		return
	}
	if err = ie.DLNGUUPTNLInformation.Encode(w); err != nil {
		err = fmt.Errorf("Encode DLNGUUPTNLInformation: %w", err)
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
		err = fmt.Errorf("Read ULNGUUPTNLInformation: %w", err)
		return
	}
	if err = ie.DLNGUUPTNLInformation.Decode(r); err != nil {
		err = fmt.Errorf("Read DLNGUUPTNLInformation: %w", err)
		return
	}
	return
}
