package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type NGRANTNLAssociationToRemoveItem struct {
	TNLAssociationTransportLayerAddress    CPTransportLayerInformation  `madatory`
	TNLAssociationTransportLayerAddressAMF *CPTransportLayerInformation `optional`
	// IEExtensions *NGRANTNLAssociationToRemoveItemExtIEs `optional`
}

func (ie *NGRANTNLAssociationToRemoveItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.TNLAssociationTransportLayerAddressAMF != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.TNLAssociationTransportLayerAddress.Encode(w); err != nil {
		err = fmt.Errorf("Encode TNLAssociationTransportLayerAddress: %w", err)
		return
	}
	if ie.TNLAssociationTransportLayerAddressAMF != nil {
		if err = ie.TNLAssociationTransportLayerAddressAMF.Encode(w); err != nil {
			err = fmt.Errorf("Encode TNLAssociationTransportLayerAddressAMF: %w", err)
			return
		}
	}
	return
}
func (ie *NGRANTNLAssociationToRemoveItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.TNLAssociationTransportLayerAddress.Decode(r); err != nil {
		err = fmt.Errorf("Read TNLAssociationTransportLayerAddress: %w", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(CPTransportLayerInformation)
		if err = tmp.Decode(r); err != nil {
			err = fmt.Errorf("Read TNLAssociationTransportLayerAddressAMF: %w", err)
			return
		}
		ie.TNLAssociationTransportLayerAddressAMF = tmp
	}
	return
}
