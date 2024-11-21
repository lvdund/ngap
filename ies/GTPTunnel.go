package ies

import "github.com/lvdund/ngap/aper"

type GTPTunnel struct {
	TransportLayerAddress *TransportLayerAddress `False,`
	GTPTEID               *GTPTEID               `False,`
	// IEExtensions GTPTunnelExtIEs `False,OPTIONAL`
}

func (ie *GTPTunnel) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.TransportLayerAddress != nil {
		if err = ie.TransportLayerAddress.Encode(w); err != nil {
			return
		}
	}
	if ie.GTPTEID != nil {
		if err = ie.GTPTEID.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *GTPTunnel) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.TransportLayerAddress = new(TransportLayerAddress)
	ie.GTPTEID = new(GTPTEID)
	if err = ie.TransportLayerAddress.Decode(r); err != nil {
		return
	}
	if err = ie.GTPTEID.Decode(r); err != nil {
		return
	}
	return
}
