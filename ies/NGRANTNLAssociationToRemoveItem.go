package ies

import "github.com/lvdund/ngap/aper"

type NGRANTNLAssociationToRemoveItem struct {
	TNLAssociationTransportLayerAddress    *CPTransportLayerInformation `False,`
	TNLAssociationTransportLayerAddressAMF *CPTransportLayerInformation `False,OPTIONAL`
	// IEExtensions NGRANTNLAssociationToRemoveItemExtIEs `False,OPTIONAL`
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
	if ie.TNLAssociationTransportLayerAddress != nil {
		if err = ie.TNLAssociationTransportLayerAddress.Encode(w); err != nil {
			return
		}
	}
	if ie.TNLAssociationTransportLayerAddressAMF != nil {
		if err = ie.TNLAssociationTransportLayerAddressAMF.Encode(w); err != nil {
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
	ie.TNLAssociationTransportLayerAddress = new(CPTransportLayerInformation)
	ie.TNLAssociationTransportLayerAddressAMF = new(CPTransportLayerInformation)
	if err = ie.TNLAssociationTransportLayerAddress.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.TNLAssociationTransportLayerAddressAMF.Decode(r); err != nil {
			return
		}
	}
	return
}
