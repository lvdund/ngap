package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type UPTransportLayerInformationItem struct {
	NGUUPTNLInformation UPTransportLayerInformation `madatory`
	// IEExtensions *UPTransportLayerInformationItemExtIEs `optional`
}

func (ie *UPTransportLayerInformationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.NGUUPTNLInformation.Encode(w); err != nil {
		err = fmt.Errorf("Encode NGUUPTNLInformation: %w", err)
		return
	}
	return
}
func (ie *UPTransportLayerInformationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.NGUUPTNLInformation.Decode(r); err != nil {
		err = fmt.Errorf("Read NGUUPTNLInformation: %w", err)
		return
	}
	return
}
