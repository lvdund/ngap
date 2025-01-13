package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type GTPTunnel struct {
	TransportLayerAddress []byte `lb:1,ub:160,madatory,valExt`
	GTPTEID               []byte `lb:4,ub:4,madatory`
	// IEExtensions *GTPTunnelExtIEs `optional`
}

func (ie *GTPTunnel) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_TransportLayerAddress := NewBITSTRING(ie.TransportLayerAddress, aper.Constraint{Lb: 1, Ub: 160}, true)
	if err = tmp_TransportLayerAddress.Encode(w); err != nil {
		err = utils.WrapError("Encode TransportLayerAddress", err)
		return
	}
	tmp_GTPTEID := NewOCTETSTRING(ie.GTPTEID, aper.Constraint{Lb: 4, Ub: 4}, false)
	if err = tmp_GTPTEID.Encode(w); err != nil {
		err = utils.WrapError("Encode GTPTEID", err)
		return
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
	tmp_TransportLayerAddress := BITSTRING{
		c:   aper.Constraint{Lb: 1, Ub: 160},
		ext: true,
	}
	if err = tmp_TransportLayerAddress.Decode(r); err != nil {
		err = utils.WrapError("Read TransportLayerAddress", err)
		return
	}
	ie.TransportLayerAddress = tmp_TransportLayerAddress.Value.Bytes
	tmp_GTPTEID := OCTETSTRING{
		c:   aper.Constraint{Lb: 4, Ub: 4},
		ext: false,
	}
	if err = tmp_GTPTEID.Decode(r); err != nil {
		err = utils.WrapError("Read GTPTEID", err)
		return
	}
	ie.GTPTEID = tmp_GTPTEID.Value
	return
}
