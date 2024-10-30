package ies

import "github.com/lvdund/ngap/aper"

type EndpointIPAddressAndPort struct {
	EndpointIPAddress *TransportLayerAddress `False,`
	PortNumber        *PortNumber            `False,`
	// IEExtensions EndpointIPAddressAndPortExtIEs `False,OPTIONAL`
}

func (ie *EndpointIPAddressAndPort) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.EndpointIPAddress != nil {
		if err = ie.EndpointIPAddress.Encode(w); err != nil {
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
func (ie *EndpointIPAddressAndPort) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.EndpointIPAddress = new(TransportLayerAddress)
	ie.PortNumber = new(PortNumber)
	if err = ie.EndpointIPAddress.Decode(r); err != nil {
		return
	}
	if err = ie.PortNumber.Decode(r); err != nil {
		return
	}
	return
}
