package ies

import "github.com/lvdund/ngap/aper"

type AMFTNLAssociationToAddItem struct {
	AMFTNLAssociationAddress *CPTransportLayerInformation `False,`
	TNLAssociationUsage      *TNLAssociationUsage         `False,OPTIONAL`
	TNLAddressWeightFactor   *TNLAddressWeightFactor      `False,`
	// IEExtensions AMFTNLAssociationToAddItemExtIEs `False,OPTIONAL`
}

func (ie *AMFTNLAssociationToAddItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.TNLAssociationUsage != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.AMFTNLAssociationAddress != nil {
		if err = ie.AMFTNLAssociationAddress.Encode(w); err != nil {
			return
		}
	}
	if ie.TNLAssociationUsage != nil {
		if err = ie.TNLAssociationUsage.Encode(w); err != nil {
			return
		}
	}
	if ie.TNLAddressWeightFactor != nil {
		if err = ie.TNLAddressWeightFactor.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *AMFTNLAssociationToAddItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	ie.AMFTNLAssociationAddress = new(CPTransportLayerInformation)
	ie.TNLAssociationUsage = new(TNLAssociationUsage)
	ie.TNLAddressWeightFactor = new(TNLAddressWeightFactor)
	if err = ie.AMFTNLAssociationAddress.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.TNLAssociationUsage.Decode(r); err != nil {
			return
		}
	}
	if err = ie.TNLAddressWeightFactor.Decode(r); err != nil {
		return
	}
	return
}
