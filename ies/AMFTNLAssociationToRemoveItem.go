package ies

import "github.com/lvdund/ngap/aper"

type AMFTNLAssociationToRemoveItem struct {
	AMFTNLAssociationAddress CPTransportLayerInformation
	// IEExtensions  *AMFTNLAssociationToRemoveItemExtIEs
}

func (ie *AMFTNLAssociationToRemoveItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.AMFTNLAssociationAddress.Encode(w); err != nil {
		return
	}
	return
}
func (ie *AMFTNLAssociationToRemoveItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.AMFTNLAssociationAddress.Decode(r); err != nil {
		return
	}
	return
}
