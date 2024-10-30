package ies

import "github.com/lvdund/ngap/aper"

type UserLocationInformationN3IWF struct {
	IPAddress  *TransportLayerAddress `False,`
	PortNumber *PortNumber            `False,`
	// IEExtensions UserLocationInformationN3IWFExtIEs `False,OPTIONAL`
}

func (ie *UserLocationInformationN3IWF) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.IPAddress != nil {
		if err = ie.IPAddress.Encode(w); err != nil {
			return
		}
	}
	if ie.PortNumber != nil {
		if err = ie.PortNumber.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *UserLocationInformationN3IWF) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.IPAddress = new(TransportLayerAddress)
	ie.PortNumber = new(PortNumber)
	if err = ie.IPAddress.Decode(r); err != nil {
		return
	}
	if err = ie.PortNumber.Decode(r); err != nil {
		return
	}
	return
}
