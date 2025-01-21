package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AMFTNLAssociationToAddItem struct {
	AMFTNLAssociationAddress CPTransportLayerInformation `madatory`
	TNLAssociationUsage      *TNLAssociationUsage        `optional`
	TNLAddressWeightFactor   int64                       `lb:0,ub:255,madatory`
	// IEExtensions *AMFTNLAssociationToAddItemExtIEs `optional`
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
	if err = ie.AMFTNLAssociationAddress.Encode(w); err != nil {
		err = utils.WrapError("Encode AMFTNLAssociationAddress", err)
		return
	}
	if ie.TNLAssociationUsage != nil {
		if err = ie.TNLAssociationUsage.Encode(w); err != nil {
			err = utils.WrapError("Encode TNLAssociationUsage", err)
			return
		}
	}
	tmp_TNLAddressWeightFactor := NewINTEGER(ie.TNLAddressWeightFactor, aper.Constraint{Lb: 0, Ub: 255}, false)
	if err = tmp_TNLAddressWeightFactor.Encode(w); err != nil {
		err = utils.WrapError("Encode TNLAddressWeightFactor", err)
		return
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
	if err = ie.AMFTNLAssociationAddress.Decode(r); err != nil {
		err = utils.WrapError("Read AMFTNLAssociationAddress", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(TNLAssociationUsage)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read TNLAssociationUsage", err)
			return
		}
		ie.TNLAssociationUsage = tmp
	}
	tmp_TNLAddressWeightFactor := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 255},
		ext: false,
	}
	if err = tmp_TNLAddressWeightFactor.Decode(r); err != nil {
		err = utils.WrapError("Read TNLAddressWeightFactor", err)
		return
	}
	ie.TNLAddressWeightFactor = int64(tmp_TNLAddressWeightFactor.Value)
	return
}
