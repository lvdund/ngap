package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AMFTNLAssociationSetupItem struct {
	AMFTNLAssociationAddress CPTransportLayerInformation `madatory`
	// IEExtensions *AMFTNLAssociationSetupItemExtIEs `optional`
}

func (ie *AMFTNLAssociationSetupItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.AMFTNLAssociationAddress.Encode(w); err != nil {
		err = utils.WrapError("Encode AMFTNLAssociationAddress", err)
		return
	}
	return
}
func (ie *AMFTNLAssociationSetupItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.AMFTNLAssociationAddress.Decode(r); err != nil {
		err = utils.WrapError("Read AMFTNLAssociationAddress", err)
		return
	}
	return
}
