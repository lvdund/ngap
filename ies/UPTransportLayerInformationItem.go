package ies

import "github.com/lvdund/ngap/aper"

type UPTransportLayerInformationItem struct {
	NGUUPTNLInformation *UPTransportLayerInformation `False,`
	// IEExtensions UPTransportLayerInformationItemExtIEs `False,OPTIONAL`
}

func (ie *UPTransportLayerInformationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.NGUUPTNLInformation != nil {
		if err = ie.NGUUPTNLInformation.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *UPTransportLayerInformationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.NGUUPTNLInformation = new(UPTransportLayerInformation)
	if err = ie.NGUUPTNLInformation.Decode(r); err != nil {
		return
	}
	return
}
