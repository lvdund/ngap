package ies

import "github.com/lvdund/ngap/aper"

type UPTransportLayerInformationPairItem struct {
	ULNGUUPTNLInformation UPTransportLayerInformation
	DLNGUUPTNLInformation UPTransportLayerInformation
	// IEExtensions  *UPTransportLayerInformationPairItemExtIEs
}

func (ie *UPTransportLayerInformationPairItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.ULNGUUPTNLInformation.Encode(w); err != nil {
		return
	}
	if err = ie.DLNGUUPTNLInformation.Encode(w); err != nil {
		return
	}
	return
}
func (ie *UPTransportLayerInformationPairItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.ULNGUUPTNLInformation.Decode(r); err != nil {
		return
	}
	if err = ie.DLNGUUPTNLInformation.Decode(r); err != nil {
		return
	}
	return
}
