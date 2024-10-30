package ies

import "github.com/lvdund/ngap/aper"

type XnTNLConfigurationInfo struct {
	XnTransportLayerAddresses         *XnTLAs    `False,`
	XnExtendedTransportLayerAddresses *XnExtTLAs `False,OPTIONAL`
	// IEExtensions XnTNLConfigurationInfoExtIEs `False,OPTIONAL`
}

func (ie *XnTNLConfigurationInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.XnExtendedTransportLayerAddresses != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.XnTransportLayerAddresses != nil {
		if err = ie.XnTransportLayerAddresses.Encode(w); err != nil {
			return
		}
	}
	if ie.XnExtendedTransportLayerAddresses != nil {
		if err = ie.XnExtendedTransportLayerAddresses.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *XnTNLConfigurationInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	ie.XnTransportLayerAddresses = new(XnTLAs)
	ie.XnExtendedTransportLayerAddresses = new(XnExtTLAs)
	if err = ie.XnTransportLayerAddresses.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.XnExtendedTransportLayerAddresses.Decode(r); err != nil {
			return
		}
	}
	return
}
