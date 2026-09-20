package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type TNLAssociationItem struct {
	TNLAssociationAddress CPTransportLayerInformation `madatory`
	Cause                 Cause                       `madatory`
	// IEExtensions *TNLAssociationItemExtIEs `optional`
}

func (ie *TNLAssociationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.TNLAssociationAddress.Encode(w); err != nil {
		err = fmt.Errorf("Encode TNLAssociationAddress: %w", err)
		return
	}
	if err = ie.Cause.Encode(w); err != nil {
		err = fmt.Errorf("Encode Cause: %w", err)
		return
	}
	return
}
func (ie *TNLAssociationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.TNLAssociationAddress.Decode(r); err != nil {
		err = fmt.Errorf("Read TNLAssociationAddress: %w", err)
		return
	}
	if err = ie.Cause.Decode(r); err != nil {
		err = fmt.Errorf("Read Cause: %w", err)
		return
	}
	return
}
