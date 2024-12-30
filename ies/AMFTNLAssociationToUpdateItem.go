package ies

import "github.com/lvdund/ngap/aper"

type AMFTNLAssociationToUpdateItem struct {
	AMFTNLAssociationAddress CPTransportLayerInformation
	TNLAssociationUsage      *TNLAssociationUsage
	TNLAddressWeightFactor   *int64
	// IEExtensions  *AMFTNLAssociationToUpdateItemExtIEs
}

func (ie *AMFTNLAssociationToUpdateItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.TNLAssociationUsage != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.TNLAddressWeightFactor != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if err = ie.AMFTNLAssociationAddress.Encode(w); err != nil {
		return
	}
	if ie.TNLAssociationUsage != nil {
		if err = ie.TNLAssociationUsage.Encode(w); err != nil {
			return
		}
	}
	if ie.TNLAddressWeightFactor != nil {
		tmp_TNLAddressWeightFactor := NewINTEGER(*ie.TNLAddressWeightFactor, aper.Constraint{Lb: 0, Ub: 255}, true)
		if err = tmp_TNLAddressWeightFactor.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *AMFTNLAssociationToUpdateItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if err = ie.AMFTNLAssociationAddress.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.TNLAssociationUsage.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_TNLAddressWeightFactor := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 255},
			ext: false,
		}
		if err = tmp_TNLAddressWeightFactor.Decode(r); err != nil {
			return
		}
		*ie.TNLAddressWeightFactor = int64(tmp_TNLAddressWeightFactor.Value)
	}
	return
}
